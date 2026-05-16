import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import './Sidebar.css';

function SidebarAdmin({ onLogout }) {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const menuItems = [
    { path: '/admin/dashboard', icon: '📊', name: 'Панель управления' },
    { path: '/admin/students', icon: '👨‍🎓', name: 'Студенты' },
    { path: '/admin/schedule', icon: '📅', name: 'Расписание' },
    { path: '/admin/library', icon: '📚', name: 'Библиотека' },
    { path: '/admin/news', icon: '📰', name: 'Новости' },
    { path: '/admin/events', icon: '🎮', name: 'ДПО' },
    { path: '/admin/achievements', icon: '🏆', name: 'Ачивки' },
    { path: '/admin/exam', icon: '🎓', name: 'Экзамен' },
    { path: '/admin/challenges', icon: '🎯', name: 'Вызовы' },
    { path: '/admin/games', icon: '🎮', name: 'Мини-игры' },
    { path: '/admin/settings', icon: '⚙️', name: 'Настройки' }
  ];

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };

  return (
    <>
      <button className="mobile-menu-btn" onClick={toggleMobileMenu}>
        <span className="hamburger-icon">☰</span>
      </button>

      {isMobileMenuOpen && (
        <div className="mobile-overlay" onClick={toggleMobileMenu}></div>
      )}

      <div className={`sidebar ${isMobileMenuOpen ? 'mobile-open' : ''}`}>
        <div className="sidebar-logo">
          <div className="sidebar-logo-icon">📱</div>
          <div className="sidebar-logo-text">AdminPanel</div>
          <div className="sidebar-role-badge">👑 Админ</div>
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

export default SidebarAdmin;