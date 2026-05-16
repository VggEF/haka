package users

import (
	"database/sql"
	"fmt"
	"strings"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Основные операции с пользователями
func (r *Repository) Create(user *models.User) error {
	query := `
        INSERT INTO users (login, password, name, role, group_name, course, faculty)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, user.Login, user.Password, user.Name, user.Role, user.Group, user.Course, user.Faculty).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *Repository) FindByID(id int) (*models.User, error) {
	var user models.User
	query := `
        SELECT id, login, name, role, group_name, course, faculty,
               email, phone, telegram, vk, github, photo, total_xp, coins,
               created_at, updated_at
        FROM users
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Login, &user.Name, &user.Role, &user.Group,
		&user.Course, &user.Faculty, &user.Email, &user.Phone, &user.Telegram,
		&user.VK, &user.GitHub, &user.Photo, &user.TotalXP, &user.Coins,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByLogin(login string) (*models.User, error) {
	var user models.User
	query := `
        SELECT id, login, name, role, group_name, course, faculty,
               email, phone, total_xp, coins
        FROM users
        WHERE login = $1
    `
	err := r.db.QueryRow(query, login).Scan(
		&user.ID, &user.Login, &user.Name, &user.Role, &user.Group,
		&user.Course, &user.Faculty, &user.Email, &user.Phone,
		&user.TotalXP, &user.Coins,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Update(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	i := 1

	for key, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE users SET %s, updated_at = NOW() WHERE id = $%d",
		strings.Join(setParts, ", "), i)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) GetAll(role string, limit, offset int) ([]models.User, error) {
	query := `
        SELECT id, login, name, role, group_name, course, faculty, email, phone
        FROM users
        WHERE role = $1
        ORDER BY id
        LIMIT $2 OFFSET $3
    `
	rows, err := r.db.Query(query, role, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Login, &user.Name, &user.Role,
			&user.Group, &user.Course, &user.Faculty, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Профили студентов
func (r *Repository) GetStudentProfile(userID int) (*models.StudentProfile, error) {
	var profile models.StudentProfile
	query := `
        SELECT user_id, hobbies, clubs, about, checklist_data, planner_data
        FROM student_profiles
        WHERE user_id = $1
    `
	err := r.db.QueryRow(query, userID).Scan(
		&profile.UserID, &profile.Hobbies, &profile.Clubs, &profile.About,
		&profile.ChecklistData, &profile.PlannerData,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Repository) CreateOrUpdateStudentProfile(profile *models.StudentProfile) error {
	query := `
        INSERT INTO student_profiles (user_id, hobbies, clubs, about, checklist_data, planner_data)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (user_id) DO UPDATE SET
            hobbies = EXCLUDED.hobbies,
            clubs = EXCLUDED.clubs,
            about = EXCLUDED.about,
            checklist_data = EXCLUDED.checklist_data,
            planner_data = EXCLUDED.planner_data
    `
	_, err := r.db.Exec(query, profile.UserID, profile.Hobbies, profile.Clubs,
		profile.About, profile.ChecklistData, profile.PlannerData)
	return err
}

// Профили преподавателей
func (r *Repository) GetTeacherProfile(userID int) (*models.TeacherProfile, error) {
	var profile models.TeacherProfile
	query := `
        SELECT user_id, department, position, degree, office, experience, achievements
        FROM teacher_profiles
        WHERE user_id = $1
    `
	err := r.db.QueryRow(query, userID).Scan(
		&profile.UserID, &profile.Department, &profile.Position, &profile.Degree,
		&profile.Office, &profile.Experience, &profile.Achievements,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Repository) CreateOrUpdateTeacherProfile(profile *models.TeacherProfile) error {
	query := `
        INSERT INTO teacher_profiles (user_id, department, position, degree, office, experience, achievements)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (user_id) DO UPDATE SET
            department = EXCLUDED.department,
            position = EXCLUDED.position,
            degree = EXCLUDED.degree,
            office = EXCLUDED.office,
            experience = EXCLUDED.experience,
            achievements = EXCLUDED.achievements
    `
	_, err := r.db.Exec(query, profile.UserID, profile.Department, profile.Position,
		profile.Degree, profile.Office, profile.Experience, profile.Achievements)
	return err
}
