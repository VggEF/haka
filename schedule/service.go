package schedule

import (
	"encoding/json"
	"errors"
	"student-app/internal/models"
	"time"
)

type Service struct {
	repo      *Repository
	apiClient *APIClient
}

func NewService(repo *Repository, apiClient *APIClient) *Service {
	return &Service{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (s *Service) Create(req *CreateScheduleRequest, userID int) (*models.Schedule, error) {
	var date time.Time
	if req.Date != "" {
		var err error
		date, err = time.Parse("2006-01-02", req.Date)
		if err != nil {
			return nil, errors.New("неверный формат даты")
		}
	} else {
		date = time.Now()
	}

	dayOfWeek := int(date.Weekday())
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	schedule := &models.Schedule{
		GroupName:  req.GroupName,
		GroupID:    req.GroupID,
		Discipline: req.Discipline,
		Teacher:    req.Teacher,
		TeacherID:  req.TeacherID,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Date:       date,
		DayOfWeek:  dayOfWeek,
		Audience:   req.Audience,
		LessonType: req.LessonType,
		WeekType:   req.WeekType,
		Comment:    req.Comment,
		IsFromAPI:  false,
		CreatedBy:  userID,
	}

	if err := s.repo.Create(schedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *Service) Update(id int, req *UpdateScheduleRequest) error {
	updates := make(map[string]interface{})

	if req.GroupName != "" {
		updates["group_name"] = req.GroupName
	}
	if req.Discipline != "" {
		updates["discipline"] = req.Discipline
	}
	if req.Teacher != "" {
		updates["teacher"] = req.Teacher
	}
	if req.StartTime != "" {
		updates["start_time"] = req.StartTime
	}
	if req.EndTime != "" {
		updates["end_time"] = req.EndTime
	}
	if req.Date != "" {
		updates["date"] = req.Date
	}
	if req.Audience != "" {
		updates["audience"] = req.Audience
	}
	if req.LessonType != "" {
		updates["lesson_type"] = req.LessonType
	}
	if req.Comment != "" {
		updates["comment"] = req.Comment
	}

	return s.repo.Update(id, updates)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByID(id int) (*ScheduleResponse, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &ScheduleResponse{
		ID:         item.ID,
		GroupName:  item.GroupName,
		GroupID:    item.GroupID,
		Discipline: item.Discipline,
		Teacher:    item.Teacher,
		TeacherID:  item.TeacherID,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Date:       item.Date,
		DayOfWeek:  item.DayOfWeek,
		Audience:   item.Audience,
		LessonType: item.LessonType,
		WeekType:   item.WeekType,
		Comment:    item.Comment,
	}, nil
}

func (s *Service) GetSchedule(query *GetScheduleQuery) ([]ScheduleResponse, error) {
	var items []models.Schedule
	var err error

	if query.GroupName != "" {
		if query.Date != "" {
			date, _ := time.Parse("2006-01-02", query.Date)
			items, err = s.repo.FindByGroup(query.GroupName, date)
		} else if query.StartDate != "" && query.EndDate != "" {
			items, err = s.repo.FindByGroupAndDateRange(query.GroupName, query.StartDate, query.EndDate)
		} else {
			items, err = s.repo.FindByGroup(query.GroupName, time.Now())
		}
	}

	if err != nil {
		return nil, err
	}

	var response []ScheduleResponse
	for _, item := range items {
		response = append(response, ScheduleResponse{
			ID:         item.ID,
			GroupName:  item.GroupName,
			GroupID:    item.GroupID,
			Discipline: item.Discipline,
			Teacher:    item.Teacher,
			TeacherID:  item.TeacherID,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Date:       item.Date,
			DayOfWeek:  item.DayOfWeek,
			Audience:   item.Audience,
			LessonType: item.LessonType,
			WeekType:   item.WeekType,
			Comment:    item.Comment,
		})
	}
	return response, nil
}

func (s *Service) SyncFromAPI(groupID int, date string) error {
	resp, err := s.apiClient.GetSchedule(groupID, date)
	if err != nil {
		return err
	}

	if resp.State != 1 {
		return errors.New("ошибка получения расписания из API")
	}

	// Сохраняем в кэш
	data, _ := json.Marshal(resp)
	cache := &models.ScheduleCache{
		GroupID:   groupID,
		Data:      string(data),
		FetchedAt: time.Now(),
	}
	s.repo.SaveCache(cache)

	// Очищаем старые записи для этой группы
	// и сохраняем новые
	for _, lesson := range resp.Data.Rasp {
		date, _ := time.Parse("2006-01-02", lesson.Date[:10])

		schedule := &models.Schedule{
			GroupName:  lesson.Group,
			GroupID:    groupID,
			Discipline: lesson.Discipline,
			Teacher:    lesson.Teacher,
			StartTime:  lesson.Start,
			EndTime:    lesson.End,
			Date:       date,
			DayOfWeek:  lesson.DayOfWeek,
			Audience:   lesson.Audience,
			LessonType: s.getLessonType(lesson.Discipline),
			IsFromAPI:  true,
		}
		s.repo.Create(schedule)
	}

	return nil
}

func (s *Service) getLessonType(discipline string) string {
	if discipline[:3] == "лек" {
		return "lecture"
	}
	if discipline[:3] == "лаб" {
		return "lab"
	}
	if discipline[:3] == "пр." {
		return "practice"
	}
	return "other"
}

func (s *Service) GetGroups() ([]Group, error) {
	return s.apiClient.GetGroupList()
}
