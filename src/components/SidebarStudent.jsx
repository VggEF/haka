import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import './Sidebar.css';

function SidebarStudent({ onLogout, userLogin }) {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const menuItems = [
    { path: '/profile', icon: '👤', name: 'Профиль' },
    { path: '/schedule', icon: '📅', name: 'Расписание' },
    { path: '/library', icon: '📚', name: 'Библиотека' },
    { path: '/news', icon: '📰', name: 'Новости' },
    { path: '/entertainment', icon: '🎮', name: 'ДПО и развлечения' },
    { path: '/achievements', icon: '🏆', name: 'Ачивки' },
    { path: '/skill-tree', icon: '🌳', name: 'Дерево навыков' },        // ИСПРАВЛЕНО
    { path: '/deadlines', icon: '⏰', name: 'Дедлайны' },
    { path: '/checklist', icon: '📋', name: 'Чек-лист' },
    { path: '/shop', icon: '🛒', name: 'Магазин' },
    { path: '/exam-generator', icon: '🎓', name: 'Экзамен' },           // ИСПРАВЛЕНО
    { path: '/challenges', icon: '🎯', name: 'Вызовы' },
    { path: '/mini-games', icon: '🎮', name: 'Мини-игры' },              // ИСПРАВЛЕНО
    { path: '/planner', icon: '📅', name: 'Планер' },
    { path: '/settings', icon: '⚙️', name: 'Настройки' }
  ];

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };

  return (
    <>
      {/* Мобильная кнопка-гамбургер */}
      <button className="mobile-menu-btn" onClick={toggleMobileMenu}>
        <span className="hamburger-icon">☰</span>
      </button>

      {/* Оверлей для мобильного меню */}
      {isMobileMenuOpen && (
        <div className="mobile-overlay" onClick={toggleMobileMenu}></div>
      )}

      {/* Сайдбар */}
      <div className={`sidebar ${isMobileMenuOpen ? 'mobile-open' : ''}`}>
        <div className="sidebar-logo">
          <div className="sidebar-logo-icon">📱</div>
          <div className="sidebar-logo-text">StudentApp</div>
          <div className="sidebar-role-badge">🎓 Студент</div>
          <button className="mobile-close-btn" onClick={toggleMobileMenu}>✕</button>
        </div>
        
        <nav className="sidebar-nav">
          {menuItems.map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) => 
                isActive ? 'sidebar-nav-item active' : 'sidebar-nav-item'
              }
              onClick={() => setIsMobileMenuOpen(false)}
            >
              <span className="sidebar-nav-icon">{item.icon}</span>
              <span>{item.name}</span>
            </NavLink>
          ))}
        </nav>

        <button onClick={onLogout} className="sidebar-logout">
          <span>🚪</span>
          <span>Выйти</span>
        </button>
      </div>
    </>
  );
}

export default SidebarStudent;