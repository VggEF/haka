import React, { useState, useEffect } from 'react';
import './Profile.css';

// Система уровней
const LEVELS = [
  { name: "🌱 Стажер", minScore: 0, maxScore: 100, color: "#9E9E9E", icon: "🌱" },
  { name: "📘 Младший специалист", minScore: 101, maxScore: 300, color: "#4CAF50", icon: "📘" },
  { name: "📚 Специалист", minScore: 301, maxScore: 600, color: "#2196F3", icon: "📚" },
  { name: "⭐ Старший специалист", minScore: 601, maxScore: 1000, color: "#FF9800", icon: "⭐" },
  { name: "🚀 Джуниор", minScore: 1001, maxScore: 1500, color: "#9C27B0", icon: "🚀" },
  { name: "🔥 Мидл", minScore: 1501, maxScore: 2500, color: "#E91E63", icon: "🔥" },
  { name: "💎 Сеньор", minScore: 2501, maxScore: 4000, color: "#00BCD4", icon: "💎" },
  { name: "👑 Тимлид+", minScore: 4001, maxScore: Infinity, color: "#FFD700", icon: "👑" },
];

const getLevelByScore = (score) => {
  return LEVELS.find(level => score >= level.minScore && score <= level.maxScore) || LEVELS[0];
};

// Типы мероприятий и их баллы для ПГАС
const PGAS_EVENT_TYPES = {
  deadline: { name: "Дедлайн", score: 10, icon: "📝" },
  lecture: { name: "Лекция", score: 15, icon: "🎓" },
  hackathon: { name: "Хакатон", score: 50, icon: "🏆" },
  meeting: { name: "Встреча", score: 5, icon: "💬" },
  defense: { name: "Защита", score: 30, icon: "📊" },
  party: { name: "Мероприятие", score: 20, icon: "🎉" },
};

// ПГАС мероприятия пользователя
const DEFAULT_PGAS_EVENTS = [
  { id: 1, title: "Хакатон по веб-разработке", date: "2026-05-10", type: "hackathon", status: "completed", score: 50 },
  { id: 2, title: "Лекция по искусственному интеллекту", date: "2026-05-12", type: "lecture", status: "completed", score: 15 },
  { id: 3, title: "Встреча с куратором", date: "2026-05-15", type: "meeting", status: "completed", score: 5 },
  { id: 4, title: "Студенческая вечеринка", date: "2026-05-20", type: "party", status: "upcoming", score: null },
  { id: 5, title: "Защита курсового проекта", date: "2026-06-01", type: "defense", status: "upcoming", score: null },
  { id: 6, title: "Дедлайн: Курсовая работа", date: "2026-05-25", type: "deadline", status: "upcoming", score: null },
];

// Обычные мероприятия (не ПГАС)
const DEFAULT_EVENTS = [
  { id: 1, title: "📝 Дедлайн: Лабораторная работа №3", date: "2026-05-20", type: "deadline", priority: "high" },
  { id: 2, title: "🎓 Лекция: «Современные технологии ИИ»", date: "2026-05-18", type: "lecture", priority: "medium" },
  { id: 3, title: "🏆 Хакатон по веб-разработке", date: "2026-05-25", type: "hackathon", priority: "high" },
  { id: 4, title: "💬 Встреча с куратором", date: "2026-05-22", type: "meeting", priority: "medium" },
  { id: 5, title: "📊 Защита курсового проекта", date: "2026-06-01", type: "defense", priority: "high" },
  { id: 6, title: "🎉 Студенческая вечеринка", date: "2026-05-30", type: "party", priority: "low" },
];

