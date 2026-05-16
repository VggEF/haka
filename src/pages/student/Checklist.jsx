import React, { useState, useEffect } from 'react';
import './Checklist.css';

const CHECKLIST_ITEMS = [
  { id: 1, title: "🎓 Получить студенческий билет", xp: 50, category: "Документы", isGameified: true },
  { id: 2, title: "📚 Записаться в библиотеку", xp: 30, category: "Инфраструктура", isGameified: true },
  { id: 3, title: "💬 Найти свою группу в чате", xp: 20, category: "Общение", isGameified: true },
  { id: 4, title: "📱 Скачать все приложения (Moodle, Teams)", xp: 40, category: "Техника", isGameified: true },
  { id: 5, title: "🏛️ Найти деканат", xp: 15, category: "Ориентация", isGameified: true },
  { id: 6, title: "🍽️ Найти столовую", xp: 10, category: "Ориентация", isGameified: true },
  { id: 7, title: "💳 Получить карту для прохода", xp: 25, category: "Документы", isGameified: true },
  { id: 8, title: "📖 Купить учебники", xp: 20, category: "Учеба", isGameified: true },
  { id: 9, title: "👨‍🏫 Познакомиться с куратором", xp: 30, category: "Общение", isGameified: true },
  { id: 10, title: "🎯 Сдать первую лабораторную", xp: 100, category: "Учеба", isGameified: true },
];

function Checklist() {
  const [completedItems, setCompletedItems] = useState([]);
  const [totalXP, setTotalXP] = useState(0);
  const [selectedCategory, setSelectedCategory] = useState('all');

  useEffect(() => {
    const saved = localStorage.getItem('checklistCompleted');
    const savedXP = localStorage.getItem('checklistXP');
    
    if (saved) setCompletedItems(JSON.parse(saved));
    if (savedXP) setTotalXP(parseInt(savedXP));
  }, []);

  const toggleItem = (item) => {
    if (completedItems.includes(item.id)) {
      const newCompleted = completedItems.filter(id => id !== item.id);
      setCompletedItems(newCompleted);
      setTotalXP(totalXP - item.xp);
      localStorage.setItem('checklistCompleted', JSON.stringify(newCompleted));
      localStorage.setItem('checklistXP', totalXP - item.xp);
    } else {
      const newCompleted = [...completedItems, item.id];
      setCompletedItems(newCompleted);
      setTotalXP(totalXP + item.xp);
      localStorage.setItem('checklistCompleted', JSON.stringify(newCompleted));
      localStorage.setItem('checklistXP', totalXP + item.xp);
      alert(`🎉 +${item.xp} XP! "${item.title}" выполнено!`);
    }
  };

  const categories = ['all', ...new Set(CHECKLIST_ITEMS.map(item => item.category))];
  
  const filteredItems = CHECKLIST_ITEMS.filter(item => 
    selectedCategory === 'all' || item.category === selectedCategory
  );

  const completedCount = completedItems.length;
  const totalCount = CHECKLIST_ITEMS.length;
  const progress = (completedCount / totalCount) * 100;

  return (
    <div className="checklist-container">
      <h2 className="checklist-title">🎯 Чек-лист первокурсника</h2>
      
      <div className="checklist-stats">
        <div className="checklist-progress">
          <div className="progress-bar">
            <div className="progress-fill" style={{ width: `${progress}%` }}></div>
          </div>
          <div className="progress-text">
            Выполнено {completedCount} из {totalCount} ({Math.round(progress)}%)
          </div>
        </div>
        <div className="xp-bonus">
          <span>⭐</span>
          <span>{totalXP} XP получено</span>
        </div>
      </div>

      <div className="checklist-categories">
        {categories.map(cat => (
          <button
            key={cat}
            className={`category-btn ${selectedCategory === cat ? 'active' : ''}`}
            onClick={() => setSelectedCategory(cat)}
          >
            {cat === 'all' ? 'Все' : cat}
          </button>
        ))}
      </div>

      <div className="checklist-items">
        {filteredItems.map(item => {
          const isCompleted = completedItems.includes(item.id);
          return (
            <div key={item.id} className={`checklist-item ${isCompleted ? 'completed' : ''}`} onClick={() => toggleItem(item)}>
              <div className="item-checkbox">{isCompleted ? '✅' : '⬜'}</div>
              <div className="item-info">
                <div className="item-title">{item.title}</div>
                <div className="item-meta">
                  <span className="item-category">{item.category}</span>
                  <span className="item-xp">+{item.xp} XP</span>
                </div>
              </div>
              {isCompleted && <div className="item-completed-badge">Готово!</div>}
            </div>
          );
        })}
      </div>

      {completedCount === totalCount && (
        <div className="checklist-complete">
          <div className="complete-emoji">🏆🎉🎓</div>
          <div className="complete-text">Поздравляем! Ты прошел чек-лист первокурсника!</div>
          <div className="complete-bonus">Бонус: +500 XP к карме студента</div>
        </div>
      )}
    </div>
  );
}

export default Checklist;