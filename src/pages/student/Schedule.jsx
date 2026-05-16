import React, { useState, useEffect, useRef } from 'react';
import './Schedule.css';

// Только группы ИВИТШ
const IVITSH_GROUPS = [
  { name: "22-ИБбо-6", id: 8190, kurs: 4 },
  { name: "22-ИСбо-1", id: 8195, kurs: 4 },
  { name: "22-ИСбо-2", id: 8196, kurs: 4 },
  { name: "22-ИСбо-3", id: 8193, kurs: 4 },
  { name: "22-ИСбо-4", id: 8197, kurs: 4 },
  { name: "22-ПМбо-1", id: 8216, kurs: 4 },
  { name: "23-ИБбо-6", id: 8259, kurs: 3 },
  { name: "23-ИСбо-1", id: 8265, kurs: 3 },
  { name: "23-ИСбо-2", id: 8266, kurs: 3 },
  { name: "23-ИСбо-3", id: 8263, kurs: 3 },
  { name: "23-ИСбо-4", id: 8267, kurs: 3 },
  { name: "23-ПМбо-1", id: 8289, kurs: 3 },
  { name: "24-ИБбо-6", id: 8346, kurs: 2 },
  { name: "24-ИМбо-3", id: 8350, kurs: 2 },
  { name: "24-ИСбо-1", id: 8540, kurs: 2 },
  { name: "24-ИСбо-2", id: 8529, kurs: 2 },
  { name: "24-ИСмо-2", id: 8353, kurs: 2 },
  { name: "24-ИТбо-4", id: 8354, kurs: 2 },
  { name: "25-ИБбо-6", id: 8446, kurs: 1 },
  { name: "25-ИСбо-1", id: 8448, kurs: 1 },
  { name: "25-ИСбо-2", id: 8451, kurs: 1 },
  { name: "25-ИСбо-3", id: 8449, kurs: 1 },
  { name: "25-ИСбо-4", id: 8453, kurs: 1 },
  { name: "25-ИСмо-2", id: 8452, kurs: 1 },
  { name: "25-ПМбо-1", id: 8474, kurs: 1 },
];

// Время пар
const LESSON_TIMES = {
  1: "08:00-09:30",
  2: "09:40-11:10",
  3: "11:20-12:50",
  4: "13:00-14:30",
  5: "14:40-16:10",
  6: "16:20-17:50",
  7: "18:00-19:30",
};

const getLessonIcon = (discipline) => {
  if (discipline.includes('лек')) return '📖';
  if (discipline.includes('лаб')) return '🔬';
  if (discipline.includes('пр.')) return '💻';
  return '📚';
};

