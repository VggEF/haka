package auth

import (
	"database/sql"
	"log"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *models.User) error {
	query := `
        INSERT INTO users (login, password, name, role, group_name, total_xp, coins, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(query, user.Login, user.Password, user.Name,
		user.Role, user.Group, user.TotalXP, user.Coins).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *Repository) FindByLogin(login string) (*models.User, error) {
	log.Printf("🔍 Поиск пользователя в БД: login='%s'", login)

	var user models.User
	var groupName sql.NullString
	var email sql.NullString
	var phone sql.NullString

	query := `
        SELECT id, login, password, name, role, group_name, total_xp, coins, email, phone, created_at, updated_at
        FROM users
        WHERE login = $1
    `
	err := r.db.QueryRow(query, login).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Name,
		&user.Role,
		&groupName,
		&user.TotalXP,
		&user.Coins,
		&email,
		&phone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("❌ Ошибка поиска: %v", err)
		return nil, err
	}

	// Преобразуем NullString в обычные строки
	if groupName.Valid {
		user.Group = groupName.String
	} else {
		user.Group = ""
	}
	if email.Valid {
		user.Email = email.String
	} else {
		user.Email = ""
	}
	if phone.Valid {
		user.Phone = phone.String
	} else {
		user.Phone = ""
	}

	log.Printf("✅ Найден пользователь: ID=%d, Login=%s, Role=%s", user.ID, user.Login, user.Role)
	return &user, nil
}

func (r *Repository) FindByID(id int) (*models.User, error) {
	log.Printf("🔍 Поиск пользователя по ID: %d", id)

	var user models.User
	var groupName sql.NullString
	var email sql.NullString
	var phone sql.NullString

	query := `
        SELECT id, login, name, role, group_name, email, phone, total_xp, coins, created_at, updated_at
        FROM users
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Login,
		&user.Name,
		&user.Role,
		&groupName,
		&email,
		&phone,
		&user.TotalXP,
		&user.Coins,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("❌ Ошибка поиска по ID: %v", err)
		return nil, err
	}

	// Преобразуем NullString в обычные строки
	if groupName.Valid {
		user.Group = groupName.String
	} else {
		user.Group = ""
	}
	if email.Valid {
		user.Email = email.String
	} else {
		user.Email = ""
	}
	if phone.Valid {
		user.Phone = phone.String
	} else {
		user.Phone = ""
	}

	log.Printf("✅ Найден пользователь по ID: %d, Login=%s", user.ID, user.Login)
	return &user, nil
}
