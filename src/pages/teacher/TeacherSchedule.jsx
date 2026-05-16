import React, { useState, useEffect } from 'react';
import './Teacher.css';

function TeacherSchedule() {
  const [schedule, setSchedule] = useState([]);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({ group: "", discipline: "", time: "", audience: "", day: "" });

  useEffect(() => {
    const saved = localStorage.getItem('teacherSchedule');
    if (saved) setSchedule(JSON.parse(saved));
  }, []);

  const saveSchedule = (newSchedule) => {
    setSchedule(newSchedule);
    localStorage.setItem('teacherSchedule', JSON.stringify(newSchedule));
  };

  const addSchedule = () => {
    if (!formData.discipline || !formData.group) {
      alert("Заполните дисциплину и группу");
      return;
    }
    saveSchedule([...schedule, { id: Date.now(), ...formData }]);
    setFormData({ group: "", discipline: "", time: "", audience: "", day: "" });
    setShowForm(false);
  };

  const deleteSchedule = (id) => {
    if (window.confirm("Удалить занятие?")) {
      saveSchedule(schedule.filter(item => item.id !== id));
    }
  };

  const days = ["Понедельник", "Вторник", "Среда", "Четверг", "Пятница"];

  return (
    <div className="teacher-page">
      <div className="teacher-header">
        <div className="teacher-title">
          <div className="teacher-title-icon">📅</div>
          <h1>Управление расписанием</h1>
        </div>
        <div className="teacher-subtitle">Добавляйте и редактируйте занятия</div>
      </div>

      <button className="btn-primary" onClick={() => setShowForm(true)}>➕ Добавить занятие</button>

      {showForm && (
        <div className="teacher-card" style={{ marginTop: 20 }}>
          <div className="card-title">Новое занятие</div>
          <div className="form-group"><label>Группа</label><input type="text" placeholder="23-ПМбо-1" value={formData.group} onChange={(e) => setFormData({...formData, group: e.target.value})} /></div>
          <div className="form-group"><label>Дисциплина</label><input type="text" placeholder="Математика" value={formData.discipline} onChange={(e) => setFormData({...formData, discipline: e.target.value})} /></div>
          <div className="form-group"><label>Время</label><input type="text" placeholder="11:50" value={formData.time} onChange={(e) => setFormData({...formData, time: e.target.value})} /></div>
          <div className="form-group"><label>Аудитория</label><input type="text" placeholder="Б-201" value={formData.audience} onChange={(e) => setFormData({...formData, audience: e.target.value})} /></div>
          <div className="form-group"><label>День недели</label><select value={formData.day} onChange={(e) => setFormData({...formData, day: e.target.value})}><option value="">Выберите</option>{days.map(d => <option key={d}>{d}</option>)}</select></div>
          <button className="btn-primary" onClick={addSchedule}>Сохранить</button>
          <button className="btn-secondary" style={{ marginLeft: 12 }} onClick={() => setShowForm(false)}>Отмена</button>
        </div>
      )}

      <div className="items-list" style={{ marginTop: 20 }}>
        {schedule.length === 0 && <div className="empty-state">Нет занятий</div>}
        {schedule.map(item => (
          <div key={item.id} className="item-card">
            <div><strong>{item.discipline}</strong> | {item.group}</div>
            <div>⏰ {item.time} | 📍 {item.audience} | 📅 {item.day}</div>
            <button className="btn-danger" onClick={() => deleteSchedule(item.id)}>🗑️ Удалить</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default TeacherSchedule;