import React, { useState, useEffect } from 'react';
import './AdminAchievements.css';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const getToken = () => localStorage.getItem('token');

function AdminAchievements() {
  const [achievements, setAchievements] = useState([]);
  const [students, setStudents] = useState([]);
  const [selectedStudent, setSelectedStudent] = useState("");
  const [selectedGroup, setSelectedGroup] = useState("");
  const [selectedAchievement, setSelectedAchievement] = useState("");
  const [loading, setLoading] = useState(false);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [newAchievement, setNewAchievement] = useState({
    title: '',
    description: '',
    xp: 0,
    icon: '🏆',
    rarity: 'common',
    category: 'Общая'
  });

  // Загрузка ачивок с сервера
  const fetchAchievements = async () => {
    console.log('🔄 Загрузка ачивок...');
    console.log('🔗 URL:', `${API_BASE_URL}/api/achievements`);
    
    try {
      const response = await fetch(`${API_BASE_URL}/api/achievements`);
      console.log('📡 Статус ответа:', response.status);
      
      const data = await response.json();
      console.log('📦 Полученные данные:', data);
      console.log('📦 Тип данных:', typeof data);
      console.log('📦 Это массив?', Array.isArray(data));
      
      if (Array.isArray(data)) {
        setAchievements(data);
        console.log('✅ Установлено ачивок:', data.length);
      } else if (data.data && Array.isArray(data.data)) {
        setAchievements(data.data);
        console.log('✅ Установлено ачивок из data.data:', data.data.length);
      } else {
        console.error('❌ Неожиданный формат данных:', data);
        setAchievements([]);
      }
    } catch (error) {
      console.error('❌ Ошибка загрузки ачивок:', error);
    }
  };

  // Загрузка студентов
  const fetchStudents = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/api/admin/users`, {
        headers: { 'Authorization': `Bearer ${getToken()}` }
      });
      const data = await response.json();
      setStudents(data.filter(u => u.role === 'student'));
    } catch (error) {
      console.error('Ошибка загрузки студентов:', error);
    }
  };

  useEffect(() => {
    fetchAchievements();
    fetchStudents();
  }, []);

  const createAchievement = async () => {
    console.log('🚀 НАЧАЛО СОЗДАНИЯ АЧИВКИ');
    console.log('📝 Данные ачивки:', newAchievement);
    console.log('🔗 URL:', `${API_BASE_URL}/api/admin/achievements`);
    console.log('🔑 Токен:', getToken());

    if (!newAchievement.title) {
      alert('Введите название ачивки');
      return;
    }

    setLoading(true);
    try {
      const response = await fetch(`${API_BASE_URL}/api/admin/achievements`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}`
        },
        body: JSON.stringify(newAchievement)
      });

      console.log('📡 Статус ответа:', response.status);
      
      const data = await response.json();
      console.log('📦 Ответ сервера:', data);

      if (response.ok) {
        console.log('✅ Ачивка создана успешно!');
        alert(`Ачивка "${newAchievement.title}" добавлена!`);
        setShowCreateModal(false);
        setNewAchievement({
          title: '',
          description: '',
          xp: 0,
          icon: '🏆',
          rarity: 'common',
          category: 'Общая'
        });
        await fetchAchievements(); // Обновляем список
      } else {
        console.error('❌ Ошибка сервера:', data);
        alert(`Ошибка: ${data.error || 'Неизвестная ошибка'}`);
      }
    } catch (error) {
      console.error('❌ Ошибка при создании:', error);
      alert('Ошибка при создании ачивки');
    } finally {
      setLoading(false);
    }
  };

  // Выдача ачивки студенту
  const awardToStudent = async () => {
    if (!selectedStudent || !selectedAchievement) {
      alert("Выберите студента и ачивку");
      return;
    }

    setLoading(true);
    try {
      const response = await fetch(`${API_BASE_URL}/api/achievements/award`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}`
        },
        body: JSON.stringify({
          user_id: parseInt(selectedStudent),
          achievement_id: parseInt(selectedAchievement)
        })
      });

      if (response.ok) {
        alert(`Ачивка выдана студенту!`);
        setSelectedStudent("");
        setSelectedAchievement("");
      } else {
        const error = await response.json();
        alert(`Ошибка: ${error.error}`);
      }
    } catch (error) {
      console.error('Ошибка:', error);
      alert('Ошибка при выдаче ачивки');
    } finally {
      setLoading(false);
    }
  };

  // Выдача ачивки группе
  const awardToGroup = async () => {
    if (!selectedGroup || !selectedAchievement) {
      alert("Выберите группу и ачивку");
      return;
    }

    setLoading(true);
    try {
      const response = await fetch(`${API_BASE_URL}/api/achievements/award/group`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${getToken()}`
        },
        body: JSON.stringify({
          group_name: selectedGroup,
          achievement_id: parseInt(selectedAchievement)
        })
      });

      if (response.ok) {
        alert(`Ачивка выдана группе!`);
        setSelectedGroup("");
        setSelectedAchievement("");
      } else {
        const error = await response.json();
        alert(`Ошибка: ${error.error}`);
      }
    } catch (error) {
      console.error('Ошибка:', error);
      alert('Ошибка при выдаче ачивки');
    } finally {
      setLoading(false);
    }
  };

  // Удаление ачивки
  const deleteAchievement = async (id, title) => {
    if (!window.confirm(`Удалить ачивку "${title}"?`)) return;

    try {
      const response = await fetch(`${API_BASE_URL}/api/achievements/${id}`, {
        method: 'DELETE',
        headers: { 'Authorization': `Bearer ${getToken()}` }
      });

      if (response.ok) {
        alert('Ачивка удалена');
        await fetchAchievements();
      } else {
        alert('Ошибка при удалении');
      }
    } catch (error) {
      console.error('Ошибка:', error);
      alert('Ошибка при удалении');
    }
  };

  const groups = [...new Set(students.map(s => s.group || "Без группы"))];

  return (
    <div className="admin-achievements">
      <h2>🏆 Управление ачивками</h2>
      
      <button className="add-btn" onClick={() => setShowCreateModal(true)} disabled={loading}>
        + Добавить ачивку
      </button>

      {/* Форма выдачи студенту */}
      <div className="award-section">
        <h3>🎯 Выдать ачивку студенту</h3>
        <div className="award-form">
          <select value={selectedStudent} onChange={(e) => setSelectedStudent(e.target.value)}>
            <option value="">Выберите студента</option>
            {students.map(s => <option key={s.id} value={s.id}>{s.name}</option>)}
          </select>
          <select value={selectedAchievement} onChange={(e) => setSelectedAchievement(e.target.value)}>
            <option value="">Выберите ачивку</option>
            {achievements.map(a => <option key={a.id} value={a.id}>{a.icon} {a.title} (+{a.xp} XP)</option>)}
          </select>
          <button onClick={awardToStudent} disabled={loading}>Выдать</button>
        </div>
      </div>

      {/* Форма выдачи группе */}
      <div className="award-section">
        <h3>👥 Выдать ачивку группе</h3>
        <div className="award-form">
          <select value={selectedGroup} onChange={(e) => setSelectedGroup(e.target.value)}>
            <option value="">Выберите группу</option>
            {groups.map(g => <option key={g} value={g}>{g}</option>)}
          </select>
          <select value={selectedAchievement} onChange={(e) => setSelectedAchievement(e.target.value)}>
            <option value="">Выберите ачивку</option>
            {achievements.map(a => <option key={a.id} value={a.id}>{a.icon} {a.title} (+{a.xp} XP)</option>)}
          </select>
          <button onClick={awardToGroup} disabled={loading}>Выдать группе</button>
        </div>
      </div>

      {/* Список ачивок */}
      <div className="achievements-list">
        <h3>📋 Список ачивок</h3>
        {achievements.length === 0 ? (
          <div className="empty-state">Ачивок пока нет</div>
        ) : (
          achievements.map(a => (
            <div key={a.id} className="achievement-card">
              <div className="achievement-icon">{a.icon}</div>
              <div className="achievement-info">
                <div className="achievement-title">{a.title}</div>
                <div className="achievement-desc">{a.description}</div>
                <div className="achievement-meta">
                  <span>+{a.xp} XP</span>
                  <span>{a.rarity}</span>
                  <span>{a.category}</span>
                </div>
              </div>
              <button className="delete-btn" onClick={() => deleteAchievement(a.id, a.title)}>
                🗑️
              </button>
            </div>
          ))
        )}
      </div>

      {/* Модальное окно создания */}
      {showCreateModal && (
        <div className="modal-overlay" onClick={() => setShowCreateModal(false)}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            <h3>➕ Создать ачивку</h3>
            <input
              type="text"
              placeholder="Название"
              value={newAchievement.title}
              onChange={(e) => setNewAchievement({...newAchievement, title: e.target.value})}
            />
            <textarea
              placeholder="Описание"
              value={newAchievement.description}
              onChange={(e) => setNewAchievement({...newAchievement, description: e.target.value})}
            />
            <input
              type="number"
              placeholder="XP награда"
              value={newAchievement.xp}
              onChange={(e) => setNewAchievement({...newAchievement, xp: parseInt(e.target.value) || 0})}
            />
            <input
              type="text"
              placeholder="Иконка (🏆)"
              value={newAchievement.icon}
              onChange={(e) => setNewAchievement({...newAchievement, icon: e.target.value})}
            />
            <select value={newAchievement.rarity} onChange={(e) => setNewAchievement({...newAchievement, rarity: e.target.value})}>
              <option value="common">Обычная</option>
              <option value="rare">Редкая</option>
              <option value="epic">Эпическая</option>
              <option value="legendary">Легендарная</option>
            </select>
            <input
              type="text"
              placeholder="Категория"
              value={newAchievement.category}
              onChange={(e) => setNewAchievement({...newAchievement, category: e.target.value})}
            />
            <div className="modal-buttons">
              <button onClick={createAchievement} disabled={loading}>Создать</button>
              <button onClick={() => setShowCreateModal(false)}>Отмена</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default AdminAchievements;