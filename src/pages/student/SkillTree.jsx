import React, { useState, useEffect } from 'react';
import './SkillTree.css';

// Данные дерева навыков
const SKILL_TREE_DATA = {
  programming: {
    name: "💻 Программирование",
    icon: "💻",
    color: "#4CAF50",
    skills: [
      { id: "python_basic", name: "🐍 Python (базовый)", xp: 50, description: "Переменные, циклы, функции", required: null, xpRequired: 50 },
      { id: "python_advanced", name: "🐍 Python (продвинутый)", xp: 100, description: "ООП, исключения, декораторы", required: "python_basic", xpRequired: 150 },
      { id: "git", name: "📦 Git и GitHub", xp: 80, description: "Ветки, коммиты, PR", required: null, xpRequired: 80 },
      { id: "sql", name: "🗄️ SQL", xp: 75, description: "SELECT, JOIN, GROUP BY", required: null, xpRequired: 75 },
      { id: "react", name: "⚛️ React", xp: 120, description: "Компоненты, хуки, состояние", required: "python_basic", xpRequired: 200 },
      { id: "algorithms", name: "📊 Алгоритмы", xp: 150, description: "Сортировки, графы, деревья", required: "python_basic", xpRequired: 250 }
    ]
  },
  softskills: {
    name: "🎯 Soft Skills",
    icon: "🎯",
    color: "#2196F3",
    skills: [
      { id: "presentation", name: "🎤 Презентации", xp: 60, description: "Структура, дизайн, выступление", required: null, xpRequired: 60 },
      { id: "teamwork", name: "🤝 Работа в команде", xp: 70, description: "Коммуникация, эмпатия", required: null, xpRequired: 70 },
      { id: "english", name: "🇬🇧 Английский для IT", xp: 100, description: "Технический английский", required: null, xpRequired: 100 },
      { id: "public_speaking", name: "🗣️ Ораторское искусство", xp: 80, description: "Выступления перед аудиторией", required: "presentation", xpRequired: 140 }
    ]
  },
  career: {
    name: "🚀 Карьера",
    icon: "🚀",
    color: "#FF9800",
    skills: [
      { id: "resume", name: "📄 Резюме", xp: 50, description: "Составление CV", required: null, xpRequired: 50 },
      { id: "interview", name: "💼 Собеседование", xp: 80, description: "Подготовка к интервью", required: "resume", xpRequired: 130 },
      { id: "linkedin", name: "🔗 LinkedIn", xp: 40, description: "Оформление профиля", required: null, xpRequired: 40 },
      { id: "portfolio", name: "📁 Портфолио", xp: 90, description: "Сбор проектов", required: "resume", xpRequired: 140 }
    ]
  }
};

