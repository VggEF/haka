import React from 'react';
import './Chats.css';

function Chats() {
  const chats = [
    { id: 1, name: '📚 Учебная группа ПМбо-23', message: 'Иван: Завтра пары по расписанию', time: '10:30' },
    { id: 2, name: '👨‍🏫 Преподаватели', message: 'Доцент: Дедлайн по лабе завтра', time: '09:15' },
    { id: 3, name: '🎉 Общий чат курса', message: 'Анна: Кто идет на пары?', time: 'Вчера' },
  ];

  return (
    <div className="chats-container">
      <h2 className="chats-title">💬 Чаты</h2>
      <div className="chats-list">
        {chats.map(chat => (
          <div key={chat.id} className="chat-item">
            <div className="chat-name">{chat.name}</div>
            <div className="chat-message">{chat.message}</div>
            <div className="chat-time">{chat.time}</div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Chats;