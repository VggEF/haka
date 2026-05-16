package users

import (
	"errors"
	"student-app/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Основные операции
func (s *Service) Create(req *CreateUserRequest) (*models.User, error) {
	existing, _ := s.repo.FindByLogin(req.Login)
	if existing != nil {
		return nil, errors.New("пользователь уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ошибка хэширования пароля")
	}

	role := req.Role
	if role == "" {
		role = "student"
	}

	user := &models.User{
		Login:    req.Login,
		Password: string(hashedPassword),
		Name:     req.Name,
		Role:     role,
		Group:    req.Group,
		Course:   req.Course,
		Faculty:  req.Faculty,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(id int, req *UpdateUserRequest) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Telegram != "" {
		updates["telegram"] = req.Telegram
	}
	if req.VK != "" {
		updates["vk"] = req.VK
	}
	if req.GitHub != "" {
		updates["github"] = req.GitHub
	}
	if req.Photo != "" {
		updates["photo"] = req.Photo
	}
	if req.Group != "" {
		updates["group_name"] = req.Group
	}
	if req.Course > 0 {
		updates["course"] = req.Course
	}
	if req.Faculty != "" {
		updates["faculty"] = req.Faculty
	}

	return s.repo.Update(id, updates)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetAll(role string, limit, offset int) ([]models.User, error) {
	return s.repo.GetAll(role, limit, offset)
}

// Профили студентов
func (s *Service) GetStudentProfile(userID int) (*StudentProfileResponse, error) {
	profile, err := s.repo.GetStudentProfile(userID)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return &StudentProfileResponse{
			Hobbies: []string{},
			Clubs:   []string{},
			About:   "",
		}, nil
	}
	return &StudentProfileResponse{
		Hobbies: profile.Hobbies,
		Clubs:   profile.Clubs,
		About:   profile.About,
	}, nil
}

func (s *Service) UpdateStudentProfile(userID int, req *UpdateStudentProfileRequest) error {
	profile := &models.StudentProfile{
		UserID:  userID,
		Hobbies: req.Hobbies,
		Clubs:   req.Clubs,
		About:   req.About,
	}
	return s.repo.CreateOrUpdateStudentProfile(profile)
}

// Профили преподавателей
func (s *Service) GetTeacherProfile(userID int) (*TeacherProfileResponse, error) {
	profile, err := s.repo.GetTeacherProfile(userID)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return &TeacherProfileResponse{}, nil
	}
	return &TeacherProfileResponse{
		Department: profile.Department,
		Position:   profile.Position,
		Degree:     profile.Degree,
		Office:     profile.Office,
		Experience: profile.Experience,
	}, nil
}

func (s *Service) UpdateTeacherProfile(userID int, req *UpdateTeacherProfileRequest) error {
	profile := &models.TeacherProfile{
		UserID:     userID,
		Department: req.Department,
		Position:   req.Position,
		Degree:     req.Degree,
		Office:     req.Office,
		Experience: req.Experience,
	}
	return s.repo.CreateOrUpdateTeacherProfile(profile)
}
