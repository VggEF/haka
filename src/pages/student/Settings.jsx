import React, { useState, useEffect } from 'react';
import './Settings.css';

function Settings() {
  // Состояния для каждого раздела (открыт/закрыт)
  const [openSections, setOpenSections] = useState({
    password: false,
    email: false,
    phone: false,
    theme: false,
    language: false,
    accessibility: false,
    version: false,
    notifications: false,
  });

  // Состояния для значений настроек
  const [settings, setSettings] = useState({
    password: '',
    newPassword: '',
    confirmPassword: '',
    email: '',
    newEmail: '',
    phone: '',
    newPhone: '',
    theme: 'light',
    language: 'ru',
    accessibilityMode: false,
    notificationsEnabled: true,
  });

  // Загрузка сохраненных настроек
  useEffect(() => {
    const savedSettings = localStorage.getItem('appSettings');
    if (savedSettings) {
      setSettings(JSON.parse(savedSettings));
    }
  }, []);

  // Сохранение настроек
  const saveSettings = (newSettings) => {
    setSettings(newSettings);
    localStorage.setItem('appSettings', JSON.stringify(newSettings));
  };

  // Переключение открытия/закрытия секции
  const toggleSection = (section) => {
    setOpenSections(prev => ({
      ...prev,
      [section]: !prev[section]
    }));
  };

  // Смена пароля
  const handleChangePassword = () => {
    if (!settings.password) {
      alert('Введите текущий пароль');
      return;
    }
    if (settings.newPassword !== settings.confirmPassword) {
      alert('Новый пароль и подтверждение не совпадают');
      return;
    }
    if (settings.newPassword.length < 6) {
      alert('Пароль должен содержать минимум 6 символов');
      return;
    }
    alert('Пароль успешно изменен!');
    setSettings({
      ...settings,
      password: '',
      newPassword: '',
      confirmPassword: '',
    });
    toggleSection('password');
  };

  // Смена email
  const handleChangeEmail = () => {
    if (!settings.newEmail) {
      alert('Введите новый email');
      return;
    }
    if (!settings.newEmail.includes('@')) {
      alert('Введите корректный email');
      return;
    }
    alert(`Email успешно изменен на ${settings.newEmail}`);
    setSettings({
      ...settings,
      email: settings.newEmail,
      newEmail: '',
    });
    toggleSection('email');
  };

  // Смена телефона
  const handleChangePhone = () => {
    if (!settings.newPhone) {
      alert('Введите новый номер телефона');
      return;
    }
    if (settings.newPhone.length < 11) {
      alert('Введите корректный номер телефона');
      return;
    }
    alert(`Номер телефона успешно изменен на ${settings.newPhone}`);
    setSettings({
      ...settings,
      phone: settings.newPhone,
      newPhone: '',
    });
    toggleSection('phone');
  };

  // Смена темы
  const handleChangeTheme = (theme) => {
    setSettings({ ...settings, theme });
    if (theme === 'dark') {
      document.body.classList.add('dark-theme');
    } else {
      document.body.classList.remove('dark-theme');
    }
    toggleSection('theme');
  };

  // Смена языка
  const handleChangeLanguage = (lang) => {
    const langNames = {
      ru: 'Русский',
      en: 'English',
      kz: 'Қазақша',
    };
    alert(`Язык изменен на ${langNames[lang]}`);
    setSettings({ ...settings, language: lang });
    toggleSection('language');
  };

  // Версия для слабовидящих
  const handleAccessibilityToggle = () => {
    const newMode = !settings.accessibilityMode;
    setSettings({ ...settings, accessibilityMode: newMode });
    
    if (newMode) {
      document.body.classList.add('accessibility-mode');
      alert('Включен режим для слабовидящих: увеличенный шрифт и высокий контраст');
    } else {
      document.body.classList.remove('accessibility-mode');
      alert('Режим для слабовидящих выключен');
    }
    toggleSection('accessibility');
  };

  // Уведомления
  const handleNotificationsToggle = () => {
    const newState = !settings.notificationsEnabled;
    setSettings({ ...settings, notificationsEnabled: newState });
    alert(`Уведомления ${newState ? 'включены' : 'выключены'}`);
    toggleSection('notifications');
  };

  // Версия приложения
  const appVersion = "2.0.1";

  return (
    <div className="settings-container">
      <h2 className="settings-title">⚙️ Настройки</h2>
      <p className="settings-subtitle">Управление параметрами аккаунта и приложения</p>

      <div className="settings-list">
        {/* === СМЕНА ПАРОЛЯ === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('password')}>
            <div className="card-icon">🔐</div>
            <div className="card-info">
              <h3>Пароль</h3>
              <p>Изменение пароля для входа в аккаунт</p>
            </div>
            <div className={`card-arrow ${openSections.password ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.password && (
            <div className="settings-card-content">
              <div className="input-group">
                <label>Текущий пароль</label>
                <input
                  type="password"
                  placeholder="Введите текущий пароль"
                  value={settings.password}
                  onChange={(e) => setSettings({ ...settings, password: e.target.value })}
                />
              </div>
              <div className="input-group">
                <label>Новый пароль</label>
                <input
                  type="password"
                  placeholder="Введите новый пароль"
                  value={settings.newPassword}
                  onChange={(e) => setSettings({ ...settings, newPassword: e.target.value })}
                />
              </div>
              <div className="input-group">
                <label>Подтверждение пароля</label>
                <input
                  type="password"
                  placeholder="Подтвердите новый пароль"
                  value={settings.confirmPassword}
                  onChange={(e) => setSettings({ ...settings, confirmPassword: e.target.value })}
                />
              </div>
              <button className="save-btn" onClick={handleChangePassword}>Сохранить</button>
            </div>
          )}
        </div>

        {/* === СМЕНА EMAIL === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('email')}>
            <div className="card-icon">📧</div>
            <div className="card-info">
              <h3>Электронная почта</h3>
              <p>Текущий email: {settings.email || 'не указан'}</p>
            </div>
            <div className={`card-arrow ${openSections.email ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.email && (
            <div className="settings-card-content">
              <div className="input-group">
                <label>Новый email</label>
                <input
                  type="email"
                  placeholder="Введите новый email"
                  value={settings.newEmail}
                  onChange={(e) => setSettings({ ...settings, newEmail: e.target.value })}
                />
              </div>
              <button className="save-btn" onClick={handleChangeEmail}>Привязать</button>
            </div>
          )}
        </div>

        {/* === СМЕНА ТЕЛЕФОНА === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('phone')}>
            <div className="card-icon">📱</div>
            <div className="card-info">
              <h3>Номер телефона</h3>
              <p>Текущий номер: {settings.phone || 'не указан'}</p>
            </div>
            <div className={`card-arrow ${openSections.phone ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.phone && (
            <div className="settings-card-content">
              <div className="input-group">
                <label>Новый номер телефона</label>
                <input
                  type="tel"
                  placeholder="+7 (XXX) XXX-XX-XX"
                  value={settings.newPhone}
                  onChange={(e) => setSettings({ ...settings, newPhone: e.target.value })}
                />
              </div>
              <button className="save-btn" onClick={handleChangePhone}>Привязать</button>
            </div>
          )}
        </div>

        {/* === ТЕМА === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('theme')}>
            <div className="card-icon">🎨</div>
            <div className="card-info">
              <h3>Тема оформления</h3>
              <p>Текущая тема: {settings.theme === 'light' ? 'Светлая' : 'Темная'}</p>
            </div>
            <div className={`card-arrow ${openSections.theme ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.theme && (
            <div className="settings-card-content">
              <div className="theme-options">
                <div 
                  className={`theme-option ${settings.theme === 'light' ? 'active' : ''}`}
                  onClick={() => handleChangeTheme('light')}
                >
                  <div className="theme-preview light-preview"></div>
                  <span>☀️ Светлая</span>
                </div>
                <div 
                  className={`theme-option ${settings.theme === 'dark' ? 'active' : ''}`}
                  onClick={() => handleChangeTheme('dark')}
                >
                  <div className="theme-preview dark-preview"></div>
                  <span>🌙 Темная</span>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* === ЯЗЫК === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('language')}>
            <div className="card-icon">🌐</div>
            <div className="card-info">
              <h3>Язык интерфейса</h3>
              <p>
                Текущий язык: {
                  settings.language === 'ru' ? 'Русский' : 
                  settings.language === 'en' ? 'English' : 'Қазақша'
                }
              </p>
            </div>
            <div className={`card-arrow ${openSections.language ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.language && (
            <div className="settings-card-content">
              <div className="language-options">
                <div 
                  className={`language-option ${settings.language === 'ru' ? 'active' : ''}`}
                  onClick={() => handleChangeLanguage('ru')}
                >
                  <span>🇷🇺</span>
                  <span>Русский</span>
                </div>
                <div 
                  className={`language-option ${settings.language === 'en' ? 'active' : ''}`}
                  onClick={() => handleChangeLanguage('en')}
                >
                  <span>🇬🇧</span>
                  <span>English</span>
                </div>
                <div 
                  className={`language-option ${settings.language === 'kz' ? 'active' : ''}`}
                  onClick={() => handleChangeLanguage('kz')}
                >
                  <span>🇰🇿</span>
                  <span>Қазақша</span>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* === ВЕРСИЯ ДЛЯ СЛАБОВИДЯЩИХ === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('accessibility')}>
            <div className="card-icon">👁️</div>
            <div className="card-info">
              <h3>Версия для слабовидящих</h3>
              <p>Режим с увеличенным шрифтом и высоким контрастом</p>
            </div>
            <div className="toggle-switch">
              <span className={`toggle-slider ${settings.accessibilityMode ? 'active' : ''}`}></span>
            </div>
          </div>
          {openSections.accessibility && (
            <div className="settings-card-content">
              <p className="accessibility-description">
                Включение этого режима увеличит размер шрифта, улучшит контрастность 
                и сделает интерфейс более удобным для людей с нарушениями зрения.
              </p>
              <button 
                className={`accessibility-btn ${settings.accessibilityMode ? 'active' : ''}`}
                onClick={handleAccessibilityToggle}
              >
                {settings.accessibilityMode ? '🔘 Выключить режим' : '🔘 Включить режим'}
              </button>
            </div>
          )}
        </div>

        {/* === УВЕДОМЛЕНИЯ === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('notifications')}>
            <div className="card-icon">🔔</div>
            <div className="card-info">
              <h3>Уведомления</h3>
              <p>Включить/выключить push-уведомления</p>
            </div>
            <div className="toggle-switch">
              <span className={`toggle-slider ${settings.notificationsEnabled ? 'active' : ''}`}></span>
            </div>
          </div>
          {openSections.notifications && (
            <div className="settings-card-content">
              <p className="notifications-description">
                {settings.notificationsEnabled 
                  ? 'Уведомления включены. Вы будете получать оповещения о новых сообщениях, обновлениях и событиях.'
                  : 'Уведомления выключены. Вы не будете получать оповещения.'}
              </p>
              <button 
                className={`notifications-btn ${settings.notificationsEnabled ? 'active' : ''}`}
                onClick={handleNotificationsToggle}
              >
                {settings.notificationsEnabled ? '🔔 Выключить уведомления' : '🔕 Включить уведомления'}
              </button>
            </div>
          )}
        </div>

        {/* === ВЕРСИЯ ПРИЛОЖЕНИЯ === */}
        <div className="settings-card">
          <div className="settings-card-header" onClick={() => toggleSection('version')}>
            <div className="card-icon">ℹ️</div>
            <div className="card-info">
              <h3>Версия приложения</h3>
              <p>Информация о текущей версии</p>
            </div>
            <div className={`card-arrow ${openSections.version ? 'open' : ''}`}>▼</div>
          </div>
          {openSections.version && (
            <div className="settings-card-content">
              <div className="version-info">
                <div className="version-number">
                  <span className="version-label">Текущая версия:</span>
                  <span className="version-value">{appVersion}</span>
                </div>
                <div className="version-date">
                  <span className="version-label">Дата сборки:</span>
                  <span className="version-value">15.05.2026</span>
                </div>
                <div className="version-changelog">
                  <h4>Что нового в версии {appVersion}:</h4>
                  <ul>
                    <li>📚 Добавлена электронная библиотека</li>
                    <li>🏆 Система достижений ПГАС</li>
                    <li>🌙 Темная тема оформления</li>
                    <li>🔍 Расширенный поиск в библиотеке</li>
                    <li>🐛 Исправлены ошибки</li>
                  </ul>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Settings;