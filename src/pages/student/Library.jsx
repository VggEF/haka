import React from 'react';
import './Library.css';

function Library() {
  const books = [
    { id: 1, title: '📖 "Алгоритмы и структуры данных"', author: 'Т. Кормен', status: 'Доступно' },
    { id: 2, title: '📖 "React Native в действии"', author: 'А. Бэнкс', status: 'На руках до 20.06' },
    { id: 3, title: '📖 "Математический анализ"', author: 'В.А. Зорич', status: 'Доступно' },
  ];

  return (
    <div className="library-container">
      <h2 className="library-title">📚 Библиотека</h2>
      <div className="library-list">
        {books.map(book => (
          <div key={book.id} className="book-item">
            <div className="book-title">{book.title}</div>
            <div className="book-author">{book.author}</div>
            <div className="book-status">{book.status}</div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Library;