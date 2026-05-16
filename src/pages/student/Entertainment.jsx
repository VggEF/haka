import React, { useState } from 'react';
import './Entertainment.css';

// Демо-данные для ДПО и развлечений
const DEMO_ACTIVITIES = [
  {
    id: 1,
    title: "🏀 Турнир по баскетболу",
    shortText: "Открытый турнир среди студенческих команд. Призы и грамоты победителям.",
    fullText: "Приглашаем студентов принять участие в ежегодном турнире по баскетболу среди факультетов. Турнир пройдет в спортивном комплексе университета. Команды могут состоять из 5-10 человек. Регистрация команд до 1 июня. Победители получат кубок, медали и ценные призы. Матчи будут судить профессиональные судьи. Все участники получат сертификаты.",
    date: "2026-06-20",
    time: "15:00",
    type: "sport",
    category: "Спорт",
    location: "Спорткомплекс",
    contact: "sport@kosgos.ru",
    image: "🏀",
    availableSpots: "15",
    price: "Бесплатно",
    organizer: "Спортивный клуб"
  },
  {
    id: 2,
    title: "🎮 Киберспортивный турнир",
    shortText: "Соревнования по CS:GO и Dota 2 среди студентов",
    fullText: "Киберспортивный клуб университета проводит турнир по CS:GO и Dota 2. Призовой фонд - 50 000₽. Турнир пройдет в компьютерном классе (ауд. Б-208). Команды из 5 человек. Обязательна предварительная регистрация. Все участники получат скины и внутриигровые предметы. Трансляция финала будет на большом экране в холле.",
    date: "2026-06-25",
    time: "10:00",
    type: "student",
    category: "Киберспорт",
    location: "Компьютерный класс Б-208",
    contact: "esports@kosgos.ru",
    image: "🎮",
    availableSpots: "8 команд",
    price: "Бесплатно",
    organizer: "Киберспортивный клуб"
  },
  {
    id: 3,
    title: "🎬 Киноклуб: «Матрица»",
    shortText: "Просмотр фильма с обсуждением в кругу единомышленников",
    fullText: "Студенческий киноклуб приглашает на просмотр культового фильма «Матрица». После просмотра состоится обсуждение с преподавателем философии. Можно будет высказать свое мнение и поучаствовать в дискуссии. Чай и печенье включены. Приходите с друзьями!",
    date: "2026-05-28",
    time: "18:00",
    type: "student",
    category: "Кино",
    location: "Актовый зал",
    contact: "cinema@kosgos.ru",
    image: "🎬",
    availableSpots: "50",
    price: "Бесплатно",
    organizer: "Студенческий клуб"
  },
  {
    id: 4,
    title: "🔬 Мастер-класс по робототехнике для школьников",
    shortText: "Практическое занятие по сборке и программированию роботов",
    fullText: "Факультет ИВИТШ приглашает школьников 9-11 классов на мастер-класс по робототехнике. Вы узнаете основы сборки роботов на платформе Arduino, научитесь программировать датчики и моторы. В конце занятия - мини-соревнование роботов. Каждый участник получит сертификат. Количество мест ограничено!",
    date: "2026-06-05",
    time: "14:00",
    type: "school",
    category: "Образование",
    location: "Лаборатория робототехники Б-301",
    contact: "robotics@kosgos.ru",
    image: "🔬",
    availableSpots: "20",
    price: "500₽ (материалы)",
    organizer: "Факультет ИВИТШ"
  },
  {
    id: 5,
    title: "⚽ Футбольный матч",
    shortText: "Товарищеский матч между студентами и преподавателями",
    fullText: "Приглашаем студентов поддержать команду университета в товарищеском матче против команды преподавателей. Матч пройдет на стадионе университета. Вход свободный. Желающие могут записаться в команду студентов - требуется предварительная регистрация.",
    date: "2026-06-12",
    time: "16:00",
    type: "sport",
    category: "Спорт",
    location: "Стадион",
    contact: "football@kosgos.ru",
    image: "⚽",
    availableSpots: "22 (игроки)",
    price: "Бесплатно",
    organizer: "Спортивный клуб"
  },
  {
    id: 6,
    title: "🎨 Вечер настольных игр",
    shortText: "Игротека для студентов: Мафия, UNO, Монополия и другие",
    fullText: "Студенческий совет организует вечер настольных игр. Будут представлены: Мафия, UNO, Монополия, Имаджинариум, Каркассон и другие. Можно приносить свои игры. Будет чай и угощения. Отличная возможность познакомиться с новыми людьми и отдохнуть после учебы.",
    date: "2026-05-29",
    time: "17:00",
    type: "student",
    category: "Игры",
    location: "Коворкинг",
    contact: "club@kosgos.ru",
    image: "🎲",
    availableSpots: "40",
    price: "Бесплатно",
    organizer: "Студенческий совет"
  },
  {
    id: 7,
    title: "🧘‍♀️ Йога на свежем воздухе",
    shortText: "Занятия йогой для начинающих в парке",
    fullText: "Спортивный клуб приглашает на занятия йогой на свежем воздухе. С собой необходимо иметь коврик и удобную одежду. Занятия подходят для любого уровня подготовки. Преподаватель - сертифицированный инструктор. Укрепление здоровья, снятие стресса, хорошее настроение гарантированы!",
    date: "2026-06-03",
    time: "09:00",
    type: "sport",
    category: "Спорт",
    location: "Городской парк",
    contact: "yoga@kosgos.ru",
    image: "🧘",
    availableSpots: "30",
    price: "Бесплатно",
    organizer: "Спортивный клуб"
  },
  {
    id: 8,
    title: "🚀 День карьеры IT",
    shortText: "Встреча с представителями IT-компаний, мастер-классы, стажировки",
    fullText: "Крупнейшее карьерное мероприятие для студентов IT-специальностей. Приглашены компании: Яндекс, Google, EPAM, Luxoft. В программе: лекции, мастер-классы, Q&A сессии, возможность пройти собеседование на стажировку. Предварительная регистрация обязательна. При себе иметь резюме.",
    date: "2026-06-18",
    time: "11:00",
    type: "school",
    category: "Карьера",
    location: "Конференц-зал",
    contact: "career@kosgos.ru",
    image: "💼",
    availableSpots: "200",
    price: "Бесплатно",
    organizer: "Центр карьеры"
  },
];

