import React, { useState, useEffect } from 'react';
import './Achievements.css';

const API_BASE_URL = import.meta.env.API_BASE_URL || 'http://localhost:8080';

const getToken = () => localStorage.getItem('token');

function Achievements() {
  const [achievements, setAchievements] = useState([]);
  const [loading, setLoading] = useState(true);
  const [totalXP, setTotalXP] = useState(0);
  const [level, setLevel] = useState(1);
  const [selectedAchievement, setSelectedAchievement] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState('all');

  // Загрузка ачивок пользователя
  const fetchMyAchievements = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/api/achievements/me`, {
        headers: { 'Authorization': `Bearer ${getToken()}` }
      });
      const data = await response.json();
      setAchievements(data);
      
      const total = data.filter(a => a.is_unlocked).reduce((sum, a) => sum + a.xp, 0);
      setTotalXP(total);
      setLevel(Math.floor(total / 100) + 1);
    } catch (error) {
      console.error('Ошибка загрузки ачивок:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMyAchievements();
  }, []);

  const getFilteredAchievements = () => {
    if (activeTab === 'unlocked') {
      return achievements.filter(a => a.is_unlocked);
    }
    if (activeTab === 'locked') {
      return achievements.filter(a => !a.is_unlocked);
    }
    return achievements;
  };

  const getRarityColor = (rarity) => {
    switch (rarity) {
      case 'common': return '#9E9E9E';
      case 'rare': return '#2196F3';
      case 'epic': return '#9C27B0';
      case 'legendary': return '#FF9800';
      default: return '#9E9E9E';
    }
  };

  const getRarityName = (rarity) => {
    switch (rarity) {
      case 'common': return 'Обычная';
      case 'rare': return 'Редкая';
      case 'epic': return 'Эпическая';
      case 'legendary': return 'Легендарная';
      default: return 'Обычная';
    }
  };

  const getLevelTitle = () => {
    if (level >= 50) return "👑 Легендарный студент";
    if (level >= 30) return "💎 Магистр знаний";
    if (level >= 20) return "⭐ Эксперт";
    if (level >= 10) return "🚀 Продвинутый";
    if (level >= 5) return "📚 Ученик";
    return "🌱 Новичок";
  };

  const filteredAchievements = getFilteredAchievements();
  const completedCount = achievements.filter(a => a.is_unlocked).length;
  const totalCount = achievements.length;
  const currentLevelXP = totalXP - (level - 1) * 100;
  const progressToNextLevel = (currentLevelXP / 100) * 100;

  if (loading) {
    return <div className="loading-container">Загрузка ачивок...</div>;
  }

  return (
    <div className="achievements-container">
      {/* Шапка с уровнем */}
      <div className="achievements-header">
        <div className="level-card">
          <div className="level-icon">🏆</div>
          <div className="level-info">
            <div className="level-title">{getLevelTitle()}</div>
            <div className="level-number">Уровень {level}</div>
            <div className="level-progress">
              <div className="level-progress-bar">
                <div className="level-progress-fill" style={{ width: `${Math.min(progressToNextLevel, 100)}%` }}></div>
              </div>
              <div className="level-progress-text">{currentLevelXP} / 100 XP до {level + 1} уровня</div>
            </div>
          </div>
          <div className="xp-total">
            <span className="xp-icon">⭐</span>
            <span className="xp-value">{totalXP}</span>
            <span className="xp-label">всего XP</span>
          </div>
        </div>
      </div>

      {/* Вкладки */}
      <div className="achievements-tabs">
        <button className={`tab-btn ${activeTab === 'all' ? 'active' : ''}`} onClick={() => setActiveTab('all')}>
          📋 Все ({totalCount})
        </button>
        <button className={`tab-btn ${activeTab === 'unlocked' ? 'active' : ''}`} onClick={() => setActiveTab('unlocked')}>
          ✅ Полученные ({completedCount})
        </button>
        <button className={`tab-btn ${activeTab === 'locked' ? 'active' : ''}`} onClick={() => setActiveTab('locked')}>
          🔒 Неполученные ({totalCount - completedCount})
        </button>
      </div>

      {/* Список ачивок */}
      <div className="achievements-list">
        {filteredAchievements.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">🏆</div>
            <div className="empty-text">Ачивок пока нет</div>
          </div>
        ) : (
          filteredAchievements.map((achievement) => (
            <div 
              key={achievement.id} 
              className={`achievement-card ${achievement.is_unlocked ? 'completed' : ''}`}
              onClick={() => {
                setSelectedAchievement(achievement);
                setIsModalOpen(true);
              }}
            >
              <div className="achievement-icon" style={{ backgroundColor: getRarityColor(achievement.rarity) }}>
                {achievement.icon || '🏆'}
              </div>
              <div className="achievement-info">
                <div className="achievement-header">
                  <h3 className="achievement-title">{achievement.title}</h3>
                  <span className="achievement-xp" style={{ background: `${getRarityColor(achievement.rarity)}20`, color: getRarityColor(achievement.rarity) }}>
                    +{achievement.xp} XP
                  </span>
                </div>
                <p className="achievement-description">{achievement.description}</p>
                <div className="achievement-footer">
                  <span className="achievement-category">{achievement.category}</span>
                  <span className="achievement-rarity" style={{ color: getRarityColor(achievement.rarity) }}>
                    {getRarityName(achievement.rarity)}
                  </span>
                </div>
              </div>
              {achievement.is_unlocked && <div className="completed-badge">✅</div>}
            </div>
          ))
        )}
      </div>

      {/* Модальное окно */}
      {isModalOpen && selectedAchievement && (
        <div className="modal-overlay" onClick={() => setIsModalOpen(false)}>
          <div className="modal-container" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header" style={{ backgroundColor: getRarityColor(selectedAchievement.rarity) }}>
              <div className="modal-icon">{selectedAchievement.icon || '🏆'}</div>
              <h3>{selectedAchievement.title}</h3>
              <button className="modal-close" onClick={() => setIsModalOpen(false)}>✕</button>
            </div>
            <div className="modal-body">
              <div className="modal-description">{selectedAchievement.description}</div>
              <div className="modal-details">
                <div className="detail-row">📊 Награда: <strong>+{selectedAchievement.xp} XP</strong></div>
                <div className="detail-row">💎 Редкость: <strong style={{ color: getRarityColor(selectedAchievement.rarity) }}>{getRarityName(selectedAchievement.rarity)}</strong></div>
                <div className="detail-row">🏷️ Категория: <strong>{selectedAchievement.category}</strong></div>
                {selectedAchievement.condition && <div className="detail-row">🎯 Условие: <strong>{selectedAchievement.condition}</strong></div>}
              </div>
              {selectedAchievement.is_unlocked && (
                <div className="already-completed">
                  <div className="complete-icon">🏆</div>
                  <div>Ачивка получена!</div>
                </div>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Achievements;