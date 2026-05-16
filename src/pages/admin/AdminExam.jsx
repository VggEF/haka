import React, { useState, useEffect } from 'react';
import './AdminExam.css';

function AdminExam() {
  const [questions, setQuestions] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [formData, setFormData] = useState({
    question: "",
    answer: "",
    difficulty: "medium",
    subject: ""
  });

  useEffect(() => {
    const saved = localStorage.getItem('examQuestions');
    if (saved) setQuestions(JSON.parse(saved));
  }, []);

  const saveToLocalStorage = (newQuestions) => {
    setQuestions(newQuestions);
    localStorage.setItem('examQuestions', JSON.stringify(newQuestions));
  };

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const setDifficulty = (diff) => {
    setFormData({ ...formData, difficulty: diff });
  };

  const addQuestion = () => {
    if (!formData.question || !formData.answer) {
      alert("Заполните вопрос и ответ");
      return;
    }
    const newQuestion = { id: Date.now(), ...formData };
    saveToLocalStorage([...questions, newQuestion]);
    resetForm();
  };

  const updateQuestion = () => {
    const updated = questions.map(q =>
      q.id === editingId ? { ...q, ...formData } : q
    );
    saveToLocalStorage(updated);
    resetForm();
  };

  const deleteQuestion = (id) => {
    if (window.confirm("Удалить вопрос?")) {
      saveToLocalStorage(questions.filter(q => q.id !== id));
    }
  };

  const editQuestion = (q) => {
    setEditingId(q.id);
    setFormData(q);
    setShowModal(true);
  };

  const resetForm = () => {
    setFormData({ question: "", answer: "", difficulty: "medium", subject: "" });
    setEditingId(null);
    setShowModal(false);
  };

  const stats = {
    total: questions.length,
    easy: questions.filter(q => q.difficulty === 'easy').length,
    medium: questions.filter(q => q.difficulty === 'medium').length,
    hard: questions.filter(q => q.difficulty === 'hard').length
  };

  const getDifficultyText = (diff) => {
    const texts = { easy: 'Легкий', medium: 'Средний', hard: 'Сложный' };
    return texts[diff] || 'Средний';
  };

  return (
    <div className="admin-exam">
      <div className="exam-header">
        <div className="exam-title">
          <div className="exam-title-icon">🎓</div>
          <h1>Управление экзаменом</h1>
        </div>
        <div className="exam-subtitle">Добавляйте вопросы для подготовки к экзаменам</div>
      </div>

      {/* Статистика */}
      <div className="exam-stats">
        <div className="stat-card">
          <div className="stat-icon">📚</div>
          <div className="stat-info">
            <h3>{stats.total}</h3>
            <p>Всего вопросов</p>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon">🟢</div>
          <div className="stat-info">
            <h3>{stats.easy}</h3>
            <p>Легких</p>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon">🟡</div>
          <div className="stat-info">
            <h3>{stats.medium}</h3>
            <p>Средних</p>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon">🔴</div>
          <div className="stat-info">
            <h3>{stats.hard}</h3>
            <p>Сложных</p>
          </div>
        </div>
      </div>

      <button className="add-question-btn" onClick={() => setShowModal(true)}>
        ➕ Добавить вопрос
      </button>

      {/* Модальное окно */}
      {showModal && (
        <div className="modal-overlay" onClick={resetForm}>
          <div className="question-modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>{editingId ? '✏️ Редактировать вопрос' : '📝 Новый вопрос'}</h3>
              <button className="modal-close" onClick={resetForm}>✕</button>
            </div>
            <div className="modal-body">
              <div className="form-group">
                <label>Вопрос</label>
                <textarea name="question" rows="3" placeholder="Введите вопрос..." value={formData.question} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>Правильный ответ</label>
                <input type="text" name="answer" placeholder="Правильный ответ" value={formData.answer} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>Предмет</label>
                <input type="text" name="subject" placeholder="Например: Программирование" value={formData.subject} onChange={handleInputChange} />
              </div>
              <div className="form-group">
                <label>Сложность</label>
                <div className="difficulty-select">
                  <div className={`difficulty-option easy ${formData.difficulty === 'easy' ? 'selected' : ''}`} onClick={() => setDifficulty('easy')}>
                    🟢 Легкий
                  </div>
                  <div className={`difficulty-option medium ${formData.difficulty === 'medium' ? 'selected' : ''}`} onClick={() => setDifficulty('medium')}>
                    🟡 Средний
                  </div>
                  <div className={`difficulty-option hard ${formData.difficulty === 'hard' ? 'selected' : ''}`} onClick={() => setDifficulty('hard')}>
                    🔴 Сложный
                  </div>
                </div>
              </div>
              <div className="form-actions">
                <button className="submit-btn" onClick={editingId ? updateQuestion : addQuestion}>
                  {editingId ? '💾 Сохранить' : '📤 Добавить'}
                </button>
                <button className="cancel-btn" onClick={resetForm}>Отмена</button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Список вопросов */}
      <div className="questions-grid">
        {questions.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">📭</div>
            <div className="empty-text">Нет вопросов. Добавьте первый!</div>
            <button className="empty-btn" onClick={() => setShowModal(true)}>➕ Добавить вопрос</button>
          </div>
        ) : (
          questions.map(q => (
            <div key={q.id} className={`question-card ${q.difficulty}`}>
              <div className="question-header">
                <div className="question-text">❓ {q.question}</div>
                <div className={`difficulty-badge ${q.difficulty}`}>
                  {getDifficultyText(q.difficulty)}
                </div>
              </div>
              <div className="question-answer">
                <strong>✅ Ответ:</strong> {q.answer}
              </div>
              {q.subject && (
                <div className="question-subject">
                  <span>📚</span> {q.subject}
                </div>
              )}
              <div className="question-actions">
                <button className="edit-question" onClick={() => editQuestion(q)}>
                  ✏️ Редактировать
                </button>
                <button className="delete-question" onClick={() => deleteQuestion(q.id)}>
                  🗑️ Удалить
                </button>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default AdminExam;