package coins

import (
	"database/sql"
	"fmt"
	"student-app/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ========== Баланс ==========
func (r *Repository) GetBalance(userID int) (int, error) {
	var balance int
	query := `SELECT coins FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userID).Scan(&balance)
	return balance, err
}

func (r *Repository) AddCoins(userID, amount int, reason string) error {
	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем баланс
	updateQuery := `UPDATE users SET coins = coins + $1 WHERE id = $2`
	_, err = tx.Exec(updateQuery, amount, userID)
	if err != nil {
		return err
	}

	// Записываем транзакцию
	insertQuery := `
        INSERT INTO transactions (user_id, amount, type, reason)
        VALUES ($1, $2, 'earn', $3)
    `
	_, err = tx.Exec(insertQuery, userID, amount, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) SpendCoins(userID, amount int, reason string) error {
	// Проверяем достаточно ли средств
	balance, err := r.GetBalance(userID)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("недостаточно коинов. Нужно: %d, доступно: %d", amount, balance)
	}

	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем баланс
	updateQuery := `UPDATE users SET coins = coins - $1 WHERE id = $2`
	_, err = tx.Exec(updateQuery, amount, userID)
	if err != nil {
		return err
	}

	// Записываем транзакцию
	insertQuery := `
        INSERT INTO transactions (user_id, amount, type, reason)
        VALUES ($1, $2, 'spend', $3)
    `
	_, err = tx.Exec(insertQuery, userID, amount, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// ========== Транзакции ==========
func (r *Repository) GetTransactions(userID int, limit, offset int) ([]models.Transaction, error) {
	if limit == 0 {
		limit = 50
	}
	query := `
        SELECT id, user_id, amount, type, reason, created_at
        FROM transactions
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Type, &t.Reason, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// ========== Магазин ==========
func (r *Repository) GetShopItems() ([]models.ShopItem, error) {
	// Для начала используем предустановленные товары
	// В реальном проекте должна быть отдельная таблица
	items := []models.ShopItem{
		{ID: 1, Name: "Скидка в столовой", Description: "Скидка 10% в столовой на один обед", Price: 50, Icon: "🍕", Category: "food"},
		{ID: 2, Name: "Билет на мероприятие", Description: "Бесплатный вход на любое мероприятие", Price: 100, Icon: "🎫", Category: "entertainment"},
		{ID: 3, Name: "Принт мерча", Description: "Фирменная футболка или худи", Price: 200, Icon: "👕", Category: "merch"},
		{ID: 4, Name: "Закрытая вечеринка", Description: "Билет на закрытую студенческую вечеринку", Price: 500, Icon: "🎉", Category: "entertainment"},
		{ID: 5, Name: "Дополнительная попытка", Description: "Дополнительная попытка на экзамене", Price: 150, Icon: "🎓", Category: "study"},
		{ID: 6, Name: "Кофе в кафетерии", Description: "Бесплатный кофе в кафетерии", Price: 30, Icon: "☕", Category: "food"},
		{ID: 7, Name: "Брендированная ручка", Description: "Фирменная ручка университета", Price: 25, Icon: "✒️", Category: "merch"},
		{ID: 8, Name: "Скидка на курсы", Description: "Скидка 20% на дополнительные курсы", Price: 300, Icon: "💻", Category: "education"},
	}
	return items, nil
}

func (r *Repository) GetShopItemByID(id int) (*models.ShopItem, error) {
	items, err := r.GetShopItems()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("товар не найден")
}

func (r *Repository) BuyItem(userID, itemID int, price int) error {
	return r.SpendCoins(userID, price, fmt.Sprintf("Покупка товара #%d", itemID))
}