function Entertainment() {
  const [searchTerm, setSearchTerm] = useState('');
  const [activeTab, setActiveTab] = useState('all'); // 'all', 'sport', 'school', 'student'
  const [selectedActivity, setSelectedActivity] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [registeredActivities, setRegisteredActivities] = useState([]);
  const [registrationMessage, setRegistrationMessage] = useState('');

  // Фильтрация мероприятий
  const getFilteredActivities = () => {
    let filtered = DEMO_ACTIVITIES;
    
    // Фильтр по типу вкладки
    if (activeTab === 'sport') {
      filtered = filtered.filter(item => item.type === 'sport');
    } else if (activeTab === 'school') {
      filtered = filtered.filter(item => item.type === 'school');
    } else if (activeTab === 'student') {
      filtered = filtered.filter(item => item.type === 'student');
    }
    
    // Фильтр по поиску
    if (searchTerm.trim()) {
      filtered = filtered.filter(item =>
        item.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
        item.shortText.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }
    
    // Сортировка по дате (сначала ближайшие)
    return filtered.sort((a, b) => new Date(a.date) - new Date(b.date));
  };

  const filteredActivities = getFilteredActivities();

  // Открытие модального окна
  const openModal = (activity) => {
    setSelectedActivity(activity);
    setIsModalOpen(true);
    setRegistrationMessage('');
  };

  // Закрытие модального окна
  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedActivity(null);
    setRegistrationMessage('');
  };

  // Запись на мероприятие
  const handleRegister = () => {
    if (registeredActivities.includes(selectedActivity.id)) {
      setRegistrationMessage('❌ Вы уже записаны на это мероприятие!');
      setTimeout(() => setRegistrationMessage(''), 3000);
      return;
    }
    
    setRegisteredActivities([...registeredActivities, selectedActivity.id]);
    setRegistrationMessage(`✅ Вы успешно записаны на "${selectedActivity.title}"!`);
    setTimeout(() => setRegistrationMessage(''), 3000);
  };

  // Форматирование даты
  const formatDate = (dateStr) => {
    const date = new Date(dateStr);
    const today = new Date();
    const tomorrow = new Date(today);
    tomorrow.setDate(tomorrow.getDate() + 1);
    
    if (date.toDateString() === today.toDateString()) {
      return 'Сегодня';
    }
    if (date.toDateString() === tomorrow.toDateString()) {
      return 'Завтра';
    }
    return date.toLocaleDateString('ru-RU', { day: 'numeric', month: 'long' });
  };

  // Проверка, записан ли пользователь
  const isRegistered = (id) => registeredActivities.includes(id);

  // Получение иконки для категории
  const getCategoryIcon = (category) => {
    const icons = {
      'Спорт': '⚽',
      'Киберспорт': '🎮',
      'Кино': '🎬',
      'Образование': '📚',
      'Игры': '🎲',
      'Карьера': '💼',
    };
    return icons[category] || '🎯';
  };

  return (
    <div className="entertainment-container">
      <h2 className="entertainment-title">🎯 ДПО и развлечения</h2>
      
      {/* Поиск */}
      <div className="entertainment-search">
        <input
          type="text"
          className="search-input"
          placeholder="🔍 Поиск мероприятий, секций, событий..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      
      {/* Вкладки */}
      <div className="entertainment-tabs">
        <button 
          className={`tab-btn ${activeTab === 'all' ? 'active' : ''}`}
          onClick={() => setActiveTab('all')}
        >
          🎯 Все
        </button>
        <button 
          className={`tab-btn ${activeTab === 'sport' ? 'active' : ''}`}
          onClick={() => setActiveTab('sport')}
        >
          ⚽ Спорт
        </button>
        <button 
          className={`tab-btn ${activeTab === 'school' ? 'active' : ''}`}
          onClick={() => setActiveTab('school')}
        >
          🎓 Школьникам
        </button>
        <button 
          className={`tab-btn ${activeTab === 'student' ? 'active' : ''}`}
          onClick={() => setActiveTab('student')}
        >
          🎉 Студенческие мероприятия
        </button>
      </div>
      
      {/* Список мероприятий */}
      <div className="entertainment-list">
        {filteredActivities.length === 0 ? (
          <div className="no-activities">
            <span className="no-activities-icon">🎯</span>
            <p>Мероприятий не найдено</p>
          </div>
        ) : (
          filteredActivities.map((activity) => (
            <div 
              key={activity.id} 
              className={`activity-card ${activity.type}-card`}
              onClick={() => openModal(activity)}
            >
              <div className="activity-card-icon">{activity.image}</div>
              <div className="activity-card-content">
                <div className="activity-card-header">
                  <span className="activity-category">
                    {getCategoryIcon(activity.category)} {activity.category}
                  </span>
                  <span className="activity-date">📅 {formatDate(activity.date)}</span>
                  <span className="activity-time">⏰ {activity.time}</span>
                </div>
                <h3 className="activity-card-title">{activity.title}</h3>
                <p className="activity-card-text">{activity.shortText}</p>
                <div className="activity-card-footer">
                  <span className="activity-location">📍 {activity.location}</span>
                  <span className="activity-price">{activity.price}</span>
                </div>
                {isRegistered(activity.id) && (
                  <span className="registered-badge">✅ Вы записаны</span>
                )}
              </div>
              <div className="activity-card-arrow">→</div>
            </div>
          ))
        )}
      </div>
      
      {/* Модальное окно */}
      {isModalOpen && selectedActivity && (
        <div className="modal-overlay" onClick={closeModal}>
          <div className="modal-container" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <div className="modal-icon">{selectedActivity.image}</div>
              <h3>{selectedActivity.title}</h3>
              <button className="modal-close" onClick={closeModal}>✕</button>
            </div>
            <div className="modal-body">
              <div className="modal-meta">
                <span className="meta-category">
                  {getCategoryIcon(selectedActivity.category)} {selectedActivity.category}
                </span>
                <span className="meta-date">📅 {formatDate(selectedActivity.date)}</span>
                <span className="meta-time">⏰ {selectedActivity.time}</span>
                <span className="meta-location">📍 {selectedActivity.location}</span>
              </div>
              
              <div className="modal-full-text">
                <p>{selectedActivity.fullText}</p>
              </div>
              
              <div className="modal-details">
                <div className="detail-row">
                  <span className="detail-label">👥 Организатор:</span>
                  <span className="detail-value">{selectedActivity.organizer}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">📊 Свободных мест:</span>
                  <span className="detail-value">{selectedActivity.availableSpots}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">💰 Стоимость:</span>
                  <span className="detail-value">{selectedActivity.price}</span>
                </div>
                <div className="detail-row">
                  <span className="detail-label">📧 Контакты:</span>
                  <span className="detail-value">{selectedActivity.contact}</span>
                </div>
              </div>
              
              {registrationMessage && (
                <div className="registration-message">
                  {registrationMessage}
                </div>
              )}
              
              <button 
                className={`register-btn ${isRegistered(selectedActivity.id) ? 'registered' : ''}`}
                onClick={handleRegister}
                disabled={isRegistered(selectedActivity.id)}
              >
                {isRegistered(selectedActivity.id) ? '✅ Вы уже записаны' : '📝 Записаться на мероприятие'}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Entertainment;