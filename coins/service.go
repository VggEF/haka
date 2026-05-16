package coins

import (
	"student-app/internal/models"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// GetUserBalance - получить баланс пользователя
func (s *Service) GetUserBalance(userID int) (int, error) {
	var total int
	err := s.db.Model(&models.Transaction{}).
		Where("user_id = ?", userID).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}

// GetBalance - алиас для GetUserBalance (для совместимости с handler)
func (s *Service) GetBalance(userID int) (int, error) {
	return s.GetUserBalance(userID)
}

// GetTransactions - получить историю транзакций
func (s *Service) GetTransactions(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := s.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

// GetShopItems - получить товары из магазина
func (s *Service) GetShopItems() ([]models.ShopItem, error) {
	var items []models.ShopItem
	err := s.db.Where("is_available = ?", true).
		Order("price ASC").
		Find(&items).Error
	return items, err
}

// BuyItem - покупка товара
func (s *Service) BuyItem(userID, itemID int) (*models.Transaction, *models.Purchase, error) {
	// Получаем товар
	var item models.ShopItem
	if err := s.db.First(&item, itemID).Error; err != nil {
		return nil, nil, err
	}

	if !item.IsAvailable {
		return nil, nil, gorm.ErrRecordNotFound
	}

	// Проверяем баланс
	balance, err := s.GetUserBalance(userID)
	if err != nil {
		return nil, nil, err
	}

	if balance < item.Price {
		return nil, nil, gorm.ErrRecordNotFound
	}

	// Начинаем транзакцию
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Списываем коины
	transaction := &models.Transaction{
		UserID:    userID,
		Amount:    -item.Price,
		Type:      "spend",
		Reason:    "purchase",
		ItemID:    &item.ID,
		CreatedAt: time.Now(),
	}

	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Создаем запись о покупке
	purchase := &models.Purchase{
		UserID:        userID,
		ItemID:        item.ID,
		TransactionID: transaction.ID,
		Price:         item.Price,
		CreatedAt:     time.Now(),
	}

	if err := tx.Create(purchase).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	// Обновляем сток
	if item.Stock > 0 {
		item.Stock--
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return nil, nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}

	return transaction, purchase, nil
}

// SpendCoins - списать коины (админский метод)
func (s *Service) SpendCoins(userID int, amount int, reason string, itemID *int) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID:    userID,
		Amount:    -amount,
		Type:      "spend",
		Reason:    reason,
		ItemID:    itemID,
		CreatedAt: time.Now(),
	}

	err := s.db.Create(transaction).Error
	return transaction, err
}

// AddCoins - начислить коины (админский метод)
func (s *Service) AddCoins(userID int, amount int, reason string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID:    userID,
		Amount:    amount,
		Type:      "earn",
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := s.db.Create(transaction).Error
	return transaction, err
}
