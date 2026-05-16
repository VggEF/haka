import React, { useState, useEffect } from 'react';
import './Deadlines.css';

// Демо-дедлайны
const DEMO_DEADLINES = [
  { id: 1, title: "Лабораторная работа №3", subject: "Программирование", date: "2026-05-20", priority: "high", status: "pending" },
  { id: 2, title: "Контрольная работа", subject: "Математический анализ", date: "2026-05-25", priority: "high", status: "pending" },
  { id: 3, title: "Курсовая работа", subject: "Базы данных", date: "2026-06-01", priority: "medium", status: "pending" },
  { id: 4, title: "Реферат", subject: "Философия", date: "2026-05-18", priority: "low", status: "pending" },
  { id: 5, title: "Презентация проекта", subject: "Web-технологии", date: "2026-05-22", priority: "high", status: "pending" }
];

function Deadlines() {
  const [deadlines, setDeadlines] = useState([]);
  const [showAddForm, setShowAddForm] = useState(false);
  const [newDeadline, setNewDeadline] = useState({ title: "", subject: "", date: "", priority: "medium" });

  useEffect(() => {
    const saved = localStorage.getItem('deadlines');
    if (saved) {
      setDeadlines(JSON.parse(saved));
    } else {
      setDeadlines(DEMO_DEADLINES);
    }
  }, []);

  useEffect(() => {
    localStorage.setItem('deadlines', JSON.stringify(deadlines));
  }, [deadlines]);

  const addDeadline = () => {
    if (!newDeadline.title || !newDeadline.subject || !newDeadline.date) return;
    
    const newId = Math.max(...deadlines.map(d => d.id), 0) + 1;
    setDeadlines([...deadlines, { ...newDeadline, id: newId, status: "pending" }]);
    setNewDeadline({ title: "", subject: "", date: "", priority: "medium" });
    setShowAddForm(false);
  };

  const completeDeadline = (id) => {
    setDeadlines(deadlines.map(d => 
      d.id === id ? { ...d, status: "completed" } : d
    ));
    alert("✅ Отлично! Задание выполнено!");
  };

  const deleteDeadline = (id) => {
    setDeadlines(deadlines.filter(d => d.id !== id));
  };

  const getPriorityColor = (priority) => {
    switch(priority) {
      case 'high': return '#f44336';
      case 'medium': return '#ff9800';
      case 'low': return '#4caf50';
      default: return '#999';
    }
  };

  const getPriorityName = (priority) => {
    switch(priority) {
      case 'high': return 'Высокий';
      case 'medium': return 'Средний';
      case 'low': return 'Низкий';
      default: return '—';
    }
  };

  const formatDate = (dateStr) => {
    const date = new Date(dateStr);
    const today = new Date();
    const diffDays = Math.ceil((date - today) / (1000 * 60 * 60 * 24));
    
    if (diffDays < 0) return `Просрочено на ${-diffDays} дн.`;
    if (diffDays === 0) return "Сегодня!";
    if (diffDays === 1) return "Завтра";
    return `Через ${diffDays} дн.`;
  };

  const sortedDeadlines = [...deadlines].sort((a, b) => new Date(a.date) - new Date(b.date));

  return (
    <div className="deadlines-container">
      <h2 className="deadlines-title">⏰ Таймлайн дедлайнов</h2>
      
      <button className="add-deadline-btn" onClick={() => setShowAddForm(true)}>
        + Добавить дедлайн
      </button>

      {showAddForm && (
        <div className="add-deadline-form">
          <input type="text" placeholder="Название" value={newDeadline.title} onChange={(e) => setNewDeadline({...newDeadline, title: e.target.value})} />
          <input type="text" placeholder="Предмет" value={newDeadline.subject} onChange={(e) => setNewDeadline({...newDeadline, subject: e.target.value})} />
          <input type="date" value={newDeadline.date} onChange={(e) => setNewDeadline({...newDeadline, date: e.target.value})} />
          <select value={newDeadline.priority} onChange={(e) => setNewDeadline({...newDeadline, priority: e.target.value})}>
            <option value="high">Высокий</option>
            <option value="medium">Средний</option>
            <option value="low">Низкий</option>
          </select>
          <button onClick={addDeadline}>Сохранить</button>
          <button onClick={() => setShowAddForm(false)}>Отмена</button>
        </div>
      )}

      <div className="timeline">
        {sortedDeadlines.map((deadline) => (
          <div key={deadline.id} className={`timeline-item ${deadline.status === 'completed' ? 'completed' : ''}`}>
            <div className="timeline-dot" style={{ background: getPriorityColor(deadline.priority) }}></div>
            <div className="timeline-content">
              <div className="timeline-header">
                <h3>{deadline.title}</h3>
                <span className="timeline-date" style={{ color: getPriorityColor(deadline.priority) }}>
                  {formatDate(deadline.date)}
                </span>
              </div>
              <div className="timeline-meta">
                <span>📚 {deadline.subject}</span>
                <span>⚠️ {getPriorityName(deadline.priority)} приоритет</span>
              </div>
              <div className="timeline-actions">
                {deadline.status !== 'completed' && (
                  <button onClick={() => completeDeadline(deadline.id)}>✅ Выполнено</button>
                )}
                <button onClick={() => deleteDeadline(deadline.id)}>🗑️ Удалить</button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Deadlines;