import React, { useEffect, useState } from 'react';
import './News.css';

const API_URL = 'http://localhost:8080';

function News() {
  const [news, setNews] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchNews = async () => {
    try {
      const response = await fetch(`${API_URL}/api/news`);

      const data = await response.json();

      console.log('NEWS:', data);

      setNews(data.data || []);
    } catch (error) {
      console.error('Ошибка загрузки новостей:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchNews();
  }, []);

  if (loading) {
    return <div>Загрузка...</div>;
  }

  return (
    <div className="news-page">
      <h1>Новости</h1>

      {news.length === 0 ? (
        <p>Новостей пока нет</p>
      ) : (
        <div className="news-list">
          {news.map((item) => (
            <div className="news-card" key={item.id}>
              {item.image_url && (
                <img
                  src={item.image_url}
                  alt={item.title}
                  className="news-image"
                />
              )}

              <h2>{item.title}</h2>

              <p>{item.short_text}</p>

              <div className="news-footer">
                <span>{item.category}</span>

                <span>
                  {new Date(item.date).toLocaleDateString('ru-RU')}
                </span>
              </div>

              {item.is_pinned && (
                <div className="pinned">
                  📌 Закреплено
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default News;