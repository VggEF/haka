import React, { useState, useEffect } from 'react';
import { getCoins, addCoins, spendCoins, SHOP_ITEMS } from '../utils/notifications';
import './CoinsWidget.css';

function CoinsWidget() {
  const [coins, setCoins] = useState(0);
  const [showShop, setShowShop] = useState(false);
  const [message, setMessage] = useState('');

  useEffect(() => {
    setCoins(getCoins());
  }, []);

  const handleBuy = (item) => {
    if (spendCoins(item.price)) {
      setCoins(getCoins());
      setMessage(`✅ Куплено: ${item.name}!`);
      setTimeout(() => setMessage(''), 3000);
    } else {
      setMessage(`❌ Не хватает коинов! Нужно ${item.price}`);
      setTimeout(() => setMessage(''), 3000);
    }
  };

  return (
    <div className="coins-widget">
      <div className="coins-display" onClick={() => setShowShop(!showShop)}>
        <span className="coin-icon">🪙</span>
        <span className="coin-amount">{coins}</span>
        <span className="coin-label">Коинов</span>
      </div>
      
      {showShop && (
        <div className="shop-popup">
          <h4>Магазин</h4>
          {SHOP_ITEMS.map(item => (
            <div key={item.id} className="shop-item" onClick={() => handleBuy(item)}>
              <span>{item.icon}</span>
              <span>{item.name}</span>
              <span>{item.price}🪙</span>
            </div>
          ))}
        </div>
      )}
      
      {message && <div className="coins-message">{message}</div>}
    </div>
  );
}

export default CoinsWidget;