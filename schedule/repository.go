package schedule

import (
	"database/sql"
	"fmt"
	"student-app/internal/models"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(item *models.Schedule) error {
	query := `
        INSERT INTO schedule (group_name, group_id, discipline, teacher, teacher_id,
                              start_time, end_time, date, day_of_week, audience,
                              lesson_type, week_type, comment, is_from_api, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
        RETURNING id, created_at
    `
	return r.db.QueryRow(query, item.GroupName, item.GroupID, item.Discipline,
		item.Teacher, item.TeacherID, item.StartTime, item.EndTime, item.Date,
		item.DayOfWeek, item.Audience, item.LessonType, item.WeekType,
		item.Comment, item.IsFromAPI, item.CreatedBy).Scan(&item.ID, &item.CreatedAt)
}

func (r *Repository) Update(id int, updates map[string]interface{}) error {
	query := `UPDATE schedule SET updated_at = NOW()`
	args := []interface{}{}
	i := 1

	if groupName, ok := updates["group_name"]; ok {
		query += fmt.Sprintf(", group_name = $%d", i)
		args = append(args, groupName)
		i++
	}
	if discipline, ok := updates["discipline"]; ok {
		query += fmt.Sprintf(", discipline = $%d", i)
		args = append(args, discipline)
		i++
	}
	if teacher, ok := updates["teacher"]; ok {
		query += fmt.Sprintf(", teacher = $%d", i)
		args = append(args, teacher)
		i++
	}
	if startTime, ok := updates["start_time"]; ok {
		query += fmt.Sprintf(", start_time = $%d", i)
		args = append(args, startTime)
		i++
	}
	if endTime, ok := updates["end_time"]; ok {
		query += fmt.Sprintf(", end_time = $%d", i)
		args = append(args, endTime)
		i++
	}
	if date, ok := updates["date"]; ok {
		query += fmt.Sprintf(", date = $%d", i)
		args = append(args, date)
		i++
	}
	if audience, ok := updates["audience"]; ok {
		query += fmt.Sprintf(", audience = $%d", i)
		args = append(args, audience)
		i++
	}
	if comment, ok := updates["comment"]; ok {
		query += fmt.Sprintf(", comment = $%d", i)
		args = append(args, comment)
		i++
	}

	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM schedule WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) FindByID(id int) (*models.Schedule, error) {
	var item models.Schedule
	query := `
        SELECT id, group_name, group_id, discipline, teacher, teacher_id,
               start_time, end_time, date, day_of_week, audience,
               lesson_type, week_type, comment, is_from_api, created_by, created_at
        FROM schedule
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&item.ID, &item.GroupName, &item.GroupID, &item.Discipline,
		&item.Teacher, &item.TeacherID, &item.StartTime, &item.EndTime,
		&item.Date, &item.DayOfWeek, &item.Audience, &item.LessonType,
		&item.WeekType, &item.Comment, &item.IsFromAPI, &item.CreatedBy, &item.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *Repository) FindByGroup(groupName string, date time.Time) ([]models.Schedule, error) {
	var items []models.Schedule
	query := `
        SELECT id, group_name, group_id, discipline, teacher, teacher_id,
               start_time, end_time, date, day_of_week, audience,
               lesson_type, week_type, comment
        FROM schedule
        WHERE group_name = $1 AND date = $2
        ORDER BY start_time
    `
	rows, err := r.db.Query(query, groupName, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Schedule
		err := rows.Scan(&item.ID, &item.GroupName, &item.GroupID, &item.Discipline,
			&item.Teacher, &item.TeacherID, &item.StartTime, &item.EndTime,
			&item.Date, &item.DayOfWeek, &item.Audience, &item.LessonType,
			&item.WeekType, &item.Comment)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) FindByGroupAndDateRange(groupName, startDate, endDate string) ([]models.Schedule, error) {
	var items []models.Schedule
	query := `
        SELECT id, group_name, group_id, discipline, teacher, teacher_id,
               start_time, end_time, date, day_of_week, audience,
               lesson_type, week_type, comment
        FROM schedule
        WHERE group_name = $1 AND date BETWEEN $2 AND $3
        ORDER BY date, start_time
    `
	rows, err := r.db.Query(query, groupName, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Schedule
		err := rows.Scan(&item.ID, &item.GroupName, &item.GroupID, &item.Discipline,
			&item.Teacher, &item.TeacherID, &item.StartTime, &item.EndTime,
			&item.Date, &item.DayOfWeek, &item.Audience, &item.LessonType,
			&item.WeekType, &item.Comment)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) SaveCache(cache *models.ScheduleCache) error {
	query := `
        INSERT INTO schedule_cache (group_id, group_name, data)
        VALUES ($1, $2, $3)
        ON CONFLICT (group_id) DO UPDATE SET data = EXCLUDED.data, fetched_at = NOW()
    `
	_, err := r.db.Exec(query, cache.GroupID, cache.GroupName, cache.Data)
	return err
}

func (r *Repository) GetCache(groupID int) (*models.ScheduleCache, error) {
	var cache models.ScheduleCache
	query := `SELECT id, group_id, group_name, data, fetched_at FROM schedule_cache WHERE group_id = $1`
	err := r.db.QueryRow(query, groupID).Scan(&cache.ID, &cache.GroupID, &cache.GroupName, &cache.Data, &cache.FetchedAt)
	if err != nil {
		return nil, err
	}
	return &cache, nil
}
