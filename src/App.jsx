import React, { useState } from 'react';
import { Routes, Route, useNavigate, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import { ROLES } from './utils/roles';
import './App.css';

// АДМИН
import AdminDashboard from './pages/admin/AdminDashboard';
import AdminStudents from './pages/admin/AdminStudents';
import AdminSchedule from './pages/admin/AdminSchedule';
import AdminLibrary from './pages/admin/AdminLibrary';
import AdminNews from './pages/admin/AdminNews';
import AdminEvents from './pages/admin/AdminEvents';
import AdminAchievements from './pages/admin/AdminAchievements';
import AdminExam from './pages/admin/AdminExam';
import AdminChallenges from './pages/admin/AdminChallenges';
import AdminGames from './pages/admin/AdminGames';
import AdminSettings from './pages/admin/AdminSettings';
import SidebarAdmin from './components/SidebarAdmin';
import AdminShop from './pages/admin/AdminShop';


// ПРЕПОДАВАТЕЛЬ
import TeacherDashboard from './pages/teacher/TeacherDashboard';
import TeacherProfile from './pages/teacher/TeacherProfile';
import TeacherSchedule from './pages/teacher/TeacherSchedule';
import TeacherStudents from './pages/teacher/TeacherStudents';
import TeacherAchievements from './pages/teacher/TeacherAchievements';
import TeacherDeadlines from './pages/teacher/TeacherDeadlines';
import SidebarTeacher from './components/SidebarTeacher';

// СТУДЕНТ
import Profile from './pages/student/Profile';
import Schedule from './pages/student/Schedule';
import Library from './pages/student/Library';
import SidebarStudent from './components/SidebarStudent';
import Achievements from './pages/student/Achievements';
import Challenges from './pages/student/Challenges';
import Chats from './pages/student/Chats';
import Checklist from './pages/student/Checklist';
import Deadlines from './pages/student/Deadlines';
import Entertainment from './pages/student/Entertainment';
import ExamGenerator from './pages/student/ExamGenerator';  // Экзамены
import MiniGames from './pages/student/MiniGames';          // Мини-игры
import News from './pages/student/News';
import Planner from './pages/student/Planner';
import Settings from './pages/student/Settings';
import Shop from './pages/student/Shop';
import SkillTree from './pages/student/SkillTree';          // Дерево навыков

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userLogin, setUserLogin] = useState('');
  const [userRole, setUserRole] = useState('');
  const navigate = useNavigate();

  const handleLogin = (login, role, name) => {
    setUserLogin(login);
    setUserRole(role);
    setIsLoggedIn(true);
    
    // Редирект в зависимости от роли
    if (role === ROLES.ADMIN) {
      navigate('/admin/dashboard');
    } else if (role === ROLES.TEACHER) {
      navigate('/teacher/dashboard');
    } else {
      navigate('/profile');
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUserLogin('');
    setUserRole('');
    navigate('/');
  };

  if (!isLoggedIn) {
    return <Login onLogin={handleLogin} />;
  }

  // АДМИН
  if (userRole === ROLES.ADMIN) {
    return (
      <div className="app-container">
        <SidebarAdmin onLogout={handleLogout} />
        <div className="app-content">
          <Routes>
            <Route path="/admin/dashboard" element={<AdminDashboard />} />
            <Route path="/admin/students" element={<AdminStudents />} />
            <Route path="/admin/schedule" element={<AdminSchedule />} />
            <Route path="/admin/library" element={<AdminLibrary />} />
            <Route path="/admin/news" element={<AdminNews />} />
            <Route path="/admin/events" element={<AdminEvents />} />
            <Route path="/admin/achievements" element={<AdminAchievements />} />
            <Route path="/admin/exam" element={<AdminExam />} />
            <Route path="/admin/challenges" element={<AdminChallenges />} />
            <Route path="/admin/games" element={<AdminGames />} />
            <Route path="/admin/settings" element={<AdminSettings />} />
            <Route path="*" element={<Navigate to="/admin/dashboard" replace />} />
            <Route path="/admin/shop" element={<AdminShop />} />
          </Routes>
        </div>
      </div>
    );
  }

  // ПРЕПОДАВАТЕЛЬ
  if (userRole === ROLES.TEACHER) {
    return (
      <div className="app-container">
        <SidebarTeacher onLogout={handleLogout} />
        <div className="app-content">
          <Routes>
            <Route path="/teacher/dashboard" element={<TeacherDashboard />} />
            <Route path="/teacher/profile" element={<TeacherProfile />} />
            <Route path="/teacher/schedule" element={<TeacherSchedule />} />
            <Route path="/teacher/students" element={<TeacherStudents />} />
            <Route path="/teacher/achievements" element={<TeacherAchievements />} />
            <Route path="/teacher/deadlines" element={<TeacherDeadlines />} />
            <Route path="*" element={<Navigate to="/teacher/dashboard" replace />} />
          </Routes>
        </div>
      </div>
    );
  }

  // СТУДЕНТ - ВСЕ МАРШРУТЫ
  return (
    <div className="app-container">
      <SidebarStudent onLogout={handleLogout} userLogin={userLogin} />
      <div className="app-content">
        <Routes>
          {/* Основные */}
          <Route path="/profile" element={<Profile login={userLogin} />} />
          <Route path="/schedule" element={<Schedule userLogin={userLogin} />} />
          <Route path="/library" element={<Library />} />
          <Route path="/achievements" element={<Achievements />} />
          <Route path="/challenges" element={<Challenges />} />
          <Route path="/chats" element={<Chats />} />
          <Route path="/checklist" element={<Checklist />} />
          <Route path="/deadlines" element={<Deadlines />} />
          <Route path="/shop" element={<Shop />} />
          
          {/* Развлечения и игры */}
          <Route path="/entertainment" element={<Entertainment />} />
          <Route path="/mini-games" element={<MiniGames />} />        {/* Мини-игры */}
          <Route path="/skill-tree" element={<SkillTree />} />        {/* Дерево навыков */}
          
          {/* Обучение и тесты */}
          <Route path="/exam-generator" element={<ExamGenerator />} /> {/* Экзамены/генератор экзаменов */}
          
          {/* Информация и настройки */}
          <Route path="/news" element={<News />} />
          <Route path="/planner" element={<Planner />} />
          <Route path="/settings" element={<Settings />} />
          
          {/* Редирект */}
          <Route path="*" element={<Navigate to="/profile" replace />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;