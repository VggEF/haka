import React, { useState, useEffect } from 'react';
import './Planner.css';

const HABITS = [
  { id: 1, name: "📖 Учиться 2 часа", xp: 20, streak: 0, lastDone: null },
  { id: 2, name: "💪 Сделать зарядку", xp: 10, streak: 0, lastDone: null },
  { id: 3, name: "📝 Записать конспект", xp: 15, streak: 0, lastDone: null },
  { id: 4, name: "🧠 Повторить материал", xp: 25, streak: 0, lastDone: null }
];

function Planner() {
  const [habits, setHabits] = useState([]);
  const [totalXP, setTotalXP] = useState(0);
  const [streakBonus, setStreakBonus] = useState(0);

  useEffect(() => {
    const savedHabits = localStorage.getItem('habits');
    const savedXP = localStorage.getItem('plannerXP');
    if (savedHabits) setHabits(JSON.parse(savedHabits));
    else setHabits(HABITS);
    if (savedXP) setTotalXP(parseInt(savedXP));
  }, []);

  const saveHabits = (newHabits) => {
    setHabits(newHabits);
    localStorage.setItem('habits', JSON.stringify(newHabits));
  };

  const toggleHabit = (habitId) => {
    const today = new Date().toDateString();
    const habit = habits.find(h => h.id === habitId);
    
    if (habit.lastDone === today) {
      alert("Сегодня вы уже отметили эту привычку!");
      return;
    }
    
    let newStreak = habit.streak + 1;
    let bonus = 0;
    
    // Бонус за серию
    if (newStreak === 7) {
      bonus = 50;
      alert("🔥🔥🔥 7 дней подряд! +50 XP бонус!");
    } else if (newStreak === 30) {
      bonus = 200;
      alert("🎉🎉🎉 30 дней без пропусков! +200 XP бонус!");
    }
    
    const earnedXP = habit.xp + bonus;
    setTotalXP(totalXP + earnedXP);
    localStorage.setItem('plannerXP', totalXP + earnedXP);
    
    const updatedHabits = habits.map(h => 
      h.id === habitId ? { ...h, streak: newStreak, lastDone: today } : h
    );
    saveHabits(updatedHabits);
    
    alert(`✅ +${habit.xp} XP за "${habit.name}"!${bonus ? ` Бонус: +${bonus} XP!` : ""}`);
  };

  return (
    <div className="planner-container">
      <h2>📅 Планер-трекер привычек</h2>
      <div className="planner-stats">
        <div className="stat">⭐ Всего XP: {totalXP}</div>
      </div>
      <div className="habits-list">
        {habits.map(habit => (
          <div key={habit.id} className="habit-card">
            <div className="habit-info">
              <div className="habit-name">{habit.name}</div>
              <div className="habit-streak">🔥 Серия: {habit.streak} дней</div>
              <div className="habit-xp">🎯 +{habit.xp} XP за день</div>
            </div>
            <button className="habit-btn" onClick={() => toggleHabit(habit.id)}>✅ Отметить</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Planner;