import React, { useState, useEffect } from 'react';
import './AdminGames.css';

function AdminGames() {
  const [games, setGames] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    icon: "🎮",
    xp: 10
  });

  useEffect(() => {
    const saved = localStorage.getItem('miniGames');
    if (saved) setGames(JSON.parse(saved));
  }, []);

  const saveToLocalStorage = (newGames) => {
    setGames(newGames);
    localStorage.setItem('miniGames', JSON.stringify(newGames));
  };

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const addGame = () => {
    if (!formData.title) {
      alert("Введите название игры");
      return;
    }
    const newGame = { id: Date.now(), ...formData };
    saveToLocalStorage([...games, newGame]);
    resetForm();
  };

  const updateGame = () => {
    const updated = games.map(g =>
      g.id === editingId ? { ...g, ...formData } : g
    );
    saveToLocalStorage(updated);
    resetForm();
  };

  const deleteGame = (id) => {
    if (window.confirm("Удалить игру?")) {
      saveToLocalStorage(games.filter(g => g.id !== id));
    }
  };

  const editGame = (game) => {
    setEditingId(game.id);
    setFormData(game);
    setShowModal(true);
  };

  const resetForm = () => {
    setFormData({ title: "", description: "", icon: "🎮", xp: 10 });
    setEditingId(null);
    setShowModal(false);
  };

  return (
    <div className="admin-games">
      <div className="games-header">
        <div className="games-title">
          <div className="games-title-icon">🎮</div>
          <h1>Управление мини-играми</h1>
        </div>
        <div className="games-subtitle">Добавляйте игры для отдыха и получения XP</div>
      </div>

      <button className="add-game-btn" onClick={() => setShowModal(true)}>
        ➕ Добавить игру
      </button>

      {/* Модальное окно */}
      {showModal && (
        <div className="modal-overlay" onClick={resetForm}>
          <div className="game-modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>{editingId ? '✏️ Редактировать игру' : '🎮 Новая игра'}</h3>
              <button className="modal-close" onClick={resetForm}>✕</button>
            </div>
            <div className="modal-body">
              <div className="form-group">
                <label>Название игры</label>
                <input type="text" name="title" placeholder="Название" value={formData.title} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>Иконка (эмодзи)</label>
                <input type="text" name="icon" placeholder="🎮" value={formData.icon} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>Описание</label>
                <textarea name="description" rows="3" placeholder="Краткое описание игры..." value={formData.description} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>XP за победу</label>
                <input type="number" name="xp" placeholder="10" value={formData.xp} onChange={handleInputChange} />
              </div>
              <div className="form-actions">
                <button className="submit-btn" onClick={editingId ? updateGame : addGame}>
                  {editingId ? '💾 Сохранить' : '📤 Добавить'}
                </button>
                <button className="cancel-btn" onClick={resetForm}>Отмена</button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Список игр */}
      <div className="games-grid">
        {games.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">🎮</div>
            <div className="empty-text">Нет игр. Добавьте первую!</div>
            <button className="empty-btn" onClick={() => setShowModal(true)}>➕ Добавить игру</button>
          </div>
        ) : (
          games.map(game => (
            <div key={game.id} className="game-card">
              <div className="game-image">
                <span>{game.icon || '🎮'}</span>
              </div>
              <div className="game-content">
                <h3 className="game-title">{game.title}</h3>
                <p className="game-description">{game.description || "Увлекательная мини-игра"}</p>
                <div className="game-footer">
                  <div className="game-xp">🎯 +{game.xp} XP за победу</div>
                  <div className="game-actions">
                    <button className="edit-game" onClick={() => editGame(game)}>
                      ✏️ Ред.
                    </button>
                    <button className="delete-game" onClick={() => deleteGame(game.id)}>
                      🗑️ Удалить
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default AdminGames;