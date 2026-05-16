import React, { useState, useEffect } from 'react';
import './Shop.css';

const Shop = () => {
  const [coins, setCoins] = useState(1250); // В реальном приложении брать из профиля
  const [purchases, setPurchases] = useState([]);
  const [activeTab, setActiveTab] = useState('all');

  const shopItems = {
    all: [
      { id: 1, name: 'Скидка на экзамен 10%', price: 500, icon: '🎓', type: 'discount', description: '10% скидка на пересдачу экзамена' },
      { id: 2, name: 'Освобождение от дедлайна', price: 1000, icon: '⏰', type: 'deadline', description: 'Продление дедлайна на 3 дня' },
      { id: 3, name: 'Секретная тема', price: 300, icon: '🔓', type: 'content', description: 'Доступ к дополнительным материалам' },
      { id: 4, name: 'Аватарка "Профи"', price: 200, icon: '👑', type: 'avatar', description: 'Эксклюзивная аватарка' },
      { id: 5, name: 'Билет на хакатон', price: 1500, icon: '🎫', type: 'event', description: 'Участие в студенческом хакатоне' },
      { id: 6, name: 'Ментор-сессия', price: 800, icon: '👨‍🏫', type: 'service', description: '30 минут с ментором' },
    ],
    discounts: [
      { id: 1, name: 'Скидка на экзамен 10%', price: 500, icon: '🎓', type: 'discount', description: '10% скидка на пересдачу экзамена' },
      { id: 7, name: 'Скидка на обучение 5%', price: 2000, icon: '💸', type: 'discount', description: 'Скидка на следующий семестр' },
    ],
    cosmetics: [
      { id: 4, name: 'Аватарка "Профи"', price: 200, icon: '👑', type: 'avatar', description: 'Эксклюзивная аватарка' },
      { id: 8, name: 'Тема "Тёмная"', price: 150, icon: '🌙', type: 'theme', description: 'Тёмная тема интерфейса' },
      { id: 9, name: 'Стикерпак', price: 100, icon: '😎', type: 'stickers', description: 'Набор стикеров для чата' },
    ],
    services: [
      { id: 2, name: 'Освобождение от дедлайна', price: 1000, icon: '⏰', type: 'deadline', description: 'Продление дедлайна на 3 дня' },
      { id: 6, name: 'Ментор-сессия', price: 800, icon: '👨‍🏫', type: 'service', description: '30 минут с ментором' },
      { id: 10, name: 'Проверка работы', price: 400, icon: '✅', type: 'service', description: 'Детальная проверка задания' },
    ]
  };

  const getCurrentItems = () => {
    if (activeTab === 'all') return shopItems.all;
    if (activeTab === 'discounts') return shopItems.discounts;
    if (activeTab === 'cosmetics') return shopItems.cosmetics;
    if (activeTab === 'services') return shopItems.services;
    return shopItems.all;
  };

  const handlePurchase = (item) => {
    if (coins >= item.price) {
      setCoins(coins - item.price);
      setPurchases([...purchases, { ...item, purchaseDate: new Date() }]);
      alert(`✅ Вы купили "${item.name}" за ${item.price} коинов!`);
    } else {
      alert(`❌ Недостаточно коинов! Нужно ${item.price}, у вас ${coins}`);
    }
  };

  return (
    <div className="shop-container">
      <div className="shop-header">
        <h1>🎁 Магазин</h1>
        <div className="coins-display">
          <span className="coin-icon">🪙</span>
          <span className="coin-amount">{coins}</span>
        </div>
      </div>

      <div className="shop-tabs">
        <button className={activeTab === 'all' ? 'tab-active' : ''} onClick={() => setActiveTab('all')}>
          Все товары
        </button>
        <button className={activeTab === 'discounts' ? 'tab-active' : ''} onClick={() => setActiveTab('discounts')}>
          Скидки
        </button>
        <button className={activeTab === 'cosmetics' ? 'tab-active' : ''} onClick={() => setActiveTab('cosmetics')}>
          Кастомизация
        </button>
        <button className={activeTab === 'services' ? 'tab-active' : ''} onClick={() => setActiveTab('services')}>
          Услуги
        </button>
      </div>

      <div className="shop-items-grid">
        {getCurrentItems().map(item => (
          <div key={item.id} className="shop-item">
            <div className="item-icon">{item.icon}</div>
            <div className="item-info">
              <h3>{item.name}</h3>
              <p>{item.description}</p>
              <div className="item-price">
                <span className="price-coins">🪙 {item.price}</span>
                <button 
                  className="buy-button"
                  onClick={() => handlePurchase(item)}
                  disabled={coins < item.price}
                >
                  Купить
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>

      {purchases.length > 0 && (
        <div className="purchases-history">
          <h2>📦 Мои покупки</h2>
          <div className="purchases-list">
            {purchases.map((purchase, idx) => (
              <div key={idx} className="purchase-item">
                <span>{purchase.icon}</span>
                <span>{purchase.name}</span>
                <span>-🪙 {purchase.price}</span>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default Shop;