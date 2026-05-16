import React, { useState, useEffect } from 'react';
import './MiniGames.css';

// Игра 1: Угадай термин
const TERMS = [
  { term: "React", definition: "Библиотека JavaScript для создания пользовательских интерфейсов" },
  { term: "Redux", definition: "Библиотека для управления состоянием приложения" },
  { term: "Git", definition: "Система контроля версий" },
  { term: "API", definition: "Интерфейс программирования приложений" },
  { term: "SQL", definition: "Язык структурированных запросов" }
];

// Игра 2: Найди аудиторию
const AUDITORIUMS = [
  { number: "Б-201", floor: "2 этаж", direction: "Направо от лифта", building: "Главный корпус" },
  { number: "Б-303", floor: "3 этаж", direction: "Налево от лестницы", building: "Главный корпус" },
  { number: "Б-208", floor: "2 этаж", direction: "В конец коридора", building: "Главный корпус" }
];

// Игра 3: Собери расписание
const SCHEDULE_PUZZLE = [
  { time: "09:00", subject: "Программирование", place: "Б-201" },
  { time: "11:00", subject: "Математика", place: "Б-303" },
  { time: "13:00", subject: "Английский", place: "Б-208" }
];

function MiniGames() {
  const [activeGame, setActiveGame] = useState(null);
  const [gameScore, setGameScore] = useState(0);
  const [currentTerm, setCurrentTerm] = useState(null);
  const [userDefinition, setUserDefinition] = useState("");
  const [feedback, setFeedback] = useState("");
  const [puzzleOrder, setPuzzleOrder] = useState([]);
  const [dragIndex, setDragIndex] = useState(null);
  const [auditoriumGuess, setAuditoriumGuess] = useState("");
  const [currentAuditorium, setCurrentAuditorium] = useState(null);

  useEffect(() => {
    const savedScore = localStorage.getItem('miniGamesScore');
    if (savedScore) setGameScore(parseInt(savedScore));
  }, []);

  const addPoints = (points) => {
    const newScore = gameScore + points;
    setGameScore(newScore);
    localStorage.setItem('miniGamesScore', newScore);
    setFeedback(`🎉 +${points} очков!`);
    setTimeout(() => setFeedback(""), 2000);
  };

  // Игра 1: Угадай термин
  const startGuessGame = () => {
    const randomTerm = TERMS[Math.floor(Math.random() * TERMS.length)];
    setCurrentTerm(randomTerm);
    setActiveGame("guess");
    setUserDefinition("");
  };

  const checkGuess = () => {
    if (userDefinition.toLowerCase().includes(currentTerm.definition.toLowerCase().slice(0, 20))) {
      addPoints(10);
      startGuessGame();
    } else {
      setFeedback(`❌ Неправильно! Правильный ответ: ${currentTerm.definition}`);
      setTimeout(() => setFeedback(""), 3000);
    }
  };

  // Игра 2: Найди аудиторию
  const startFindAuditorium = () => {
    const randomAud = AUDITORIUMS[Math.floor(Math.random() * AUDITORIUMS.length)];
    setCurrentAuditorium(randomAud);
    setActiveGame("find");
    setAuditoriumGuess("");
  };

  const checkAuditorium = () => {
    if (auditoriumGuess === currentAuditorium.number) {
      addPoints(15);
      startFindAuditorium();
    } else {
      setFeedback(`❌ Не угадал! Подсказка: ${currentAuditorium.direction}, ${currentAuditorium.floor}`);
      setTimeout(() => setFeedback(""), 3000);
    }
  };

  // Игра 3: Собери расписание
  const startPuzzle = () => {
    const shuffled = [...SCHEDULE_PUZZLE].sort(() => Math.random() - 0.5);
    setPuzzleOrder(shuffled);
    setActiveGame("puzzle");
  };

  const handleDragStart = (index) => {
    setDragIndex(index);
  };

  const handleDragOver = (e) => {
    e.preventDefault();
  };

  const handleDrop = (targetIndex) => {
    if (dragIndex === null) return;
    const newOrder = [...puzzleOrder];
    const draggedItem = newOrder[dragIndex];
    newOrder.splice(dragIndex, 1);
    newOrder.splice(targetIndex, 0, draggedItem);
    setPuzzleOrder(newOrder);
    setDragIndex(null);
    
    // Проверка правильности
    const isCorrect = newOrder.every((item, idx) => item.time === SCHEDULE_PUZZLE[idx].time);
    if (isCorrect && newOrder.length === SCHEDULE_PUZZLE.length) {
      addPoints(25);
      setTimeout(() => startPuzzle(), 2000);
    }
  };

  return (
    <div className="minigames-container">
      <h2 className="minigames-title">🎮 Mini-Games для прокрастинации</h2>
      
      <div className="minigames-score">
        <span>🏆 Очки за игры: {gameScore}</span>
      </div>

      <div className="minigames-grid">
        <div className="game-card" onClick={startGuessGame}>
          <div className="game-icon">📚</div>
          <h3>Угадай термин</h3>
          <p>Проверь свои знания IT-терминов</p>
        </div>
        <div className="game-card" onClick={startFindAuditorium}>
          <div className="game-icon">🔍</div>
          <h3>Найди аудиторию</h3>
          <p>Ориентируйся в корпусах</p>
        </div>
        <div className="game-card" onClick={startPuzzle}>
          <div className="game-icon">🧩</div>
          <h3>Собери расписание</h3>
          <p>Правильный порядок пар</p>
        </div>
      </div>

      {feedback && <div className="game-feedback">{feedback}</div>}

      {/* Игра 1 */}
      {activeGame === "guess" && currentTerm && (
        <div className="game-modal">
          <div className="game-modal-content">
            <h3>Угадай термин</h3>
            <div className="term-question">{currentTerm.definition}</div>
            <input type="text" placeholder="Введите термин" value={userDefinition} onChange={(e) => setUserDefinition(e.target.value)} />
            <button onClick={checkGuess}>Проверить</button>
            <button onClick={() => setActiveGame(null)}>Закрыть</button>
          </div>
        </div>
      )}

      {/* Игра 2 */}
      {activeGame === "find" && currentAuditorium && (
        <div className="game-modal">
          <div className="game-modal-content">
            <h3>Найди аудиторию</h3>
            <div className="auditorium-clue">
              <p>🔍 Подсказки:</p>
              <p>🏢 {currentAuditorium.building}</p>
              <p>📌 {currentAuditorium.floor}</p>
              <p>➡️ {currentAuditorium.direction}</p>
            </div>
            <input type="text" placeholder="Номер аудитории (например, Б-201)" value={auditoriumGuess} onChange={(e) => setAuditoriumGuess(e.target.value)} />
            <button onClick={checkAuditorium}>Проверить</button>
            <button onClick={() => setActiveGame(null)}>Закрыть</button>
          </div>
        </div>
      )}

      {/* Игра 3 */}
      {activeGame === "puzzle" && puzzleOrder.length > 0 && (
        <div className="game-modal puzzle-modal">
          <div className="game-modal-content">
            <h3>Собери расписание в правильном порядке</h3>
            <div className="puzzle-list">
              {puzzleOrder.map((item, idx) => (
                <div key={idx} className="puzzle-item" draggable onDragStart={() => handleDragStart(idx)} onDragOver={handleDragOver} onDrop={() => handleDrop(idx)}>
                  <span className="puzzle-time">{item.time}</span>
                  <span className="puzzle-subject">{item.subject}</span>
                  <span className="puzzle-place">{item.place}</span>
                </div>
              ))}
            </div>
            <button onClick={() => setActiveGame(null)}>Закрыть</button>
          </div>
        </div>
      )}
    </div>
  );
}

export default MiniGames;