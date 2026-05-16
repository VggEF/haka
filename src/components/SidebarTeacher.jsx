import React, { useState, useEffect } from 'react';
import { NavLink } from 'react-router-dom';
import './Sidebar.css';

function SidebarTeacher({ onLogout }) {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const menuItems = [
    { path: '/teacher/dashboard', icon: '📊', name: 'Панель' },
    { path: '/teacher/profile', icon: '👤', name: 'Профиль' },
    { path: '/teacher/schedule', icon: '📅', name: 'Расписание' },
    { path: '/teacher/students', icon: '👨‍🎓', name: 'Студенты' },
    { path: '/teacher/achievements', icon: '🏆', name: 'Ачивки' },
    { path: '/teacher/deadlines', icon: '⏰', name: 'Дедлайны' }
  ];

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };

  useEffect(() => {
    if (isMobileMenuOpen) {
      document.body.classList.add('no-scroll');
    } else {
      document.body.classList.remove('no-scroll');
    }
    return () => document.body.classList.remove('no-scroll');
  }, [isMobileMenuOpen]);

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
          <div className="sidebar-logo-text">TeacherApp</div>
          <div className="sidebar-role-badge">👨‍🏫 Преподаватель</div>
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

export default SidebarTeacher;