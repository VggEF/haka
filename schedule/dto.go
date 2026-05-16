package schedule

import "time"

type CreateScheduleRequest struct {
	GroupName  string `json:"group_name" binding:"required"`
	GroupID    int    `json:"group_id"`
	Discipline string `json:"discipline" binding:"required"`
	Teacher    string `json:"teacher"`
	TeacherID  int    `json:"teacher_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Date       string `json:"date"`
	DayOfWeek  int    `json:"day_of_week"`
	Audience   string `json:"audience"`
	LessonType string `json:"lesson_type"`
	WeekType   int    `json:"week_type"`
	Comment    string `json:"comment"`
}

type UpdateScheduleRequest struct {
	GroupName  string `json:"group_name"`
	Discipline string `json:"discipline"`
	Teacher    string `json:"teacher"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Date       string `json:"date"`
	Audience   string `json:"audience"`
	LessonType string `json:"lesson_type"`
	Comment    string `json:"comment"`
}

type ScheduleResponse struct {
	ID         int       `json:"id"`
	GroupName  string    `json:"group_name"`
	GroupID    int       `json:"group_id"`
	Discipline string    `json:"discipline"`
	Teacher    string    `json:"teacher"`
	TeacherID  int       `json:"teacher_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
	Date       time.Time `json:"date"`
	DayOfWeek  int       `json:"day_of_week"`
	Audience   string    `json:"audience"`
	LessonType string    `json:"lesson_type"`
	WeekType   int       `json:"week_type"`
	Comment    string    `json:"comment"`
}

type GetScheduleQuery struct {
	GroupName string `form:"group_name"`
	GroupID   int    `form:"group_id"`
	Date      string `form:"date"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Week      int    `form:"week"`
}

type ExternalScheduleResponse struct {
	Data struct {
		Rasp []struct {
			Kod          int    `json:"код"`
			Date         string `json:"дата"`
			StartTime    string `json:"датаНачала"`
			EndTime      string `json:"датаОкончания"`
			Start        string `json:"начало"`
			End          string `json:"конец"`
			DayOfWeek    int    `json:"деньНедели"`
			DayName      string `json:"день_недели"`
			Discipline   string `json:"дисциплина"`
			Teacher      string `json:"преподаватель"`
			Audience     string `json:"аудитория"`
			Group        string `json:"группа"`
			LessonNumber int    `json:"номерЗанятия"`
			Color        string `json:"цвет"`
		} `json:"rasp"`
		Info struct {
			Group struct {
				Name    string `json:"name"`
				GroupID int    `json:"groupID"`
				Year    string `json:"year"`
			} `json:"group"`
		} `json:"info"`
	} `json:"data"`
	State int    `json:"state"`
	Msg   string `json:"msg"`
}
