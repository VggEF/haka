import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Teacher.css';

function TeacherDashboard() {
  const navigate = useNavigate();
  const [stats, setStats] = useState({
    students: 0,
    achievements: 0,
    deadlines: 0,
    schedule: 0
  });

  useEffect(() => {
    // Загружаем данные
    const students = JSON.parse(localStorage.getItem('students') || '[]');
    const achievements = JSON.parse(localStorage.getItem('teacherAchievements') || '[]');
    const deadlines = JSON.parse(localStorage.getItem('teacherDeadlines') || '[]');
    const schedule = JSON.parse(localStorage.getItem('teacherSchedule') || '[]');
    
    setStats({
      students: students.length,
      achievements: achievements.length,
      deadlines: deadlines.length,
      schedule: schedule.length
    });
  }, []);

  const quickLinks = [
    { title: "Студенты", icon: "👨‍🎓", count: stats.students, path: "/teacher/students", color: "#667eea" },
    { title: "Расписание", icon: "📅", count: stats.schedule, path: "/teacher/schedule", color: "#4facfe" },
    { title: "Ачивки", icon: "🏆", count: stats.achievements, path: "/teacher/achievements", color: "#f5576c" },
    { title: "Дедлайны", icon: "⏰", count: stats.deadlines, path: "/teacher/deadlines", color: "#ff9800" }
  ];

  return (
    <div className="teacher-page">
      <div className="teacher-header">
        <div className="teacher-title">
          <div className="teacher-title-icon">📊</div>
          <h1>Панель преподавателя</h1>
        </div>
        <div className="teacher-subtitle">Управляйте расписанием, ачивками и дедлайнами</div>
      </div>

      <div className="stats-grid">
        {quickLinks.map((link, index) => (
          <div key={index} className="stat-card" onClick={() => navigate(link.path)}>
            <div className="stat-header">
              <div className="stat-icon">{link.icon}</div>
              <div className="stat-value">{link.count}</div>
            </div>
            <div className="stat-label">{link.title}</div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default TeacherDashboard;