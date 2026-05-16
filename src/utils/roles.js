// Типы ролей
export const ROLES = {
  ADMIN: 'admin',
  TEACHER: 'teacher',
  STAFF: 'staff',
  STUDENT: 'student'
};

// МЕНЮ ДЛЯ АДМИНА
export const MENU_ITEMS = {
  [ROLES.ADMIN]: [
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
    { path: '/admin/shop', icon: '🛒', name: 'Магазин' },
    { path: '/admin/settings', icon: '⚙️', name: 'Настройки' }
  ],
  [ROLES.STUDENT]: [
    { path: '/profile', icon: '👤', name: 'Профиль' },
    { path: '/schedule', icon: '📅', name: 'Расписание' },
    { path: '/library', icon: '📚', name: 'Библиотека' },
    { path: '/news', icon: '📰', name: 'Новости' },
    { path: '/entertainment', icon: '🎮', name: 'ДПО и развлечения' },
    { path: '/achievements', icon: '🏆', name: 'Ачивки' },
    { path: '/skilltree', icon: '🌳', name: 'Дерево навыков' },
    { path: '/deadlines', icon: '⏰', name: 'Дедлайны' },
    { path: '/checklist', icon: '📋', name: 'Чек-лист' },
    { path: '/exam', icon: '🎓', name: 'Экзамен' },
    { path: '/challenges', icon: '🎯', name: 'Вызовы' },
    { path: '/minigames', icon: '🎮', name: 'Мини-игры' },
    { path: '/planner', icon: '📅', name: 'Планер' },
    { path: '/shop', icon: '🛒', name: 'Магазин' },
    { path: '/settings', icon: '⚙️', name: 'Настройки' }
  ],
  [ROLES.TEACHER]: [
    { path: '/teacher/dashboard', icon: '📊', name: 'Панель' },
    { path: '/teacher/profile', icon: '👤', name: 'Профиль' },
    { path: '/teacher/schedule', icon: '📅', name: 'Расписание' },
    { path: '/teacher/students', icon: '👨‍🎓', name: 'Студенты' },
    { path: '/teacher/achievements', icon: '🏆', name: 'Ачивки' },
    { path: '/teacher/deadlines', icon: '⏰', name: 'Дедлайны' },
    { path: '/teacher/shop', icon: '🛒', name: 'Магазин' }
  ]
};

// Демо-пользователи (только для отображения подсказок, реальная проверка через API)
export const DEMO_USERS = [
  { login: "admin", password: "admin123", role: ROLES.ADMIN, name: "Администратор" },
  { login: "teacher", password: "teacher123", role: ROLES.TEACHER, name: "Преподаватель" },
  { login: "student", password: "student123", role: ROLES.STUDENT, name: "Алексей Студент" }
];

// API базовый URL
const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// Получение токена из localStorage
export const getToken = () => {
  return localStorage.getItem('token');
};

// Проверка авторизации
export const isAuthenticated = () => {
  return !!localStorage.getItem('token');
};

// Получение текущего пользователя из localStorage
export const getCurrentUser = () => {
  const userStr = localStorage.getItem('user');
  if (userStr) {
    try {
      return JSON.parse(userStr);
    } catch (e) {
      return null;
    }
  }
  return null;
};

// Получение роли текущего пользователя
export const getUserRole = () => {
  const user = getCurrentUser();
  return user?.role || null;
};

// Получение логина текущего пользователя
export const getUserLogin = () => {
  const user = getCurrentUser();
  return user?.login || null;
};

// Сохранение роли (для обратной совместимости)
export const saveUserRole = (login, role) => {
  localStorage.setItem(`userRole_${login}`, role);
};

// Выход из системы
export const logout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  // Очищаем старые данные
  const keys = Object.keys(localStorage);
  keys.forEach(key => {
    if (key.startsWith('userRole_')) {
      localStorage.removeItem(key);
    }
  });
};

// Проверка роли пользователя
export const hasRole = (role) => {
  const userRole = getUserRole();
  return userRole === role;
};

// Проверка, является ли пользователь администратором
export const isAdmin = () => hasRole(ROLES.ADMIN);

// Проверка, является ли пользователь преподавателем
export const isTeacher = () => hasRole(ROLES.TEACHER);

// Проверка, является ли пользователь студентом
export const isStudent = () => hasRole(ROLES.STUDENT);