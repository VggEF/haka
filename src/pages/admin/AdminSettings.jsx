import React, { useState, useEffect } from 'react';
import './AdminSettings.css';

function AdminSettings() {
  const [settings, setSettings] = useState({
    theme: 'light',
    notifications: true,
    language: 'ru'
  });

  useEffect(() => {
    const saved = localStorage.getItem('adminSettings');
    if (saved) setSettings(JSON.parse(saved));
  }, []);

  const saveSettings = (key, value) => {
    const updated = { ...settings, [key]: value };
    setSettings(updated);
    localStorage.setItem('adminSettings', JSON.stringify(updated));
    alert(`Настройка сохранена!`);
  };

  return (
    <div className="admin-page">
      <h2>⚙️ Настройки</h2>
      
      <div className="settings-list">
        <div className="setting-item">
          <span>🌙 Тема оформления</span>
          <select value={settings.theme} onChange={(e) => saveSettings('theme', e.target.value)}>
            <option value="light">Светлая</option>
            <option value="dark">Темная</option>
          </select>
        </div>
        
        <div className="setting-item">
          <span>🔔 Уведомления</span>
          <select value={settings.notifications} onChange={(e) => saveSettings('notifications', e.target.value === 'true')}>
            <option value="true">Включены</option>
            <option value="false">Выключены</option>
          </select>
        </div>
        
        <div className="setting-item">
          <span>🌐 Язык интерфейса</span>
          <select value={settings.language} onChange={(e) => saveSettings('language', e.target.value)}>
            <option value="ru">Русский</option>
            <option value="en">English</option>
          </select>
        </div>
      </div>
    </div>
  );
}

export default AdminSettings;