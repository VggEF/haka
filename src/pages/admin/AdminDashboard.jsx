import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { DEMO_USERS, ROLES } from '../../utils/roles';
import './AdminDashboard.css';

function AdminDashboard() {
  const navigate = useNavigate();
  const [stats, setStats] = useState({});
  const [recentActivity, setRecentActivity] = useState([]);

  useEffect(() => {
    const students = DEMO_USERS.filter(u => u.role === ROLES.STUDENT).length;
    const teachers = DEMO_USERS.filter(u => u.role === ROLES.TEACHER).length;
    const staff = DEMO_USERS.filter(u => u.role === ROLES.STAFF).length;
    
    setStats({
      totalUsers: DEMO_USERS.length,
      students,
      teachers,
      staff,
      books: JSON.parse(localStorage.getItem('libraryBooks') || '[]').length,
      news: JSON.parse(localStorage.getItem('newsItems') || '[]').length,
      events: JSON.parse(localStorage.getItem('entertainmentEvents') || '[]').length,
      achievements: JSON.parse(localStorage.getItem('customAchievements') || '[]').length,
      questions: JSON.parse(localStorage.getItem('examQuestions') || '[]').length,
      games: JSON.parse(localStorage.getItem('miniGames') || '[]').length
    });

    // Демо-активность
    setRecentActivity([
      { id: 1, action: "Добавлен новый студент", user: "Администратор", time: "5 минут назад", icon: "👨‍🎓" },
      { id: 2, action: "Обновлено расписание", user: "Администратор", time: "1 час назад", icon: "📅" },
      { id: 3, action: "Добавлена новая книга", user: "Библиотекарь", time: "3 часа назад", icon: "📚" },
      { id: 4, action: "Создано мероприятие", user: "Организатор", time: "Вчера", icon: "🎉" }
    ]);
  }, []);

  const quickLinks = [
    { title: "Студенты", icon: "👨‍🎓", count: stats.students, path: "/admin/students", color: "#667eea" },
    { title: "Расписание", icon: "📅", count: stats.schedule, path: "/admin/schedule", color: "#f093fb" },
    { title: "Библиотека", icon: "📚", count: stats.books, path: "/admin/library", color: "#4facfe" },
    { title: "Новости", icon: "📰", count: stats.news, path: "/admin/news", color: "#43e97b" },
    { title: "ДПО", icon: "🎮", count: stats.events, path: "/admin/events", color: "#fa709a" },
    { title: "Ачивки", icon: "🏆", count: stats.achievements, path: "/admin/achievements", color: "#f5576c" },
    { title: "Экзамен", icon: "🎓", count: stats.questions, path: "/admin/exam", color: "#ff9800" },
    { title: "Мини-игры", icon: "🎮", count: stats.games, path: "/admin/games", color: "#9c27b0" },
    { title: "Настройки", icon: "⚙️", count: null, path: "/admin/settings", color: "#607d8b" }
  ];

  return (
    <div className="admin-dashboard">
      <div className="dashboard-header">
        <div className="dashboard-title">
          <div className="dashboard-title-icon">📊</div>
          <h1>Панель управления</h1>
        </div>
        <div className="dashboard-subtitle">Добро пожаловать в админ-панель. Здесь вы можете управлять всем контентом.</div>
      </div>

      {/* Быстрые ссылки */}
      <div className="quick-links-section">
        <h2 className="section-title">🚀 Быстрые ссылки</h2>
        <div className="quick-links-grid">
          {quickLinks.map((link, index) => (
            <div 
              key={index} 
              className="quick-link-card" 
              onClick={() => navigate(link.path)}
              style={{ borderBottom: `3px solid ${link.color}` }}
            >
              <div className="quick-link-icon" style={{ background: link.color }}>{link.icon}</div>
              <div className="quick-link-info">
                <h3>{link.title}</h3>
                {link.count !== undefined && <p>{link.count} элементов</p>}
              </div>
              <div className="quick-link-arrow">→</div>
            </div>
          ))}
        </div>
      </div>

      {/* Последняя активность */}
      <div className="recent-activity">
        <h2 className="section-title">🕐 Последняя активность</h2>
        <div className="activity-list">
          {recentActivity.map(activity => (
            <div key={activity.id} className="activity-item">
              <div className="activity-icon">{activity.icon}</div>
              <div className="activity-info">
                <div className="activity-text">{activity.action}</div>
                <div className="activity-meta">{activity.user} • {activity.time}</div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default AdminDashboard;