const DEFAULT_PROFILE = {
  photo: null,
  firstName: "Иван",
  lastName: "Иванов",
  patronymic: "Иванович",
  group: "ПМбо-23",
  course: 3,
  faculty: "ИВИТШ",
  email: "ivan.ivanov@student.ru",
  phone: "+7 (999) 123-45-67",
  telegram: "@ivanov_ivan",
  vk: "vk.com/ivanov",
  github: "github.com/ivanov",
  hobbies: ["Программирование", "Спорт", "Шахматы", "Чтение книг"],
  clubs: ["Студенческий совет", "Киберспортивный клуб", "Клуб робототехники"],
  events: DEFAULT_EVENTS,
  pgasEvents: DEFAULT_PGAS_EVENTS,
  grades: [
    { subject: "Математический анализ", grade: 5, type: "Экзамен", teacher: "Сухов А.К." },
    { subject: "Программирование", grade: 5, type: "Экзамен", teacher: "Леготин Д.Л." },
    { subject: "Алгебра и геометрия", grade: 4, type: "Экзамен", teacher: "Секованов В.С." },
    { subject: "Физика", grade: 3, type: "Зачет", teacher: "Петров С.И." },
    { subject: "Английский язык", grade: 5, type: "Зачет", teacher: "Смирнова Е.А." },
    { subject: "Web-технологии", grade: 4, type: "Экзамен", teacher: "Леготин Д.Л." },
    { subject: "Базы данных", grade: 5, type: "Экзамен", teacher: "Ивков В.А." },
    { subject: "Философия", grade: 4, type: "Зачет", teacher: "Козлов А.М." },
  ]
};

