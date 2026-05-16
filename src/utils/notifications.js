// Умные уведомления
export const checkDeadlines = (deadlines) => {
  const today = new Date();
  const notifications = [];
  
  deadlines.forEach(deadline => {
    const deadlineDate = new Date(deadline.date);
    const diffDays = Math.ceil((deadlineDate - today) / (1000 * 60 * 60 * 24));
    
    if (diffDays === 1 && deadline.status !== 'completed') {
      notifications.push(`⚠️ Завтра дедлайн: "${deadline.title}"!`);
    }
    if (diffDays === 0 && deadline.status !== 'completed') {
      notifications.push(`🔥 Сегодня последний день: "${deadline.title}"! Сдай и получи +20 XP!`);
    }
  });
  
  return notifications;
};

// Внутренняя валюта - Коины
const COINS_KEY = 'userCoins';

export const getCoins = () => {
  const saved = localStorage.getItem(COINS_KEY);
  return saved ? parseInt(saved) : 0;
};

export const addCoins = (amount) => {
  const current = getCoins();
  localStorage.setItem(COINS_KEY, current + amount);
  return current + amount;
};

export const spendCoins = (amount) => {
  const current = getCoins();
  if (current >= amount) {
    localStorage.setItem(COINS_KEY, current - amount);
    return true;
  }
  return false;
};

// Магазин за коины
export const SHOP_ITEMS = [
  { id: 1, name: "Скидка в столовой", price: 50, icon: "🍕", description: "Скидка 10% в столовой" },
  { id: 2, name: "Билет на мероприятие", price: 100, icon: "🎫", description: "Билет на любое мероприятие" },
  { id: 3, name: "Принт мерча", price: 200, icon: "👕", description: "Фирменный мерч университета" },
  { id: 4, name: "Закрытая вечеринка", price: 500, icon: "🎉", description: "Билет на закрытую вечеринку" }
];