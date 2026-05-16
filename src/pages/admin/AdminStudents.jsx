import React, { useState } from 'react';
import { DEMO_USERS, ROLES } from '../../utils/roles';
import './AdminStudents.css';

function AdminStudents() {
  const [students, setStudents] = useState(DEMO_USERS.filter(u => u.role === ROLES.STUDENT));
  const [newStudent, setNewStudent] = useState({ login: "", name: "", group: "", password: "" });

  const addStudent = () => {
    if (!newStudent.login || !newStudent.name) {
      alert("Заполните логин и ФИО");
      return;
    }
    const newUser = {
      login: newStudent.login,
      password: newStudent.password || "student123",
      role: ROLES.STUDENT,
      name: newStudent.name,
      group: newStudent.group || "Не указана"
    };
    setStudents([...students, newUser]);
    setNewStudent({ login: "", name: "", group: "", password: "" });
    alert(`Студент ${newStudent.name} добавлен!`);
  };

  const deleteStudent = (login) => {
    if (window.confirm(`Удалить студента ${login}?`)) {
      setStudents(students.filter(s => s.login !== login));
    }
  };

  // Получение инициалов для аватара
  const getInitials = (name) => {
    const parts = name.split(' ');
    if (parts.length >= 2) {
      return `${parts[0][0]}${parts[1][0]}`;
    }
    return name[0] || '?';
  };

  return (
    <div className="admin-students">
      <div className="students-header">
        <div className="students-title">
          <div className="students-title-icon">👨‍🎓</div>
          <h1>Управление студентами</h1>
        </div>
        <div className="students-subtitle">Добавляйте и управляйте студентами вашего университета</div>
      </div>

      {/* Карточка добавления */}
      <div className="add-student-card">
        <div className="card-title">
          <span className="card-title-icon">➕</span>
          <span>Добавить нового студента</span>
        </div>
        <div className="add-student-form">
          <div className="form-field">
            <label>Логин</label>
            <input 
              type="text" 
              placeholder="например: 23-ПМбо-999" 
              value={newStudent.login} 
              onChange={(e) => setNewStudent({...newStudent, login: e.target.value})} 
            />
          </div>
          <div className="form-field">
            <label>ФИО</label>
            <input 
              type="text" 
              placeholder="Иванов Иван Иванович" 
              value={newStudent.name} 
              onChange={(e) => setNewStudent({...newStudent, name: e.target.value})} 
            />
          </div>
          <div className="form-field">
            <label>Группа</label>
            <input 
              type="text" 
              placeholder="23-ПМбо-1" 
              value={newStudent.group} 
              onChange={(e) => setNewStudent({...newStudent, group: e.target.value})} 
            />
          </div>
          <div className="form-field">
            <label>Пароль</label>
            <input 
              type="text" 
              placeholder="По умолчанию: student123" 
              value={newStudent.password} 
              onChange={(e) => setNewStudent({...newStudent, password: e.target.value})} 
            />
          </div>
          <button className="submit-btn" onClick={addStudent}>
            ➕ Добавить студента
          </button>
        </div>
      </div>

      {/* Таблица студентов */}
      <div className="students-table-container">
        <div className="table-header">
          <div className="table-title">
            <span>📋</span>
            <span>Список студентов</span>
          </div>
          <div className="table-stats">Всего: {students.length}</div>
        </div>
        
        {students.length === 0 ? (
          <div className="empty-state">
            <div className="empty-icon">👨‍🎓</div>
            <div className="empty-text">Нет студентов. Добавьте первого!</div>
          </div>
        ) : (
          <table className="students-table">
            <thead>
              <tr>
                <th>Студент</th>
                <th>Логин</th>
                <th>Группа</th>
                <th>Роль</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {students.map(student => (
                <tr key={student.login}>
                  <td>
                    <div className="student-cell">
                      <div className="student-avatar">{getInitials(student.name)}</div>
                      <div className="student-info">
                        <span className="student-name">{student.name}</span>
                        <span className="student-login">ID: {student.login}</span>
                      </div>
                    </div>
                  </td>
                  <td>{student.login}</td>
                  <td><span className="group-badge">{student.group || "—"}</span></td>
                  <td><span className="role-badge">🎓 Студент</span></td>
                  <td>
                    <button className="delete-btn" onClick={() => deleteStudent(student.login)} title="Удалить">
                      🗑️
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}

export default AdminStudents;