function Profile({ login }) {
  const [profile, setProfile] = useState(DEFAULT_PROFILE);
  const [isEditing, setIsEditing] = useState(false);
  const [editForm, setEditForm] = useState(DEFAULT_PROFILE);
  const [tempPhoto, setTempPhoto] = useState(null);
  const [userScore, setUserScore] = useState(0);
  
  // Состояние для мероприятий
  const [showEventForm, setShowEventForm] = useState(false);
  const [showPgasEventForm, setShowPgasEventForm] = useState(false);
  const [newEvent, setNewEvent] = useState({ title: "", date: "", type: "deadline", priority: "medium" });
  const [newPgasEvent, setNewPgasEvent] = useState({ title: "", date: "", type: "hackathon" });
  
  useEffect(() => {
    if (login) {
      const savedScore = localStorage.getItem(`score_${login}`);
      setUserScore(savedScore ? parseInt(savedScore) : 0);
      
      const savedProfile = localStorage.getItem(`profile_${login}`);
      if (savedProfile) {
        setProfile(JSON.parse(savedProfile));
      }
    }
  }, [login]);
  
  // Подсчет общего количества баллов ПГАС
  const calculateTotalPgasScore = () => {
    const events = isEditing ? editForm.pgasEvents : profile.pgasEvents;
    return events
      .filter(e => e.status === "completed")
      .reduce((sum, e) => sum + (e.score || 0), 0);
  };
  
  const currentLevel = getLevelByScore(userScore + calculateTotalPgasScore());
  
  const handleEditClick = () => {
    setEditForm(profile);
    setTempPhoto(null);
    setIsEditing(true);
  };
  
  const handleSaveProfile = () => {
    let updatedProfile = { ...editForm };
    if (tempPhoto) {
      updatedProfile.photo = tempPhoto;
    }
    setProfile(updatedProfile);
    localStorage.setItem(`profile_${login}`, JSON.stringify(updatedProfile));
    setIsEditing(false);
  };
  
  const handlePhotoUpload = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setTempPhoto(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };
  
  const handleAddHobby = () => {
    const newHobby = prompt("Введите новое хобби:");
    if (newHobby && newHobby.trim()) {
      setEditForm({
        ...editForm,
        hobbies: [...editForm.hobbies, newHobby.trim()]
      });
    }
  };
  
  const handleRemoveHobby = (index) => {
    const newHobbies = editForm.hobbies.filter((_, i) => i !== index);
    setEditForm({ ...editForm, hobbies: newHobbies });
  };
  
  const handleAddClub = () => {
    const newClub = prompt("Введите название кружка/секции:");
    if (newClub && newClub.trim()) {
      setEditForm({
        ...editForm,
        clubs: [...editForm.clubs, newClub.trim()]
      });
    }
  };
  
  const handleRemoveClub = (index) => {
    const newClubs = editForm.clubs.filter((_, i) => i !== index);
    setEditForm({ ...editForm, clubs: newClubs });
  };
  
  // Добавление обычного мероприятия
  const handleAddEvent = () => {
    if (!newEvent.title.trim()) {
      alert("Введите название мероприятия");
      return;
    }
    if (!newEvent.date) {
      alert("Выберите дату");
      return;
    }
    
    const event = {
      id: Date.now(),
      title: newEvent.title,
      date: newEvent.date,
      type: newEvent.type,
      priority: newEvent.priority
    };
    
    const updatedEvents = [...editForm.events, event].sort((a, b) => new Date(a.date) - new Date(b.date));
    setEditForm({ ...editForm, events: updatedEvents });
    setNewEvent({ title: "", date: "", type: "deadline", priority: "medium" });
    setShowEventForm(false);
  };
  
  const handleRemoveEvent = (eventId) => {
    const updatedEvents = editForm.events.filter(e => e.id !== eventId);
    setEditForm({ ...editForm, events: updatedEvents });
  };
  
  // Добавление ПГАС мероприятия
  const handleAddPgasEvent = () => {
    if (!newPgasEvent.title.trim()) {
      alert("Введите название мероприятия");
      return;
    }
    if (!newPgasEvent.date) {
      alert("Выберите дату");
      return;
    }
    
    const event = {
      id: Date.now(),
      title: newPgasEvent.title,
      date: newPgasEvent.date,
      type: newPgasEvent.type,
      status: "upcoming",
      score: null
    };
    
    const updatedEvents = [...editForm.pgasEvents, event].sort((a, b) => new Date(a.date) - new Date(b.date));
    setEditForm({ ...editForm, pgasEvents: updatedEvents });
    setNewPgasEvent({ title: "", date: "", type: "hackathon" });
    setShowPgasEventForm(false);
  };
  
  // Отметка ПГАС мероприятия как выполненного
  const handleCompletePgasEvent = (eventId) => {
    const updatedEvents = editForm.pgasEvents.map(event => {
      if (event.id === eventId && event.status === "upcoming") {
        const eventType = PGAS_EVENT_TYPES[event.type];
        return {
          ...event,
          status: "completed",
          score: eventType.score
        };
      }
      return event;
    });
    setEditForm({ ...editForm, pgasEvents: updatedEvents });
  };
  
  const handleRemovePgasEvent = (eventId) => {
    const updatedEvents = editForm.pgasEvents.filter(e => e.id !== eventId);
    setEditForm({ ...editForm, pgasEvents: updatedEvents });
  };
  
  const getEventIcon = (type) => {
    const icons = {
      deadline: "📝",
      lecture: "🎓",
      hackathon: "🏆",
      meeting: "💬",
      defense: "📊",
      party: "🎉"
    };
    return icons[type] || "📌";
  };
  
  const getPgasEventIcon = (type) => {
    return PGAS_EVENT_TYPES[type]?.icon || "📌";
  };
  
  const getPriorityColor = (priority) => {
    if (priority === "high") return "#f44336";
    if (priority === "medium") return "#ff9800";
    return "#4caf50";
  };
  
  const formatEventDate = (dateStr) => {
    const date = new Date(dateStr);
    const today = new Date();
    const tomorrow = new Date(today);
    tomorrow.setDate(tomorrow.getDate() + 1);
    
    if (date.toDateString() === today.toDateString()) {
      return "Сегодня";
    }
    if (date.toDateString() === tomorrow.toDateString()) {
      return "Завтра";
    }
    return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'long' });
  };
  
  const getGradeColor = (grade) => {
    if (grade >= 5) return '#4CAF50';
    if (grade >= 4) return '#2196F3';
    if (grade >= 3) return '#FF9800';
    return '#F44336';
  };
  
  const getGradeText = (grade) => {
    if (grade >= 5) return 'Отлично';
    if (grade >= 4) return 'Хорошо';
    if (grade >= 3) return 'Удовлетворительно';
    return 'Неудовлетворительно';
  };
  
  const calculateAverageGrade = () => {
    const sum = profile.grades.reduce((acc, g) => acc + g.grade, 0);
    return (sum / profile.grades.length).toFixed(2);
  };

  return (
    <div className="profile-container">
      <div className="profile-content">
        {/* Левая колонка - фиксированная */}
        <div className="profile-left">
          <div className="sticky-sidebar">
            {/* Фото */}
            <div className="photo-container">
              {isEditing ? (
                <>
                  {tempPhoto ? (
                    <img src={tempPhoto} alt="Фото" className="profile-photo" />
                  ) : editForm.photo ? (
                    <img src={editForm.photo} alt="Фото" className="profile-photo" />
                  ) : (
                    <div className="profile-photo-placeholder">
                      <span className="photo-emoji">👨‍🎓</span>
                    </div>
                  )}
                  <label className="photo-upload-label">
                    📷 Загрузить фото
                    <input type="file" accept="image/*" onChange={handlePhotoUpload} hidden />
                  </label>
                </>
              ) : (
                <>
                  {profile.photo ? (
                    <img src={profile.photo} alt="Фото" className="profile-photo" />
                  ) : (
                    <div className="profile-photo-placeholder">
                      <span className="photo-emoji">👨‍🎓</span>
                    </div>
                  )}
                </>
              )}
            </div>
            
            {/* Блок ПГАС */}
            <div className="pgas-container">
              <div className="pgas-header">
                <div className="pgas-title">
                  <span className="pgas-icon">🏆</span>
                  <span>ПГАС</span>
                </div>
                <div className="pgas-total-score">
                  {calculateTotalPgasScore()} баллов
                </div>
              </div>
              
              <div className="pgas-stats">
                <div className="pgas-stat">
                  <span className="stat-value">{(isEditing ? editForm.pgasEvents : profile.pgasEvents).filter(e => e.status === "completed").length}</span>
                  <span className="stat-label">Выполнено</span>
                </div>
                <div className="pgas-stat">
                  <span className="stat-value">{(isEditing ? editForm.pgasEvents : profile.pgasEvents).filter(e => e.status === "upcoming").length}</span>
                  <span className="stat-label">В плане</span>
                </div>
              </div>
              
              <div className="pgas-events-header">
                <span>Мои мероприятия</span>
                {isEditing && (
                  <button className="add-event-btn" onClick={() => setShowPgasEventForm(true)}>+</button>
                )}
              </div>
              
              {/* Форма добавления ПГАС мероприятия */}
              {showPgasEventForm && (
                <div className="event-form">
                  <input
                    type="text"
                    placeholder="Название мероприятия"
                    value={newPgasEvent.title}
                    onChange={(e) => setNewPgasEvent({ ...newPgasEvent, title: e.target.value })}
                    className="event-input"
                  />
                  <input
                    type="date"
                    value={newPgasEvent.date}
                    onChange={(e) => setNewPgasEvent({ ...newPgasEvent, date: e.target.value })}
                    className="event-input"
                  />
                  <select
                    value={newPgasEvent.type}
                    onChange={(e) => setNewPgasEvent({ ...newPgasEvent, type: e.target.value })}
                    className="event-select"
                  >
                    <option value="deadline">📝 Дедлайн (+10)</option>
                    <option value="lecture">🎓 Лекция (+15)</option>
                    <option value="hackathon">🏆 Хакатон (+50)</option>
                    <option value="meeting">💬 Встреча (+5)</option>
                    <option value="defense">📊 Защита (+30)</option>
                    <option value="party">🎉 Мероприятие (+20)</option>
                  </select>
                  <div className="event-form-buttons">
                    <button className="event-submit" onClick={handleAddPgasEvent}>Добавить</button>
                    <button className="event-cancel" onClick={() => setShowPgasEventForm(false)}>Отмена</button>
                  </div>
                </div>
              )}
              
              {/* Список ПГАС мероприятий (скроллится внутри) */}
              <div className="pgas-events-list">
                {(isEditing ? editForm.pgasEvents : profile.pgasEvents).length === 0 ? (
                  <div className="no-events">Нет мероприятий</div>
                ) : (
                  (isEditing ? editForm.pgasEvents : profile.pgasEvents)
                    .sort((a, b) => new Date(a.date) - new Date(b.date))
                    .map(event => (
                      <div key={event.id} className={`pgas-event-item ${event.status}`}>
                        <div className="event-icon">{getPgasEventIcon(event.type)}</div>
                        <div className="event-info">
                          <div className="event-title">{event.title}</div>
                          <div className="event-date">{formatEventDate(event.date)}</div>
                        </div>
                        <div className="event-score">
                          {event.status === "completed" ? (
                            <span className="score-badge">+{event.score}</span>
                          ) : (
                            isEditing && (
                              <button 
                                className="complete-btn"
                                onClick={() => handleCompletePgasEvent(event.id)}
                                title="Отметить как выполненное"
                              >
                                ✅
                              </button>
                            )
                          )}
                          {isEditing && event.status === "upcoming" && (
                            <button className="event-delete" onClick={() => handleRemovePgasEvent(event.id)}>✕</button>
                          )}
                        </div>
                      </div>
                    ))
                )}
              </div>
            </div>
            
            {/* Блок обычных мероприятий (скроллится внутри) */}
            <div className="events-container">
              <div className="events-header">
                <h3 className="events-title">📅 Ближайшие мероприятия</h3>
                {isEditing && (
                  <button className="add-event-btn" onClick={() => setShowEventForm(true)}>+</button>
                )}
              </div>
              
              {showEventForm && (
                <div className="event-form">
                  <input
                    type="text"
                    placeholder="Название"
                    value={newEvent.title}
                    onChange={(e) => setNewEvent({ ...newEvent, title: e.target.value })}
                    className="event-input"
                  />
                  <input
                    type="date"
                    value={newEvent.date}
                    onChange={(e) => setNewEvent({ ...newEvent, date: e.target.value })}
                    className="event-input"
                  />
                  <select
                    value={newEvent.type}
                    onChange={(e) => setNewEvent({ ...newEvent, type: e.target.value })}
                    className="event-select"
                  >
                    <option value="deadline">Дедлайн</option>
                    <option value="lecture">Лекция</option>
                    <option value="hackathon">Хакатон</option>
                    <option value="meeting">Встреча</option>
                    <option value="defense">Защита</option>
                    <option value="party">Вечеринка</option>
                  </select>
                  <select
                    value={newEvent.priority}
                    onChange={(e) => setNewEvent({ ...newEvent, priority: e.target.value })}
                    className="event-select"
                  >
                    <option value="high">Высокий приоритет</option>
                    <option value="medium">Средний приоритет</option>
                    <option value="low">Низкий приоритет</option>
                  </select>
                  <div className="event-form-buttons">
                    <button className="event-submit" onClick={handleAddEvent}>Добавить</button>
                    <button className="event-cancel" onClick={() => setShowEventForm(false)}>Отмена</button>
                  </div>
                </div>
              )}
              
              <div className="events-list">
                {(isEditing ? editForm.events : profile.events).length === 0 ? (
                  <div className="no-events">Нет ближайших мероприятий</div>
                ) : (
                  (isEditing ? editForm.events : profile.events)
                    .sort((a, b) => new Date(a.date) - new Date(b.date))
                    .slice(0, 5)
                    .map(event => (
                      <div key={event.id} className="event-item">
                        <div className="event-icon">{getEventIcon(event.type)}</div>
                        <div className="event-info">
                          <div className="event-title">{event.title}</div>
                          <div className="event-date" style={{ color: getPriorityColor(event.priority) }}>
                            {formatEventDate(event.date)}
                          </div>
                        </div>
                        {isEditing && (
                          <button className="event-delete" onClick={() => handleRemoveEvent(event.id)}>✕</button>
                        )}
                      </div>
                    ))
                )}
              </div>
            </div>
          </div>
        </div>
        
        {/* Правая колонка - скроллится */}
        <div className="profile-right">
          {/* Заголовок */}
          <div className="profile-header">
            <div className="profile-name-section">
              {isEditing ? (
                <div className="edit-name-fields">
                  <input
                    type="text"
                    value={editForm.lastName}
                    onChange={(e) => setEditForm({ ...editForm, lastName: e.target.value })}
                    placeholder="Фамилия"
                    className="edit-input-small"
                  />
                  <input
                    type="text"
                    value={editForm.firstName}
                    onChange={(e) => setEditForm({ ...editForm, firstName: e.target.value })}
                    placeholder="Имя"
                    className="edit-input-small"
                  />
                  <input
                    type="text"
                    value={editForm.patronymic}
                    onChange={(e) => setEditForm({ ...editForm, patronymic: e.target.value })}
                    placeholder="Отчество"
                    className="edit-input-small"
                  />
                </div>
              ) : (
                <h1 className="profile-name">
                  {profile.lastName} {profile.firstName} {profile.patronymic}
                </h1>
              )}
              
              <div className="profile-badges">
                <div className="profile-group-badge">
                  📚 {profile.group} | {profile.course} курс | {profile.faculty}
                </div>
                <div className="profile-rank-badge" style={{ backgroundColor: currentLevel.color }}>
                  {currentLevel.icon} {currentLevel.name}
                </div>
              </div>
            </div>
            
            {!isEditing && (
              <button className="edit-profile-btn" onClick={handleEditClick}>
                ✏️ Редактировать
              </button>
            )}
          </div>
          
          {/* Контакты */}
          <div className="profile-section">
            <h3 className="section-title">📞 Контактная информация</h3>
            <div className="contact-info">
              {isEditing ? (
                <>
                  <div className="contact-row">
                    <span className="contact-icon">📧</span>
                    <input type="email" value={editForm.email} onChange={(e) => setEditForm({ ...editForm, email: e.target.value })} className="edit-input" />
                  </div>
                  <div className="contact-row">
                    <span className="contact-icon">📱</span>
                    <input type="tel" value={editForm.phone} onChange={(e) => setEditForm({ ...editForm, phone: e.target.value })} className="edit-input" />
                  </div>
                  <div className="contact-row">
                    <span className="contact-icon">💬</span>
                    <input type="text" value={editForm.telegram} onChange={(e) => setEditForm({ ...editForm, telegram: e.target.value })} className="edit-input" placeholder="Telegram" />
                  </div>
                  <div className="contact-row">
                    <span className="contact-icon">📘</span>
                    <input type="text" value={editForm.vk} onChange={(e) => setEditForm({ ...editForm, vk: e.target.value })} className="edit-input" placeholder="ВКонтакте" />
                  </div>
                  <div className="contact-row">
                    <span className="contact-icon">🐙</span>
                    <input type="text" value={editForm.github} onChange={(e) => setEditForm({ ...editForm, github: e.target.value })} className="edit-input" placeholder="GitHub" />
                  </div>
                </>
              ) : (
                <>
                  <div className="contact-row"><span className="contact-icon">📧</span><span>{profile.email}</span></div>
                  <div className="contact-row"><span className="contact-icon">📱</span><span>{profile.phone}</span></div>
                  <div className="contact-row"><span className="contact-icon">💬</span><span>{profile.telegram}</span></div>
                  <div className="contact-row"><span className="contact-icon">📘</span><span>{profile.vk}</span></div>
                  <div className="contact-row"><span className="contact-icon">🐙</span><span>{profile.github}</span></div>
                </>
              )}
            </div>
          </div>
          
          {/* Хобби и кружки */}
          <div className="profile-row-sections">
            <div className="profile-section half">
              <h3 className="section-title">🎯 Хобби</h3>
              <div className="tags-list">
                {isEditing ? (
                  <>
                    {editForm.hobbies.map((hobby, idx) => (
                      <div key={idx} className="tag-item">
                        <span className="tag">{hobby}</span>
                        <button className="tag-remove" onClick={() => handleRemoveHobby(idx)}>✕</button>
                      </div>
                    ))}
                    <button className="add-tag-btn" onClick={handleAddHobby}>+ Добавить</button>
                  </>
                ) : (
                  profile.hobbies.map((hobby, idx) => <span key={idx} className="tag">{hobby}</span>)
                )}
              </div>
            </div>
            
            <div className="profile-section half">
              <h3 className="section-title">🏆 Кружки и секции</h3>
              <div className="tags-list">
                {isEditing ? (
                  <>
                    {editForm.clubs.map((club, idx) => (
                      <div key={idx} className="tag-item">
                        <span className="tag club-tag">{club}</span>
                        <button className="tag-remove" onClick={() => handleRemoveClub(idx)}>✕</button>
                      </div>
                    ))}
                    <button className="add-tag-btn" onClick={handleAddClub}>+ Добавить</button>
                  </>
                ) : (
                  profile.clubs.map((club, idx) => <span key={idx} className="tag club-tag">{club}</span>)
                )}
              </div>
            </div>
          </div>
          
          {/* Успеваемость */}
          <div className="profile-section">
            <h3 className="section-title">📊 Успеваемость</h3>
            <div className="grades-summary">
              <div className="average-grade">
                <span className="average-value">{calculateAverageGrade()}</span>
                <span className="average-label">Средний балл</span>
              </div>
            </div>
            <div className="grades-list">
              <div className="grades-header">
                <span>Дисциплина</span>
                <span>Оценка</span>
                <span>Тип</span>
                <span>Преподаватель</span>
              </div>
              {profile.grades.map((grade, idx) => (
                <div key={idx} className="grade-row">
                  <span className="grade-subject">{grade.subject}</span>
                  <span className="grade-value" style={{ color: getGradeColor(grade.grade) }}>
                    {grade.grade} ({getGradeText(grade.grade)})
                  </span>
                  <span className="grade-type">{grade.type}</span>
                  <span className="grade-teacher">{grade.teacher}</span>
                </div>
              ))}
            </div>
          </div>
          
          {isEditing && (
            <div className="edit-actions">
              <button className="save-btn" onClick={handleSaveProfile}>💾 Сохранить</button>
              <button className="cancel-btn" onClick={() => setIsEditing(false)}>❌ Отмена</button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Profile;