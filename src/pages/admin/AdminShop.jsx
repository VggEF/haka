import React, { useState } from 'react';
import './AdminShop.css';

const AdminShop = () => {
  const [shopItems, setShopItems] = useState([
    { id: 1, name: 'Скидка на экзамен 10%', price: 500, icon: '🎓', type: 'discount', available: true },
    { id: 2, name: 'Освобождение от дедлайна', price: 1000, icon: '⏰', type: 'deadline', available: true },
    { id: 3, name: 'Секретная тема', price: 300, icon: '🔓', type: 'content', available: true },
  ]);

  const [newItem, setNewItem] = useState({ name: '', price: '', icon: '🎁', type: 'other' });
  const [editingId, setEditingId] = useState(null);

  const handleAddItem = () => {
    if (newItem.name && newItem.price) {
      setShopItems([...shopItems, { 
        id: Date.now(), 
        ...newItem, 
        price: parseInt(newItem.price),
        available: true 
      }]);
      setNewItem({ name: '', price: '', icon: '🎁', type: 'other' });
    }
  };

  const handleDeleteItem = (id) => {
    setShopItems(shopItems.filter(item => item.id !== id));
  };

  const handleToggleAvailable = (id) => {
    setShopItems(shopItems.map(item => 
      item.id === id ? { ...item, available: !item.available } : item
    ));
  };

  const handleUpdatePrice = (id, newPrice) => {
    setShopItems(shopItems.map(item => 
      item.id === id ? { ...item, price: parseInt(newPrice) } : item
    ));
    setEditingId(null);
  };

  return (
    <div className="admin-shop-container">
      <h1>🏪 Управление магазином</h1>
      
      <div className="admin-stats">
        <div className="stat-card">
          <h3>📦 Всего товаров</h3>
          <p>{shopItems.length}</p>
        </div>
        <div className="stat-card">
          <h3>🟢 Активных</h3>
          <p>{shopItems.filter(i => i.available).length}</p>
        </div>
        <div className="stat-card">
          <h3>💰 Средняя цена</h3>
          <p>{(shopItems.reduce((sum, i) => sum + i.price, 0) / shopItems.length || 0).toFixed(0)}</p>
        </div>
      </div>

      {/* Добавление нового товара */}
      <div className="add-item-form">
        <h2>➕ Добавить товар</h2>
        <div className="form-group">
          <input 
            type="text" 
            placeholder="Название товара"
            value={newItem.name}
            onChange={(e) => setNewItem({...newItem, name: e.target.value})}
          />
          <input 
            type="number" 
            placeholder="Цена в коинах"
            value={newItem.price}
            onChange={(e) => setNewItem({...newItem, price: e.target.value})}
          />
          <input 
            type="text" 
            placeholder="Иконка (эмодзи)"
            value={newItem.icon}
            onChange={(e) => setNewItem({...newItem, icon: e.target.value})}
          />
          <select 
            value={newItem.type}
            onChange={(e) => setNewItem({...newItem, type: e.target.value})}
          >
            <option value="discount">Скидка</option>
            <option value="deadline">Дедлайн</option>
            <option value="content">Контент</option>
            <option value="avatar">Аватар</option>
            <option value="service">Услуга</option>
            <option value="other">Другое</option>
          </select>
          <button onClick={handleAddItem}>➕ Добавить</button>
        </div>
      </div>

      {/* Список товаров */}
      <div className="items-list">
        <h2>📋 Товары в магазине</h2>
        <table className="items-table">
          <thead>
            <tr>
              <th>Иконка</th>
              <th>Название</th>
              <th>Цена</th>
              <th>Тип</th>
              <th>Статус</th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            {shopItems.map(item => (
              <tr key={item.id}>
                <td className="item-icon">{item.icon}</td>
                <td>{item.name}</td>
                <td>
                  {editingId === item.id ? (
                    <input 
                      type="number" 
                      defaultValue={item.price}
                      onBlur={(e) => handleUpdatePrice(item.id, e.target.value)}
                      autoFocus
                    />
                  ) : (
                    <span onClick={() => setEditingId(item.id)} className="editable-price">
                      🪙 {item.price}
                    </span>
                  )}
                </td>
                <td>{item.type}</td>
                <td>
                  <button 
                    className={`status-btn ${item.available ? 'active' : 'inactive'}`}
                    onClick={() => handleToggleAvailable(item.id)}
                  >
                    {item.available ? 'Активен' : 'Скрыт'}
                  </button>
                </td>
                <td>
                  <button className="delete-btn" onClick={() => handleDeleteItem(item.id)}>
                    🗑️
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* Статистика продаж (заглушка) */}
      <div className="sales-stats">
        <h2>📊 Статистика продаж (тестовые данные)</h2>
        <div className="stats-grid">
          <div>Сегодня: 15 покупок</div>
          <div>За неделю: 127 покупок</div>
          <div>Всего продано: 2,450 коинов</div>
          <div>Популярный товар: Скидка на экзамен</div>
        </div>
      </div>
    </div>
  );
};

export default AdminShop;