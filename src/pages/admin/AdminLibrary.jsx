import React, { useState, useEffect } from 'react';
import './AdminLibrary.css';

function AdminLibrary() {
  const [books, setBooks] = useState([]);
  const [newBook, setNewBook] = useState({ title: "", author: "", publisher: "", year: "", description: "" });

  useEffect(() => {
    const saved = localStorage.getItem('libraryBooks');
    if (saved) setBooks(JSON.parse(saved));
  }, []);

  const addBook = () => {
    if (!newBook.title || !newBook.author) {
      alert("Заполните название и автора");
      return;
    }
    const book = {
      id: Date.now(),
      ...newBook,
      type: "book"
    };
    const updated = [...books, book];
    setBooks(updated);
    localStorage.setItem('libraryBooks', JSON.stringify(updated));
    setNewBook({ title: "", author: "", publisher: "", year: "", description: "" });
    alert(`Книга "${book.title}" добавлена!`);
  };

  const deleteBook = (id) => {
    const updated = books.filter(b => b.id !== id);
    setBooks(updated);
    localStorage.setItem('libraryBooks', JSON.stringify(updated));
  };

  return (
    <div className="admin-page">
      <h2>📚 Управление библиотекой</h2>
      
      <div className="add-form">
        <input type="text" placeholder="Название книги" value={newBook.title} onChange={(e) => setNewBook({...newBook, title: e.target.value})} />
        <input type="text" placeholder="Автор" value={newBook.author} onChange={(e) => setNewBook({...newBook, author: e.target.value})} />
        <input type="text" placeholder="Издательство" value={newBook.publisher} onChange={(e) => setNewBook({...newBook, publisher: e.target.value})} />
        <input type="text" placeholder="Год" value={newBook.year} onChange={(e) => setNewBook({...newBook, year: e.target.value})} />
        <input type="text" placeholder="Описание" value={newBook.description} onChange={(e) => setNewBook({...newBook, description: e.target.value})} />
        <button onClick={addBook}>+ Добавить</button>
      </div>

      <div className="items-list">
        {books.length === 0 && <div className="empty-state">Книг пока нет</div>}
        {books.map(book => (
          <div key={book.id} className="item-card">
            <div>📖 <strong>{book.title}</strong> - {book.author}</div>
            <div>🏢 {book.publisher} ({book.year})</div>
            <div>{book.description}</div>
            <button onClick={() => deleteBook(book.id)}>🗑️</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default AdminLibrary;