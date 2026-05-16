import React, { useState, useEffect } from 'react';
import './Teacher.css';

function TeacherAchievements() {
  const [achievements, setAchievements] = useState([]);
  const [students, setStudents] = useState([]);
  const [selectedStudent, setSelectedStudent] = useState("");
  const [selectedGroup, setSelectedGroup] = useState("");
  const [selectedAchievement, setSelectedAchievement] = useState("");
  const [newAchievement, setNewAchievement] = useState({ title: "", description: "", xp: 50 });

  useEffect(() => {
    const saved = localStorage.getItem('teacherAchievements');
    if (saved) setAchievements(JSON.parse(saved));
    const savedStudents = JSON.parse(localStorage.getItem('students') || '[]');
    setStudents(savedStudents);
  }, []);

  const addAchievement = () => {
    if (!newAchievement.title) return alert("Введите название");
    const achievement = { id: Date.now(), ...newAchievement };
    setAchievements([...achievements, achievement]);
    localStorage.setItem('teacherAchievements', JSON.stringify([...achievements, achievement]));
    setNewAchievement({ title: "", description: "", xp: 50 });
  };

  const awardToStudent = () => {
    if (!selectedStudent || !selectedAchievement) return alert("Выберите студента и ачивку");
    const achievement = achievements.find(a => a.id === parseInt(selectedAchievement));
    if (!achievement) return;
    const studentAchievements = JSON.parse(localStorage.getItem(`achievements_${selectedStudent}`) || '[]');
    if (studentAchievements.includes(achievement.id)) return alert("У студента уже есть эта ачивка");
    studentAchievements.push(achievement.id);
    localStorage.setItem(`achievements_${selectedStudent}`, JSON.stringify(studentAchievements));
    const currentXP = parseInt(localStorage.getItem(`score_${selectedStudent}`) || '0');
    localStorage.setItem(`score_${selectedStudent}`, currentXP + achievement.xp);
    alert(`Ачивка выдана! +${achievement.xp} XP`);
  };

  const groups = [...new Set(students.map(s => s.group || "Без группы"))];

  return (
    <div className="teacher-page">
      <div className="teacher-header">
        <div className="teacher-title">
          <div className="teacher-title-icon">🏆</div>
          <h1>Ачивки для студентов</h1>
        </div>
        <div className="teacher-subtitle">Создавайте ачивки и выдавайте их студентам</div>
      </div>

      <div className="teacher-card">
        <div className="card-title">➕ Создать ачивку</div>
        <div className="form-group"><input type="text" placeholder="Название" value={newAchievement.title} onChange={(e) => setNewAchievement({...newAchievement, title: e.target.value})} /></div>
        <div className="form-group"><textarea placeholder="Описание" rows="2" value={newAchievement.description} onChange={(e) => setNewAchievement({...newAchievement, description: e.target.value})} /></div>
        <div className="form-group"><input type="number" placeholder="XP" value={newAchievement.xp} onChange={(e) => setNewAchievement({...newAchievement, xp: parseInt(e.target.value)})} /></div>
        <button className="btn-primary" onClick={addAchievement}>➕ Создать</button>
      </div>

      <div className="teacher-card">
        <div className="card-title">🎯 Выдать ачивку</div>
        <div className="form-group"><select value={selectedStudent} onChange={(e) => setSelectedStudent(e.target.value)}><option value="">Выберите студента</option>{students.map(s => <option key={s.login} value={s.login}>{s.name}</option>)}</select></div>
        <div className="form-group"><select value={selectedAchievement} onChange={(e) => setSelectedAchievement(e.target.value)}><option value="">Выберите ачивку</option>{achievements.map(a => <option key={a.id} value={a.id}>{a.title} (+{a.xp})</option>)}</select></div>
        <button className="btn-primary" onClick={awardToStudent}>🎁 Выдать</button>
      </div>

      <div className="teacher-card">
        <div className="card-title">📋 Список ачивок</div>
        {achievements.map(a => <div key={a.id} className="item-card"><div>🏆 {a.title}</div><div>{a.description} | +{a.xp} XP</div></div>)}
        {achievements.length === 0 && <div>Нет созданных ачивок</div>}
      </div>
    </div>
  );
}

export default TeacherAchievements;