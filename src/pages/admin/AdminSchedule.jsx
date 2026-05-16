import React, { useState, useEffect } from 'react';
import './AdminSchedule.css';

function AdminSchedule() {
  const [schedule, setSchedule] = useState([]);
  const [showForm, setShowForm] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [formData, setFormData] = useState({
    group: "",
    discipline: "",
    teacher: "",
    time: "",
    audience: "",
    day: "",
    comment: ""
  });

  useEffect(() => {
    const saved = localStorage.getItem('adminSchedule');
    if (saved) setSchedule(JSON.parse(saved));
  }, []);

  const saveToLocalStorage = (newSchedule) => {
    setSchedule(newSchedule);
    localStorage.setItem('adminSchedule', JSON.stringify(newSchedule));
  };

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const addSchedule = () => {
    if (!formData.discipline || !formData.group) {
      alert("Заполните обязательные поля");
      return;
    }
    const newItem = {
      id: Date.now(),
      ...formData,
      comment: formData.comment || ""
    };
    saveToLocalStorage([...schedule, newItem]);
    resetForm();
  };

  const updateSchedule = () => {
    const updated = schedule.map(item =>
      item.id === editingId ? { ...item, ...formData } : item
    );
    saveToLocalStorage(updated);
    resetForm();
  };

  const deleteSchedule = (id) => {
    if (window.confirm("Удалить занятие?")) {
      saveToLocalStorage(schedule.filter(item => item.id !== id));
    }
  };

  const editSchedule = (item) => {
    setEditingId(item.id);
    setFormData(item);
    setShowForm(true);
  };

  const updateComment = (id, comment) => {
    const updated = schedule.map(item =>
      item.id === id ? { ...item, comment } : item
    );
    saveToLocalStorage(updated);
  };

  const resetForm = () => {
    setFormData({ group: "", discipline: "", teacher: "", time: "", audience: "", day: "", comment: "" });
    setEditingId(null);
    setShowForm(false);
  };

  const days = ["Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"];

  return (
    <div className="admin-schedule">
      <div className="schedule-header">
        <div className="schedule-title">
          <div className="schedule-title-icon">📅</div>
          <h1>Управление расписанием</h1>
        </div>
        <div className="schedule-subtitle">Добавляйте, редактируйте и комментируйте занятия</div>
      </div>

      <button className="add-schedule-btn" onClick={() => setShowForm(true)}>
        ➕ Добавить занятие
      </button>

      {/* Форма добавления/редактирования */}
      <div className={`add-schedule-form ${showForm ? 'open' : ''}`}>
        <div className="form-title">
          <span>{editingId ? '✏️' : '➕'}</span>
          <span>{editingId ? 'Редактировать занятие' : 'Новое занятие'}</span>
        </div>
        <div className="form-grid">
          <div className="form-field">
            <label>Группа *</label>
            <input type="text" name="group" placeholder="23-ПМбо-1" value={formData.group} onChange={handleInputChange} />
          </div>
          <div className="form-field">
            <label>Дисциплина *</label>
            <input type="text" name="discipline" placeholder="Математика" value={formData.discipline} onChange={handleInputChange} />
          </div>
          <div className="form-field">
            <label>Преподаватель</label>
            <input type="text" name="teacher" placeholder="Иванов И.И." value={formData.teacher} onChange={handleInputChange} />
          </div>
          <div className="form-field">
            <label>Время</label>
            <input type="text" name="time" placeholder="11:50" value={formData.time} onChange={handleInputChange} />
          </div>
          <div className="form-field">
            <label>Аудитория</label>
            <input type="text" name="audience" placeholder="Б-201" value={formData.audience} onChange={handleInputChange} />
          </div>
          <div className="form-field">
            <label>День недели</label>
            <select name="day" value={formData.day} onChange={handleInputChange}>
              <option value="">Выберите день</option>
              {days.map(d => <option key={d} value={d}>{d}</option>)}
            </select>
          </div>
        </div>
        <div className="form-actions">
          <button className="submit-btn" onClick={editingId ? updateSchedule : addSchedule}>
            {editingId ? '💾 Сохранить' : '➕ Добавить'}
          </button>
          <button className="cancel-btn" onClick={resetForm}>Отмена</button>
        </div>
      </div>

      {/* Список занятий */}
      <div className="schedule-list">
        {schedule.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">📭</div>
            <div className="empty-text">Нет занятий в расписании</div>
            <button className="empty-btn" onClick={() => setShowForm(true)}>➕ Добавить первое занятие</button>
          </div>
        ) : (
          schedule.map((item) => {
            const [showCommentInput, setShowCommentInput] = useState(false);
            const [tempComment, setTempComment] = useState(item.comment || "");

            return (
              <div key={item.id} className="schedule-item" style={{ borderLeftColor: '#f5576c' }}>
                <div className="schedule-item-header">
                  <div className="schedule-discipline">{item.discipline}</div>
                  <div className="schedule-badge">
                    {item.group && <span className="badge badge-group">📚 {item.group}</span>}
                    {item.time && <span className="badge badge-time">⏰ {item.time}</span>}
                    {item.day && <span className="badge badge-group">📅 {item.day}</span>}
                  </div>
                </div>
                
                <div className="schedule-details">
                  {item.teacher && (
                    <div className="detail-item">
                      <span className="detail-icon">👨‍🏫</span>
                      <span>{item.teacher}</span>
                    </div>
                  )}
                  {item.audience && (
                    <div className="detail-item">
                      <span className="detail-icon">📍</span>
                      <span>{item.audience}</span>
                    </div>
                  )}
                </div>

                {/* Комментарий */}
                <div className="comment-section">
                  <div className="comment-label">
                    <span>💬</span>
                    <span>Комментарий преподавателя</span>
                  </div>
                  {showCommentInput ? (
                    <>
                      <textarea
                        className="comment-input"
                        rows="2"
                        placeholder="Добавьте комментарий к занятию..."
                        value={tempComment}
                        onChange={(e) => setTempComment(e.target.value)}
                      />
                      <div className="comment-actions">
                        <button 
                          className="save-comment-btn"
                          onClick={() => {
                            updateComment(item.id, tempComment);
                            setShowCommentInput(false);
                          }}
                        >
                          💾 Сохранить
                        </button>
                        <button 
                          className="cancel-btn"
                          onClick={() => {
                            setTempComment(item.comment || "");
                            setShowCommentInput(false);
                          }}
                        >
                          Отмена
                        </button>
                      </div>
                    </>
                  ) : (
                    <>
                      <div className="comment-text">
                        {item.comment ? item.comment : "Нет комментариев"}
                      </div>
                      <div style={{ marginTop: '8px' }}>
                        <button 
                          className="comment-btn"
                          onClick={() => setShowCommentInput(true)}
                        >
                          ✏️ Добавить комментарий
                        </button>
                      </div>
                    </>
                  )}
                </div>

                <div className="schedule-actions">
                  <button className="edit-btn" onClick={() => editSchedule(item)}>
                    ✏️ Редактировать
                  </button>
                  <button className="delete-btn" onClick={() => deleteSchedule(item.id)}>
                    🗑️ Удалить
                  </button>
                </div>
              </div>
            );
          })
        )}
      </div>
    </div>
  );
}

export default AdminSchedule;