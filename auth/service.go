package auth

import (
	"errors"
	"log"
	"student-app/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(req *RegisterRequest) (*models.User, error) {
	// Проверка на существующего пользователя
	existing, _ := s.repo.FindByLogin(req.Login)
	if existing != nil {
		return nil, errors.New("пользователь уже существует")
	}

	role := req.Role
	if role == "" {
		role = "student"
	}

	user := &models.User{
		Login:    req.Login,
		Password: req.Password, // Сохраняем пароль как есть (без хеширования)
		Name:     req.Name,
		Role:     role,
		Group:    req.Group,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, errors.New("ошибка создания пользователя")
	}

	return user, nil
}

func (s *Service) Login(req *LoginRequest) (*models.User, error) {
	log.Printf("🔐 Попытка входа: login=%s, password=%s", req.Login, req.Password)

	user, err := s.repo.FindByLogin(req.Login)
	if err != nil {
		log.Printf("❌ Пользователь не найден")
		return nil, errors.New("неверный логин или пароль")
	}

	log.Printf("✅ Пользователь найден: ID=%d", user.ID)
	log.Printf("📦 Пароль в БД: %s", user.Password)
	log.Printf("🔑 Введённый пароль: %s", req.Password)

	// Простое сравнение строк (без bcrypt)
	if user.Password != req.Password {
		log.Printf("❌ Пароли не совпадают!")
		return nil, errors.New("неверный логин или пароль")
	}

	log.Printf("✅ Вход успешен!")
	return user, nil
}

func (s *Service) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}