function Schedule({ userLogin = "23-ПМбо-014" }) {
  const [searchTerm, setSearchTerm] = useState('');
  const [filteredGroups, setFilteredGroups] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedGroup, setSelectedGroup] = useState(null);
  const [schedule, setSchedule] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  
  const [modalOpen, setModalOpen] = useState(false);
  const [selectedLesson, setSelectedLesson] = useState(null);
  const [taskText, setTaskText] = useState('');
  
  const searchRef = useRef(null);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (searchRef.current && !searchRef.current.contains(event.target)) {
        setShowDropdown(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  useEffect(() => {
    if (searchTerm.length >= 1) {
      const filtered = IVITSH_GROUPS.filter(group => 
        group.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setFilteredGroups(filtered);
      setShowDropdown(true);
    } else {
      setFilteredGroups([]);
      setShowDropdown(false);
    }
  }, [searchTerm]);

  const fetchSchedule = async (group) => {
    setLoading(true);
    setError('');
    setSelectedGroup(group);
    setSearchTerm(group.name);
    setShowDropdown(false);
    
    const today = new Date().toISOString().split('T')[0];

    try {
      const response = await fetch(
        `https://eios.kosgos.ru/api/Rasp?idGroup=${group.id}&sdate=${today}`,
        {
          headers: {
            'accept': 'application/json',
            'client-version': '2026-04-10T12:06:09.286Z',
          }
        }
      );

      if (!response.ok) throw new Error('Ошибка загрузки');

      const data = await response.json();
      
      if (data.state === 1) {
        setSchedule(data.data?.rasp || []);
      } else {
        setError('Расписание не найдено');
      }
    } catch (err) {
      setError('Не удалось загрузить расписание');
    } finally {
      setLoading(false);
    }
  };

  const openModal = (lesson) => {
    setSelectedLesson(lesson);
    setModalOpen(true);
    setTaskText('');
  };

  const closeModal = () => {
    setModalOpen(false);
    setSelectedLesson(null);
    setTaskText('');
  };

  const handleSubmit = () => {
    if (!taskText.trim()) {
      alert('Напишите задание');
      return;
    }
    alert(`Задание по предмету "${selectedLesson?.дисциплина}" отправлено!\n\nВаш ответ: ${taskText}`);
    closeModal();
  };

  const getDaysWithLessons = () => {
    const daysMap = new Map();
    
    schedule.forEach(lesson => {
      const date = lesson.дата.split('T')[0];
      if (!daysMap.has(date)) {
        daysMap.set(date, []);
      }
      daysMap.get(date).push(lesson);
    });
    
    const sortedDays = Array.from(daysMap.entries()).sort((a, b) => 
      new Date(a[0]) - new Date(b[0])
    );
    
    return sortedDays.map(([date, lessons]) => ({
      date: date,
      dayOfWeek: new Date(date).getDay(),
      lessons: lessons.sort((a, b) => a.номерЗанятия - b.номерЗанятия)
    }));
  };

  const getDayName = (dayOfWeek) => {
    const days = ['Воскресенье', 'Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота'];
    return days[dayOfWeek];
  };

  const formatDate = (dateStr) => {
    const date = new Date(dateStr);
    return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'long' });
  };

  const isToday = (dateStr) => {
    const today = new Date().toISOString().split('T')[0];
    return dateStr === today;
  };

  const daysWithLessons = getDaysWithLessons();

  useEffect(() => {
    if (userLogin) {
      const parts = userLogin.split('-');
      if (parts.length >= 2) {
        const searchQuery = `${parts[0]}-${parts[1]}`;
        setSearchTerm(searchQuery);
        
        const foundGroup = IVITSH_GROUPS.find(g => g.name.startsWith(searchQuery));
        if (foundGroup) {
          fetchSchedule(foundGroup);
        }
      }
    }
  }, [userLogin]);

  return (
    <div className="schedule-container">
      <h2 className="schedule-title">📅 Расписание занятий</h2>
      
      <div className="schedule-search" ref={searchRef}>
        <input
          type="text"
          className="schedule-search-input"
          placeholder="🔍 Введите название группы (например: ПМбо, ИСбо)"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          onFocus={() => searchTerm.length >= 1 && setShowDropdown(true)}
        />
        
        {showDropdown && filteredGroups.length > 0 && (
          <div className="search-results">
            {filteredGroups.map(group => (
              <div key={group.id} className="search-result-item" onClick={() => fetchSchedule(group)}>
                <div className="result-name">{group.name}</div>
                <div className="result-course">{group.kurs} курс</div>
              </div>
            ))}
          </div>
        )}
        
        {showDropdown && searchTerm.length >= 1 && filteredGroups.length === 0 && (
          <div className="search-no-results">Группы не найдены</div>
        )}
      </div>

      {selectedGroup && (
        <div className="group-info">
          📚 <strong>{selectedGroup.name}</strong> ({selectedGroup.kurs} курс)
          {userLogin && <span className="user-login"> | {userLogin}</span>}
        </div>
      )}

      {loading && (
        <div className="schedule-loading">
          <div className="loading-spinner"></div>
          <span>Загрузка расписания...</span>
        </div>
      )}

      {error && <div className="schedule-error">⚠️ {error}</div>}

      {!loading && !error && selectedGroup && daysWithLessons.length > 0 && (
        <div className="schedule-days">
          {daysWithLessons.map((day, idx) => (
            <div key={idx} className={`day-card ${isToday(day.date) ? 'today-card' : ''}`}>
              <div className="day-card-header">
                <div className="day-name">{getDayName(day.dayOfWeek)}</div>
                <div className="day-date">{formatDate(day.date)}</div>
                {isToday(day.date) && <div className="today-badge">Сегодня</div>}
              </div>
              <div className="day-card-body">
                {day.lessons.length > 0 ? (
                  day.lessons.map((lesson, lessonIdx) => (
                    <div 
                      key={lessonIdx} 
                      className="lesson-card"
                      onClick={() => openModal(lesson)}
                    >
                      <div className="lesson-time">
                        {LESSON_TIMES[lesson.номерЗанятия]}
                      </div>
                      <div className="lesson-content">
                        <div className="lesson-title">
                          <span className="lesson-icon">{getLessonIcon(lesson.дисциплина)}</span>
                          <span className="lesson-name">{lesson.дисциплина}</span>
                        </div>
                        <div className="lesson-details">
                          <div className="lesson-teacher">
                            <span className="detail-icon">👨‍🏫</span>
                            <span>{lesson.преподаватель}</span>
                          </div>
                          <div className="lesson-room">
                            <span className="detail-icon">📍</span>
                            <span>{lesson.аудитория}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  ))
                ) : (
                  <div className="no-lessons-card">
                    <span className="no-lessons-icon">📭</span>
                    <span>Нет пар</span>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      )}

      {!loading && !error && selectedGroup && daysWithLessons.length === 0 && (
        <div className="schedule-empty">
          <div className="empty-icon">📭</div>
          <div>Нет расписания на текущую неделю</div>
        </div>
      )}

      {/* МОДАЛЬНОЕ ОКНО 50% ширины и высоты */}
      {modalOpen && selectedLesson && (
        <div className="modal-overlay" onClick={closeModal}>
          <div className="modal-container" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>📝 {selectedLesson.дисциплина}</h3>
              <button className="modal-close-btn" onClick={closeModal}>✕</button>
            </div>
            <div className="modal-body">
              <div className="lesson-details-modal">
                <div className="detail-row">
                  <span className="detail-label">📅 Дата:</span>
                  <span>{formatDate(selectedLesson.дата)}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">⏰ Время:</span>
                  <span>{LESSON_TIMES[selectedLesson.номерЗанятия]}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">👨‍🏫 Преподаватель:</span>
                  <span>{selectedLesson.преподаватель}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">📍 Аудитория:</span>
                  <span>{selectedLesson.аудитория}</span>
                </div>
              </div>
              
              <div className="task-section-modal">
                <label className="task-label-modal">📌 Задание преподавателя:</label>
                <textarea
                  className="task-textarea-modal"
                  placeholder="Введите текст задания..."
                  value={taskText}
                  onChange={(e) => setTaskText(e.target.value)}
                  rows={5}
                />
              </div>
              
              <button className="submit-btn-modal" onClick={handleSubmit}>
                📤 Отправить задание
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Schedule;