function SkillTree() {
  const [unlockedSkills, setUnlockedSkills] = useState([]);
  const [userXP, setUserXP] = useState(0);
  const [selectedSkill, setSelectedSkill] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [activeBranch, setActiveBranch] = useState('programming');

  useEffect(() => {
    const savedUnlocked = localStorage.getItem('unlockedSkills');
    const savedXP = localStorage.getItem('totalXP');
    
    if (savedUnlocked) setUnlockedSkills(JSON.parse(savedUnlocked));
    if (savedXP) setUserXP(parseInt(savedXP) || 0);
  }, []);

  const isUnlocked = (skillId) => unlockedSkills.includes(skillId);
  
  const canUnlock = (skill) => {
    if (isUnlocked(skill.id)) return false;
    if (skill.required && !unlockedSkills.includes(skill.required)) return false;
    if (userXP < skill.xpRequired) return false;
    return true;
  };

  const unlockSkill = (skill) => {
    if (!canUnlock(skill)) return;
    
    const newUnlocked = [...unlockedSkills, skill.id];
    setUnlockedSkills(newUnlocked);
    localStorage.setItem('unlockedSkills', JSON.stringify(newUnlocked));
    
    alert(`🎉 Навык "${skill.name}" открыт! +${skill.xp} XP к карме`);
    setSelectedSkill(null);
    setIsModalOpen(false);
  };

  const openModal = (skill) => {
    setSelectedSkill(skill);
    setIsModalOpen(true);
  };

  const currentBranch = SKILL_TREE_DATA[activeBranch];

  return (
    <div className="skilltree-container">
      <h2 className="skilltree-title">🌳 Дерево навыков</h2>
      <div className="skilltree-stats">
        <div className="stat-card">
          <span className="stat-icon">⭐</span>
          <span className="stat-value">{userXP}</span>
          <span className="stat-label">Всего XP</span>
        </div>
        <div className="stat-card">
          <span className="stat-icon">🔓</span>
          <span className="stat-value">{unlockedSkills.length}</span>
          <span className="stat-label">Навыков открыто</span>
        </div>
      </div>

      {/* Ветки */}
      <div className="skilltree-branches">
        {Object.entries(SKILL_TREE_DATA).map(([key, branch]) => (
          <button
            key={key}
            className={`branch-btn ${activeBranch === key ? 'active' : ''}`}
            onClick={() => setActiveBranch(key)}
            style={activeBranch === key ? { background: branch.color } : {}}
          >
            {branch.icon} {branch.name}
          </button>
        ))}
      </div>

      {/* Дерево навыков */}
      <div className="skilltree-tree">
        {currentBranch.skills.map((skill, index) => {
          const unlocked = isUnlocked(skill.id);
          const can = canUnlock(skill);
          
          return (
            <div key={skill.id} className="skill-node">
              <div 
                className={`skill-card ${unlocked ? 'unlocked' : ''} ${can ? 'can-unlock' : ''}`}
                onClick={() => openModal(skill)}
              >
                <div className="skill-icon">{skill.icon || '📚'}</div>
                <div className="skill-info">
                  <div className="skill-name">{skill.name}</div>
                  <div className="skill-xp">🔒 {skill.xpRequired} XP</div>
                </div>
                {unlocked && <div className="unlocked-badge">✅</div>}
                {can && !unlocked && <div className="can-unlock-badge">🔓</div>}
              </div>
              {index < currentBranch.skills.length - 1 && (
                <div className={`skill-connector ${unlocked ? 'active' : ''}`}></div>
              )}
            </div>
          );
        })}
      </div>

      {/* Модальное окно */}
      {isModalOpen && selectedSkill && (
        <div className="modal-overlay" onClick={() => setIsModalOpen(false)}>
          <div className="modal-container" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header" style={{ background: currentBranch.color }}>
              <div className="modal-icon">{selectedSkill.icon || '📚'}</div>
              <h3>{selectedSkill.name}</h3>
              <button className="modal-close" onClick={() => setIsModalOpen(false)}>✕</button>
            </div>
            <div className="modal-body">
              <p className="modal-description">{selectedSkill.description}</p>
              <div className="modal-details">
                <div className="detail-row">
                  <span>📊 Требуется XP:</span>
                  <span>{selectedSkill.xpRequired}</span>
                </div>
                <div className="detail-row">
                  <span>🏆 Награда:</span>
                  <span>+{selectedSkill.xp} XP</span>
                </div>
                {selectedSkill.required && (
                  <div className="detail-row">
                    <span>🔗 Требуется навык:</span>
                    <span>{SKILL_TREE_DATA[activeBranch].skills.find(s => s.id === selectedSkill.required)?.name}</span>
                  </div>
                )}
              </div>
              {!isUnlocked(selectedSkill.id) && (
                <button 
                  className={`unlock-btn ${canUnlock(selectedSkill) ? '' : 'disabled'}`}
                  onClick={() => unlockSkill(selectedSkill)}
                  disabled={!canUnlock(selectedSkill)}
                >
                  {canUnlock(selectedSkill) ? '🔓 Открыть навык' : '🔒 Недостаточно XP'}
                </button>
              )}
              {isUnlocked(selectedSkill.id) && (
                <div className="already-unlocked">✅ Навык уже открыт</div>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default SkillTree;