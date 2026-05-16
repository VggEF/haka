import React, { useEffect, useState } from 'react';
import './AdminNews.css';

const API_BASE_URL = 'http://localhost:8080';

function AdminNews() {
  const [news, setNews] = useState([]);
  const [loading, setLoading] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [editingId, setEditingId] = useState(null);

  const [formData, setFormData] = useState({
    title: '',
    short_text: '',
    full_text: '',
    image_url: '',
    category: 'Новость',
    is_pinned: false,
  });

  useEffect(() => {
    fetchNews();
  }, []);

  const fetchNews = async () => {
    try {
      setLoading(true);

      const response = await fetch(
        `${API_BASE_URL}/api/news`
      );

      const data = await response.json();

      console.log('NEWS:', data);

      if (Array.isArray(data)) {
        setNews(data);
      } else {
        setNews(data.data || []);
      }

    } catch (error) {
      console.error('Ошибка загрузки:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e) => {
    const { name, value, type, checked } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]:
        type === 'checkbox'
          ? checked
          : value,
    }));
  };

  const resetForm = () => {
    setEditingId(null);

    setFormData({
      title: '',
      short_text: '',
      full_text: '',
      image_url: '',
      category: 'Новость',
      is_pinned: false,
    });

    setShowModal(false);
  };

  const createNews = async () => {

    console.log('BUTTON CLICKED');

    try {

      console.log('SENDING:', formData);

      const response = await fetch(
        'http://localhost:8080/api/news',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(formData),
        }
      );

      console.log('STATUS:', response.status);

      const text = await response.text();

      console.log('RESPONSE:', text);

    } catch (error) {

      console.error('FETCH ERROR:', error);

    }
  };

  const updateNews = async () => {
    try {
      const response = await fetch(
        `${API_BASE_URL}/api/news/${editingId}`,
        {
          method: 'PUT',
          headers: {
            'Content-Type':
              'application/json',
          },
          body: JSON.stringify(formData),
        }
      );

      const data =
        await response.json();

      if (!response.ok) {
        throw new Error(
          data.error ||
            'Ошибка обновления'
        );
      }

      alert('Новость обновлена');

      await fetchNews();

      resetForm();

    } catch (error) {
      console.error(error);

      alert(error.message);
    }
  };

  const deleteNews = async (id) => {
    const confirmed =
      window.confirm(
        'Удалить новость?'
      );

    if (!confirmed) return;

    try {
      const response = await fetch(
        `${API_BASE_URL}/api/news/${id}`,
        {
          method: 'DELETE',
        }
      );

      if (!response.ok) {
        throw new Error(
          'Ошибка удаления'
        );
      }

      alert('Новость удалена');

      await fetchNews();

    } catch (error) {
      console.error(error);

      alert(error.message);
    }
  };

  const editNews = (item) => {
    setEditingId(item.id);

    setFormData({
      title: item.title || '',
      short_text:
        item.short_text || '',
      full_text:
        item.full_text || '',
      image_url:
        item.image_url || '',
      category:
        item.category ||
        'Новость',
      is_pinned:
        item.is_pinned || false,
    });

    setShowModal(true);
  };

  return (
    <div className="admin-news-page">

      <div className="admin-header">
        <h1>
          Управление новостями
        </h1>

        <button
          className="create-btn"
          onClick={() =>
            setShowModal(true)
          }
        >
          ➕ Создать
        </button>
      </div>

      {loading && (
        <div>Загрузка...</div>
      )}

      <div className="admin-news-grid">
        {news.map((item) => (
          <div
            className="admin-news-card"
            key={item.id}
          >
            {item.image_url && (
              <img
                src={item.image_url}
                alt={item.title}
                className="admin-news-image"
              />
            )}

            <div className="admin-news-top">
              <span>
                {item.category}
              </span>

              {item.is_pinned && (
                <span className="pinned">
                  📌 Закреплено
                </span>
              )}
            </div>

            <h3>{item.title}</h3>

            <p>
              {item.short_text}
            </p>

            <small>
              {item.created_at
                ? new Date(
                    item.created_at
                  ).toLocaleDateString(
                    'ru-RU'
                  )
                : ''}
            </small>

            <div className="admin-news-actions">
              <button
                onClick={() =>
                  editNews(item)
                }
              >
                ✏️ Изменить
              </button>

              <button
                onClick={() =>
                  deleteNews(item.id)
                }
              >
                🗑️ Удалить
              </button>
            </div>
          </div>
        ))}
      </div>

      {showModal && (
        <div
          className="modal-overlay"
          onClick={resetForm}
        >
          <div
            className="news-modal"
            onClick={(e) =>
              e.stopPropagation()
            }
          >
            <h2>
              {editingId
                ? 'Редактирование'
                : 'Создание новости'}
            </h2>

            <input
              type="text"
              name="title"
              placeholder="Заголовок"
              value={formData.title}
              onChange={
                handleInputChange
              }
            />

            <textarea
              name="short_text"
              placeholder="Краткое описание"
              value={
                formData.short_text
              }
              onChange={
                handleInputChange
              }
            />

            <textarea
              name="full_text"
              placeholder="Полный текст"
              value={
                formData.full_text
              }
              onChange={
                handleInputChange
              }
            />

            <input
              type="text"
              name="image_url"
              placeholder="URL картинки"
              value={
                formData.image_url
              }
              onChange={
                handleInputChange
              }
            />

            <select
              name="category"
              value={formData.category}
              onChange={
                handleInputChange
              }
            >
              <option value="Новость">
                Новость
              </option>

              <option value="Объявление">
                Объявление
              </option>

              <option value="Мероприятие">
                Мероприятие
              </option>
            </select>

            <label>
              <input
                type="checkbox"
                name="is_pinned"
                checked={
                  formData.is_pinned
                }
                onChange={
                  handleInputChange
                }
              />

              Закрепить
            </label>

            <div className="modal-actions">

              <button
                type="button"
                onClick={(e) => {
                  e.stopPropagation();

                  if (editingId) {
                    updateNews();
                  } else {
                    createNews();
                  }
                }}
              >
                {editingId
                  ? 'Сохранить'
                  : 'Создать'}
              </button>

              <button
                type="button"
                onClick={(e) => {
                  e.stopPropagation();
                  resetForm();
                }}
              >
                Отмена
              </button>

            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default AdminNews;