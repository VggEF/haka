import React, { useState, useEffect } from 'react';
import './AdminEvents.css';

function AdminEvents() {
  const [events, setEvents] = useState([]);
  const [filter, setFilter] = useState('all');
  const [showModal, setShowModal] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [formData, setFormData] = useState({
    title: "",
    shortText: "",
    fullText: "",
    date: "",
    time: "",
    type: "student",
    category: "",
    location: "",
    price: "Бесплатно",
    organizer: "",
    image: "🎯",
    availableSpots: "",
    contact: ""
  });

  useEffect(() => {
    const saved = localStorage.getItem('entertainmentEvents');
    if (saved) setEvents(JSON.parse(saved));
  }, []);

  const saveToLocalStorage = (newEvents) => {
    setEvents(newEvents);
    localStorage.setItem('entertainmentEvents', JSON.stringify(newEvents));
  };

  const handleInputChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const addEvent = () => {
    if (!formData.title || !formData.date) {
      alert("Заполните название и дату");
      return;
    }
    const newEvent = {
      id: Date.now(),
      ...formData,
      image: getEventIcon(formData.type)
    };
    saveToLocalStorage([...events, newEvent]);
    resetForm();
  };

  const updateEvent = () => {
    const updated = events.map(item =>
      item.id === editingId ? { ...item, ...formData, image: getEventIcon(formData.type) } : item
    );
    saveToLocalStorage(updated);
    resetForm();
  };

  const deleteEvent = (id) => {
    if (window.confirm("Удалить мероприятие?")) {
      saveToLocalStorage(events.filter(item => item.id !== id));
    }
  };

  const editEvent = (item) => {
    setEditingId(item.id);
    setFormData(item);
    setShowModal(true);
  };

  const resetForm = () => {
    setFormData({
      title: "",
      shortText: "",
      fullText: "",
      date: "",
      time: "",
      type: "student",
      category: "",
      location: "",
      price: "Бесплатно",
      organizer: "",
      image: "🎯",
      availableSpots: "",
      contact: ""
    });
    setEditingId(null);
    setShowModal(false);
  };

  const getEventIcon = (type) => {
    const icons = {
      sport: '🏀',
      school: '🔬',
      student: '🎮'
    };
    return icons[type] || '🎯';
  };

  const getTypeBadgeClass = (type) => {
    const classes = {
      sport: 'sport',
      school: 'school',
      student: 'student'
    };
    return classes[type] || 'student';
  };

  const getTypeName = (type) => {
    const names = {
      sport: 'Спорт',
      school: 'Школьникам',
      student: 'Студенческое'
    };
    return names[type] || 'Мероприятие';
  };

  const filteredEvents = filter === 'all' 
    ? events 
    : events.filter(e => e.type === filter);

  return (
    <div className="admin-events">
      <div className="events-header">
        <div className="events-title">
          <div className="events-title-icon">🎮</div>
          <h1>Управление ДПО</h1>
        </div>
        <div className="events-subtitle">Создавайте мероприятия для студентов и школьников</div>
      </div>

      {/* Фильтры - без кнопки создания */}
      <div className="events-filters">
        <button className={`filter-btn ${filter === 'all' ? 'active' : ''}`} onClick={() => setFilter('all')}>
          Все
        </button>
        <button className={`filter-btn ${filter === 'sport' ? 'active' : ''}`} onClick={() => setFilter('sport')}>
          🏀 Спорт
        </button>
        <button className={`filter-btn ${filter === 'school' ? 'active' : ''}`} onClick={() => setFilter('school')}>
          🔬 Школьникам
        </button>
        <button className={`filter-btn ${filter === 'student' ? 'active' : ''}`} onClick={() => setFilter('student')}>
          🎮 Студенческие
        </button>
      </div>

      {/* Кнопка создания отдельно */}
      <button className="add-event-btn" onClick={() => setShowModal(true)}>
        + Создать мероприятие
      </button>

      {/* Модальное окно */}
      {showModal && (
        <div className="modal-overlay" onClick={resetForm}>
          <div className="event-modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>
                {editingId ? '✏️ Редактировать мероприятие' : '🎉 Создать мероприятие'}
              </h3>
              <button className="modal-close" onClick={resetForm}>✕</button>
            </div>
            <div className="modal-body">
              <div className="form-grid">
                <div className="form-group">
                  <label>Название *</label>
                  <input type="text" name="title" placeholder="Название мероприятия" value={formData.title} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Тип мероприятия *</label>
                  <select name="type" value={formData.type} onChange={handleInputChange}>
                    <option value="sport">🏀 Спорт</option>
                    <option value="school">🔬 Школьникам</option>
                    <option value="student">🎮 Студенческое</option>
                  </select>
                </div>
                <div className="form-group">
                  <label>Дата *</label>
                  <input type="date" name="date" value={formData.date} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Время</label>
                  <input type="text" name="time" placeholder="Например: 15:00" value={formData.time} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Место проведения</label>
                  <input type="text" name="location" placeholder="Аудитория, корпус" value={formData.location} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Стоимость</label>
                  <input type="text" name="price" placeholder="Бесплатно / 500₽" value={formData.price} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Организатор</label>
                  <input type="text" name="organizer" placeholder="Кто проводит" value={formData.organizer} onChange={handleInputChange} />
                </div>
                <div className="form-group">
                  <label>Свободных мест</label>
                  <input type="text" name="availableSpots" placeholder="Количество мест" value={formData.availableSpots} onChange={handleInputChange} />
                </div>
                <div className="form-group full-width">
                  <label>Краткое описание</label>
                  <textarea name="shortText" rows="2" placeholder="Краткое описание мероприятия..." value={formData.shortText} onChange={handleInputChange} />
                </div>
                <div className="form-group full-width">
                  <label>Полное описание</label>
                  <textarea name="fullText" rows="4" placeholder="Подробное описание, программа, требования..." value={formData.fullText} onChange={handleInputChange} />
                </div>
                <div className="form-actions">
                  <button className="submit-btn" onClick={editingId ? updateEvent : addEvent}>
                    {editingId ? '💾 Сохранить' : '📤 Опубликовать'}
                  </button>
                  <button className="cancel-btn" onClick={resetForm}>Отмена</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Список мероприятий */}
      <div className="events-grid">
        {filteredEvents.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">🎯</div>
            <div className="empty-text">Нет мероприятий. Создайте первое!</div>
          </div>
        ) : (
          filteredEvents.map((event) => (
            <div key={event.id} className="event-card">
              <div className={`event-badge ${getTypeBadgeClass(event.type)}`}>
                {getTypeName(event.type)}
              </div>
              <div className="event-image">
                <span>{event.image || getEventIcon(event.type)}</span>
                <div className="event-date">📅 {event.date}</div>
              </div>
              <div className="event-content">
                <h3 className="event-title">{event.title}</h3>
                <div className="event-meta">
                  {event.time && <span>⏰ {event.time}</span>}
                  {event.location && <span>📍 {event.location}</span>}
                  {event.organizer && <span>👥 {event.organizer}</span>}
                </div>
                <p className="event-description">{event.shortText || event.fullText?.substring(0, 100)}</p>
                <div className="event-footer">
                  <div className="event-price">{event.price || "Бесплатно"}</div>
                  <div className="event-actions">
                    <button className="edit-event" onClick={() => editEvent(event)}>
                      ✏️ Ред.
                    </button>
                    <button className="delete-event" onClick={() => deleteEvent(event.id)}>
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

export default AdminEvents;