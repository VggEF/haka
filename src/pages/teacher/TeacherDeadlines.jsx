import React, { useState, useEffect } from 'react';
import './Teacher.css';

function TeacherDeadlines() {
  const [deadlines, setDeadlines] = useState([]);
  const [formData, setFormData] = useState({ title: "", group: "", subject: "", dueDate: "", description: "" });

  useEffect(() => {
    const saved = localStorage.getItem('teacherDeadlines');
    if (saved) setDeadlines(JSON.parse(saved));
  }, []);

  const addDeadline = () => {
    if (!formData.title || !formData.dueDate) return alert("Заполните название и дату");
    const newDeadline = { id: Date.now(), ...formData, status: "active" };
    setDeadlines([...deadlines, newDeadline]);
    localStorage.setItem('teacherDeadlines', JSON.stringify([...deadlines, newDeadline]));
    setFormData({ title: "", group: "", subject: "", dueDate: "", description: "" });
    alert("Дедлайн создан!");
  };

  const deleteDeadline = (id) => {
    if (window.confirm("Удалить дедлайн?")) {
      setDeadlines(deadlines.filter(d => d.id !== id));
      localStorage.setItem('teacherDeadlines', JSON.stringify(deadlines.filter(d => d.id !== id)));
    }
  };

  return (
    <div className="teacher-page">
      <div className="teacher-header">
        <div className="teacher-title">
          <div className="teacher-title-icon">⏰</div>
          <h1>Дедлайны для студентов</h1>
        </div>
        <div className="teacher-subtitle">Создавайте дедлайны по домашним заданиям</div>
      </div>

      <div className="teacher-card">
        <div className="card-title">➕ Новый дедлайн</div>
        <div className="form-group"><input type="text" placeholder="Название задания" value={formData.title} onChange={(e) => setFormData({...formData, title: e.target.value})} /></div>
        <div className="form-group"><input type="text" placeholder="Группа (23-ПМбо-1)" value={formData.group} onChange={(e) => setFormData({...formData, group: e.target.value})} /></div>
        <div className="form-group"><input type="text" placeholder="Предмет" value={formData.subject} onChange={(e) => setFormData({...formData, subject: e.target.value})} /></div>
        <div className="form-group"><input type="date" value={formData.dueDate} onChange={(e) => setFormData({...formData, dueDate: e.target.value})} /></div>
        <div className="form-group"><textarea placeholder="Описание" rows="2" value={formData.description} onChange={(e) => setFormData({...formData, description: e.target.value})} /></div>
        <button className="btn-primary" onClick={addDeadline}>➕ Создать дедлайн</button>
      </div>

      <div className="items-list">
        {deadlines.map(d => (
          <div key={d.id} className="item-card">
            <div><strong>📝 {d.title}</strong> | {d.subject}</div>
            <div>👥 {d.group || "Все группы"} | ⏰ до {d.dueDate}</div>
            <div>{d.description}</div>
            <button className="btn-danger" onClick={() => deleteDeadline(d.id)}>🗑️ Удалить</button>
          </div>
        ))}
        {deadlines.length === 0 && <div className="empty-state">Нет дедлайнов</div>}
      </div>
    </div>
  );
}

export default TeacherDeadlines;