import React, { useState, useEffect } from 'react';
import './Challenges.css';

const DEMO_CHALLENGES = [
  { 
    id: 1, 
    title: "🏃‍♂️ Кто решит больше задач за неделю", 
    status: "active", 
    participants: ["Ты", "Иван", "Петр"], 
    scores: { "Ты": 5, "Иван": 7, "Петр": 3 }, 
    endDate: "2026-05-20" 
  },
  { 
    id: 2, 
    title: "📚 Кто получит больше отличных оценок", 
    status: "active", 
    participants: ["Ты", "Анна", "Мария"], 
    scores: { "Ты": 2, "Анна": 4, "Мария": 1 }, 
    endDate: "2026-05-25" 
  },
  { 
    id: 3, 
    title: "💻 Лучший проект по программированию", 
    status: "active", 
    participants: ["Ты", "Денис", "Ольга"], 
    scores: { "Ты": 85, "Денис": 92, "Ольга": 78 }, 
    endDate: "2026-06-01" 
  }
];

function Challenges() {
  const [challenges, setChallenges] = useState([]);
  const [newChallenge, setNewChallenge] = useState({ title: "", participants: "", endDate: "" });
  const [showCreate, setShowCreate] = useState(false);

  useEffect(() => {
    const saved = localStorage.getItem('challenges');
    if (saved) {
      setChallenges(JSON.parse(saved));
    } else {
      setChallenges(DEMO_CHALLENGES);
    }
  }, []);

  useEffect(() => {
    localStorage.setItem('challenges', JSON.stringify(challenges));
  }, [challenges]);

  const createChallenge = () => {
    if (!newChallenge.title || !newChallenge.participants) {
      alert("Заполните название и участников");
      return;
    }
    
    const participantsList = ["Ты", ...newChallenge.participants.split(',').map(p => p.trim())];
    const scores = {};
    participantsList.forEach(p => scores[p] = 0);
    
    const newId = Math.max(...challenges.map(c => c.id), 0) + 1;
    
    setChallenges([...challenges, {
      id: newId,
      title: newChallenge.title,
      status: "active",
      participants: participantsList,
      scores: scores,
      endDate: newChallenge.endDate || "2026-06-30"
    }]);
    
    setNewChallenge({ title: "", participants: "", endDate: "" });
    setShowCreate(false);
    alert("✅ Вызов создан!");
  };

  const updateScore = (challengeId, participant, delta) => {
    setChallenges(challenges.map(c => {
      if (c.id === challengeId) {
        const newScores = { ...c.scores };
        newScores[participant] = (newScores[participant] || 0) + delta;
        return { ...c, scores: newScores };
      }
      return c;
    }));
  };

  const deleteChallenge = (challengeId) => {
    if (window.confirm("Удалить вызов?")) {
      setChallenges(challenges.filter(c => c.id !== challengeId));
    }
  };

  const getWinner = (scores) => {
    const entries = Object.entries(scores);
    if (entries.length === 0) return null;
    return entries.reduce((a, b) => a[1] > b[1] ? a : b);
  };

  const formatDate = (dateStr) => {
    const date = new Date(dateStr);
    return date.toLocaleDateString('ru-RU');
  };

  return (
    <div className="challenges-container">
      <h2 className="challenges-title">🎯 Вызовы с друзьями</h2>
      
      <button className="create-challenge-btn" onClick={() => setShowCreate(true)}>
        + Создать вызов
      </button>

      {showCreate && (
        <div className="create-challenge-form">
          <input 
            type="text" 
            placeholder="Название вызова" 
            value={newChallenge.title} 
            onChange={(e) => setNewChallenge({...newChallenge, title: e.target.value})}
          />
          <input 
            type="text" 
            placeholder="Участники (через запятую, кроме тебя)" 
            value={newChallenge.participants} 
            onChange={(e) => setNewChallenge({...newChallenge, participants: e.target.value})}
          />
          <input 
            type="date" 
            value={newChallenge.endDate} 
            onChange={(e) => setNewChallenge({...newChallenge, endDate: e.target.value})}
          />
          <div className="form-buttons">
            <button className="submit-btn" onClick={createChallenge}>Создать</button>
            <button className="cancel-btn" onClick={() => setShowCreate(false)}>Отмена</button>
          </div>
        </div>
      )}

      <div className="challenges-list">
        {challenges.length === 0 ? (
          <div className="no-challenges">
            <span>🎯</span>
            <p>Нет активных вызовов. Создай свой первый вызов!</p>
          </div>
        ) : (
          challenges.map(challenge => {
            const sorted = Object.entries(challenge.scores).sort((a,b) => b[1] - a[1]);
            const winner = getWinner(challenge.scores);
            
            return (
              <div key={challenge.id} className="challenge-card">
                <div className="challenge-header">
                  <h3>{challenge.title}</h3>
                  <button className="delete-btn" onClick={() => deleteChallenge(challenge.id)}>🗑️</button>
                </div>
                <div className="challenge-meta">
                  <span>📅 До {formatDate(challenge.endDate)}</span>
                </div>
                
                <div className="challenge-leaderboard">
                  <div className="leaderboard-header">
                    <span>Участник</span>
                    <span>Очки</span>
                    <span>Действие</span>
                  </div>
                  {sorted.map(([name, score]) => (
                    <div key={name} className={`leaderboard-item ${name === winner?.[0] ? 'winner' : ''}`}>
                      <span>{name === winner?.[0] && '🏆 '}{name}</span>
                      <span className="score-value">{score}</span>
                      {name === "Ты" && (
                        <div className="score-actions">
                          <button onClick={() => updateScore(challenge.id, "Ты", 1)}>+1</button>
                          <button onClick={() => updateScore(challenge.id, "Ты", -1)}>-1</button>
                        </div>
                      )}
                      {name !== "Ты" && <span></span>}
                    </div>
                  ))}
                </div>
                
                {winner?.[0] === "Ты" && winner?.[1] > 0 && (
                  <div className="winner-badge">
                    🏆 Вы лидируете! Так держать!
                  </div>
                )}
              </div>
            );
          })
        )}
      </div>
    </div>
  );
}

export default Challenges;