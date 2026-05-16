import React, { useState, useEffect } from 'react';
import './Teacher.css';

function TeacherStudents() {
  const [students, setStudents] = useState([]);

  useEffect(() => {
    const saved = JSON.parse(localStorage.getItem('students') || '[]');
    setStudents(saved);
  }, []);

  return (
    <div className="teacher-page">
      <div className="teacher-header">
        <div className="teacher-title">
          <div className="teacher-title-icon">👨‍🎓</div>
          <h1>Мои студенты</h1>
        </div>
        <div className="teacher-subtitle">Список студентов, обучающихся у вас</div>
      </div>

      <div className="teacher-card">
        {students.length === 0 ? (
          <div className="empty-state">Нет студентов</div>
        ) : (
          students.map(s => (
            <div key={s.login} className="item-card">
              <div><strong>{s.name}</strong> | {s.login}</div>
              <div>📚 Группа: {s.group || "Не указана"}</div>
              <div>🎓 Статус: Студент</div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default TeacherStudents;