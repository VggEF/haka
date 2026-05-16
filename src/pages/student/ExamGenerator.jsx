import React, { useState } from 'react';
import './ExamGenerator.css';

const QUESTIONS_DB = {
  "Программирование": [
    { question: "Что такое переменная?", answer: "Именованная область памяти для хранения данных", difficulty: "easy" },
    { question: "Что такое цикл for?", answer: "Конструкция для повторения блока кода определенное количество раз", difficulty: "easy" },
    { question: "Что такое функция?", answer: "Блок кода, который можно вызывать многократно", difficulty: "medium" },
    { question: "Что такое рекурсия?", answer: "Функция, вызывающая саму себя", difficulty: "hard" },
    { question: "Что такое ООП?", answer: "Парадигма программирования, основанная на объектах и классах", difficulty: "hard" }
  ],
  "Математика": [
    { question: "Что такое производная?", answer: "Скорость изменения функции", difficulty: "medium" },
    { question: "Что такое интеграл?", answer: "Площадь под кривой функции", difficulty: "hard" }
  ]
};

function ExamGenerator() {
  const [subject, setSubject] = useState("Программирование");
  const [currentQuestion, setCurrentQuestion] = useState(null);
  const [showAnswer, setShowAnswer] = useState(false);
  const [score, setScore] = useState(0);
  const [asked, setAsked] = useState([]);

  const generateQuestion = () => {
    const available = QUESTIONS_DB[subject].filter((_, idx) => !asked.includes(idx));
    if (available.length === 0) {
      alert("🎉 Ты ответил на все вопросы! Начни заново.");
      setAsked([]);
      const randomIndex = Math.floor(Math.random() * QUESTIONS_DB[subject].length);
      setCurrentQuestion(QUESTIONS_DB[subject][randomIndex]);
      setShowAnswer(false);
      return;
    }
    const randomIndex = Math.floor(Math.random() * available.length);
    const question = available[randomIndex];
    const originalIndex = QUESTIONS_DB[subject].findIndex(q => q.question === question.question);
    setCurrentQuestion(question);
    setAsked([...asked, originalIndex]);
    setShowAnswer(false);
  };

  const checkAnswer = (userAnswer) => {
    if (!currentQuestion) return;
    if (userAnswer.toLowerCase().trim() === currentQuestion.answer.toLowerCase().trim()) {
      const points = currentQuestion.difficulty === "easy" ? 10 : currentQuestion.difficulty === "medium" ? 20 : 30;
      setScore(score + points);
      alert(`✅ Правильно! +${points} очков`);
      generateQuestion();
    } else {
      alert(`❌ Неправильно! Правильный ответ: ${currentQuestion.answer}`);
    }
  };

  return (
    <div className="exam-container">
      <h2>🎓 Генератор вопросов для экзамена</h2>
      
      <div className="exam-score">
        <span>🏆 Очки: {score}</span>
      </div>

      <div className="exam-controls">
        <select value={subject} onChange={(e) => setSubject(e.target.value)}>
          <option>Программирование</option>
          <option>Математика</option>
        </select>
        <button onClick={generateQuestion}>🎲 Сгенерировать вопрос</button>
      </div>

      {currentQuestion && (
        <div className="exam-question">
          <div className="question-text">{currentQuestion.question}</div>
          <div className="question-difficulty">
            {currentQuestion.difficulty === "easy" && "🟢 Легкий"}
            {currentQuestion.difficulty === "medium" && "🟡 Средний"}
            {currentQuestion.difficulty === "hard" && "🔴 Сложный"}
          </div>
          <input type="text" className="exam-input" placeholder="Твой ответ..." onKeyPress={(e) => e.key === 'Enter' && checkAnswer(e.target.value)} />
          <button className="exam-check-btn" onClick={() => checkAnswer(document.querySelector('.exam-input').value)}>Проверить</button>
        </div>
      )}
    </div>
  );
}

export default ExamGenerator;