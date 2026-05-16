import React, { useState, useEffect } from 'react';
import './AdminChallenges.css';

function AdminChallenges() {
  const [challenges, setChallenges] = useState([]);

  useEffect(() => {
    const saved = localStorage.getItem('challenges');
    if (saved) setChallenges(JSON.parse(saved));
  }, []);

  const deleteChallenge = (id) => {
    if (window.confirm("Удалить вызов? Он может нарушать правила университета.")) {
      const updated = challenges.filter(c => c.id !== id);
      setChallenges(updated);
      localStorage.setItem('challenges', JSON.stringify(updated));
      alert("Вызов удален");
    }
  };

  return (
    <div className="admin-page">
      <h2>🎯 Модерация вызовов</h2>
      
      <div className="items-list">
        {challenges.length === 0 && <div className="empty-state">Активных вызовов нет</div>}
        {challenges.map(ch => (
          <div key={ch.id} className="item-card">
            <div>🎯 <strong>{ch.title}</strong></div>
            <div>👥 Участники: {ch.participants?.join(', ')}</div>
            <div>📅 До {ch.endDate}</div>
            <button className="danger-btn" onClick={() => deleteChallenge(ch.id)}>🚫 Удалить</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default AdminChallenges;