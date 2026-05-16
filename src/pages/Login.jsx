import React, { useState } from 'react';
import './Login.css';

function Login({ onLogin }) {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    setError('');

    try {
      const response = await fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          login,
          password,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        setError(data.error || 'Ошибка входа');
        return;
      }

      localStorage.setItem('token', data.token);
      localStorage.setItem('userRole', data.user.role);
      localStorage.setItem('userLogin', data.user.login);
      localStorage.setItem('userName', data.user.name);

      onLogin(
        data.user.login,
        data.user.role,
        data.user.name
      );
    } catch (err) {
      setError('Сервер недоступен');
    }
  };

  return (
    <div className="login-container">
      <div className="login-form">
        <h1 className="login-title">Авторизация</h1>

        <div className="login-info-box">
          <div>Тестовые аккаунты:</div>

          <div className="demo-users">
            <div>admin / admin123</div>
            <div>teacher / teacher123</div>
            <div>23-ПМбо-014 / student123</div>
          </div>
        </div>

        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Логин"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            className="login-input"
          />

          <input
            type="password"
            placeholder="Пароль"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="login-input"
          />

          {error && (
            <div className="login-error">
              {error}
            </div>
          )}

          <button type="submit" className="login-button">
            Войти
          </button>
        </form>
      </div>
    </div>
  );
}

export default Login;