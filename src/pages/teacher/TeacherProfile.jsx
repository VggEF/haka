import React, { useState, useEffect } from 'react';
import './TeacherProfile.css';

function TeacherProfile() {
  const [profile, setProfile] = useState({
    photo: null,
    lastName: "Иванов",
    firstName: "Иван",
    patronymic: "Иванович",
    department: "ИВИТШ",
    position: "Старший преподаватель",
    degree: "Кандидат технических наук",
    email: "ivan.ivanov@kosgos.ru",
    phone: "+7 (999) 123-45-67",
    telegram: "@prof_ivanov",
    office: "Б-301",
    experience: 12,
    achievements: [
      { id: 1, title: "Лучший преподаватель года", date: "2025", icon: "🏆" },
      { id: 2, title: "Публикация в журнале ВАК", date: "2024", icon: "📝" },
      { id: 3, title: "Руководство победителем олимпиады", date: "2025", icon: "🎓" }
    ],
    courses: [
      { id: 1, name: "Программирование", group: "23-ПМбо-1", hours: 72, students: 24 },
      { id: 2, name: "Базы данных", group: "23-ПМбо-1", hours: 54, students: 24 },
      { id: 3, name: "Web-технологии", group: "23-ПМбо-2", hours: 72, students: 22 },
      { id: 4, name: "Алгоритмы и структуры данных", group: "23-ПМбо-2", hours: 64, students: 22 }
    ],
    schedule: [
      { day: "Понедельник", time: "11:50-13:20", subject: "Программирование", group: "23-ПМбо-1", room: "Б-201" },
      { day: "Вторник", time: "09:40-11:10", subject: "Базы данных", group: "23-ПМбо-1", room: "Б-303" },
      { day: "Среда", time: "13:00-14:30", subject: "Web-технологии", group: "23-ПМбо-2", room: "Б-208" },
      { day: "Пятница", time: "11:50-13:20", subject: "Алгоритмы", group: "23-ПМбо-2", room: "Б-301" }
    ]
  });

  const [isEditing, setIsEditing] = useState(false);
  const [editForm, setEditForm] = useState(profile);
  const [tempPhoto, setTempPhoto] = useState(null);
  const [activeTab, setActiveTab] = useState('info');

  useEffect(() => {
    const saved = localStorage.getItem('teacherProfile');
    if (saved) setProfile(JSON.parse(saved));
  }, []);

  const handleSave = () => {
    let updated = { ...editForm };
    if (tempPhoto) updated.photo = tempPhoto;
    setProfile(updated);
    localStorage.setItem('teacherProfile', JSON.stringify(updated));
    setIsEditing(false);
  };

  const handlePhotoUpload = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => setTempPhoto(reader.result);
      reader.readAsDataURL(file);
    }
  };

  const handleAddAchievement = () => {
    const title = prompt("Название достижения");
    const date = prompt("Год");
    if (title && date) {
      setEditForm({
        ...editForm,
        achievements: [...editForm.achievements, { id: Date.now(), title, date, icon: "🏆" }]
      });
    }
  };

  const handleRemoveAchievement = (id) => {
    setEditForm({
      ...editForm,
      achievements: editForm.achievements.filter(a => a.id !== id)
    });
  };

  const formatDate = (dateStr) => {
    const date = new Date(dateStr);
    return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'long', year: 'numeric' });
  };

  return (
    <div className="teacher-profile">
      <div className="teacher-profile-header">
        <div className="teacher-profile-title">
          <div className="teacher-profile-icon">👨‍🏫</div>
          <h1>Профиль преподавателя</h1>
        </div>
        <div className="teacher-profile-subtitle">Личная информация и профессиональная деятельность</div>
      </div>

      <div className="teacher-profile-content">
        {/* Левая колонка - фото и контакты */}
        <div className="teacher-profile-left">
          <div className="teacher-photo-card">
            {isEditing ? (
              <>
                {tempPhoto || editForm.photo ? (
                  <img src={tempPhoto || editForm.photo} alt="Фото" className="teacher-photo" />
                ) : (
                  <div className="teacher-photo-placeholder">👨‍🏫</div>
                )}
                <label className="teacher-photo-upload">
                  📷 Загрузить фото
                  <input type="file" accept="image/*" onChange={handlePhotoUpload} hidden />
                </label>
              </>
            ) : (
              profile.photo ? (
                <img src={profile.photo} alt="Фото" className="teacher-photo" />
              ) : (
                <div className="teacher-photo-placeholder">👨‍🏫</div>
              )
            )}
            
            <div className="teacher-name">
              <h2>{profile.lastName} {profile.firstName} {profile.patronymic}</h2>
              <p className="teacher-position">{profile.position}</p>
              <p className="teacher-degree">{profile.degree}</p>
            </div>
          </div>

          <div className="teacher-contacts-card">
            <h3>📞 Контакты</h3>
            {isEditing ? (
              <>
                <div className="contact-row"><span>📧</span><input value={editForm.email} onChange={(e) => setEditForm({...editForm, email: e.target.value})} /></div>
                <div className="contact-row"><span>📱</span><input value={editForm.phone} onChange={(e) => setEditForm({...editForm, phone: e.target.value})} /></div>
                <div className="contact-row"><span>💬</span><input value={editForm.telegram} onChange={(e) => setEditForm({...editForm, telegram: e.target.value})} /></div>
                <div className="contact-row"><span>🏢</span><input value={editForm.office} onChange={(e) => setEditForm({...editForm, office: e.target.value})} /></div>
              </>
            ) : (
              <>
                <div className="contact-row"><span>📧</span>{profile.email}</div>
                <div className="contact-row"><span>📱</span>{profile.phone}</div>
                <div className="contact-row"><span>💬</span>{profile.telegram}</div>
                <div className="contact-row"><span>🏢</span>Аудитория {profile.office}</div>
              </>
            )}
          </div>

          <div className="teacher-stats-card">
            <h3>📊 Статистика</h3>
            <div className="stats-grid">
              <div className="stat-item"><div className="stat-value">{profile.experience}</div><div className="stat-label">лет опыта</div></div>
              <div className="stat-item"><div className="stat-value">{profile.courses.length}</div><div className="stat-label">курсов</div></div>
              <div className="stat-item"><div className="stat-value">{profile.courses.reduce((sum, c) => sum + c.students, 0)}</div><div className="stat-label">студентов</div></div>
              <div className="stat-item"><div className="stat-value">{profile.achievements.length}</div><div className="stat-label">достижений</div></div>
            </div>
          </div>
        </div>

        {/* Правая колонка - вкладки */}
        <div className="teacher-profile-right">
          <div className="teacher-tabs">
            <button className={`teacher-tab ${activeTab === 'info' ? 'active' : ''}`} onClick={() => setActiveTab('info')}>
              📋 Основное
            </button>
            <button className={`teacher-tab ${activeTab === 'courses' ? 'active' : ''}`} onClick={() => setActiveTab('courses')}>
              📚 Курсы
            </button>
            <button className={`teacher-tab ${activeTab === 'achievements' ? 'active' : ''}`} onClick={() => setActiveTab('achievements')}>
              🏆 Достижения
            </button>
            <button className={`teacher-tab ${activeTab === 'schedule' ? 'active' : ''}`} onClick={() => setActiveTab('schedule')}>
              📅 Расписание
            </button>
          </div>

          <div className="teacher-tab-content">
            {/* Вкладка: Основное */}
            {activeTab === 'info' && (
              <div className="teacher-card">
                <div className="card-header">
                  <h3>📋 Основная информация</h3>
                  {!isEditing && <button className="edit-btn" onClick={() => setIsEditing(true)}>✏️ Редактировать</button>}
                </div>
                
                {isEditing ? (
                  <>
                    <div className="form-row"><label>Фамилия</label><input value={editForm.lastName} onChange={(e) => setEditForm({...editForm, lastName: e.target.value})} /></div>
                    <div className="form-row"><label>Имя</label><input value={editForm.firstName} onChange={(e) => setEditForm({...editForm, firstName: e.target.value})} /></div>
                    <div className="form-row"><label>Отчество</label><input value={editForm.patronymic} onChange={(e) => setEditForm({...editForm, patronymic: e.target.value})} /></div>
                    <div className="form-row"><label>Кафедра</label><input value={editForm.department} onChange={(e) => setEditForm({...editForm, department: e.target.value})} /></div>
                    <div className="form-row"><label>Должность</label><input value={editForm.position} onChange={(e) => setEditForm({...editForm, position: e.target.value})} /></div>
                    <div className="form-row"><label>Ученая степень</label><input value={editForm.degree} onChange={(e) => setEditForm({...editForm, degree: e.target.value})} /></div>
                    <div className="form-row"><label>Стаж (лет)</label><input type="number" value={editForm.experience} onChange={(e) => setEditForm({...editForm, experience: parseInt(e.target.value)})} /></div>
                    <div className="edit-actions">
                      <button className="save-btn" onClick={handleSave}>💾 Сохранить</button>
                      <button className="cancel-btn" onClick={() => { setIsEditing(false); setEditForm(profile); }}>❌ Отмена</button>
                    </div>
                  </>
                ) : (
                  <div className="info-grid">
                    <div className="info-row"><span>📚 Кафедра:</span><strong>{profile.department}</strong></div>
                    <div className="info-row"><span>🎓 Должность:</span><strong>{profile.position}</strong></div>
                    <div className="info-row"><span>📜 Ученая степень:</span><strong>{profile.degree}</strong></div>
                    <div className="info-row"><span>⏱️ Стаж работы:</span><strong>{profile.experience} лет</strong></div>
                  </div>
                )}
              </div>
            )}

            {/* Вкладка: Курсы */}
            {activeTab === 'courses' && (
              <div className="teacher-card">
                <div className="card-header"><h3>📚 Ведомые курсы</h3></div>
                <div className="courses-list">
                  {profile.courses.map(course => (
                    <div key={course.id} className="course-card">
                      <div className="course-icon">📖</div>
                      <div className="course-info">
                        <h4>{course.name}</h4>
                        <p>Группа: {course.group} | Часов: {course.hours} | Студентов: {course.students}</p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}

            {/* Вкладка: Достижения */}
            {activeTab === 'achievements' && (
              <div className="teacher-card">
                <div className="card-header">
                  <h3>🏆 Достижения и награды</h3>
                  {isEditing && <button className="add-btn" onClick={handleAddAchievement}>+ Добавить</button>}
                </div>
                <div className="achievements-list">
                  {editForm.achievements.map(ach => (
                    <div key={ach.id} className="achievement-card">
                      <div className="achievement-icon">{ach.icon}</div>
                      <div className="achievement-info">
                        <h4>{ach.title}</h4>
                        <p>{ach.date}</p>
                      </div>
                      {isEditing && <button className="remove-btn" onClick={() => handleRemoveAchievement(ach.id)}>🗑️</button>}
                    </div>
                  ))}
                  {profile.achievements.length === 0 && <div className="empty-state">Нет добавленных достижений</div>}
                </div>
              </div>
            )}

            {/* Вкладка: Расписание */}
            {activeTab === 'schedule' && (
              <div className="teacher-card">
                <div className="card-header"><h3>📅 Расписание занятий</h3></div>
                <div className="schedule-list">
                  {profile.schedule.map((item, idx) => (
                    <div key={idx} className="schedule-card">
                      <div className="schedule-day">{item.day}</div>
                      <div className="schedule-time">{item.time}</div>
                      <div className="schedule-subject">{item.subject}</div>
                      <div className="schedule-group">{item.group}</div>
                      <div className="schedule-room">{item.room}</div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default TeacherProfile;