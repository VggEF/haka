package configs

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server      ServerConfig      `yaml:"server"`
	Database    DatabaseConfig    `yaml:"database"`
	Redis       RedisConfig       `yaml:"redis"`
	JWT         JWTConfig         `yaml:"jwt"`
	CORS        CORSConfig        `yaml:"cors"`
	RateLimit   RateLimitConfig   `yaml:"rate_limit"`
	Log         LogConfig         `yaml:"log"`
	Upload      UploadConfig      `yaml:"upload"`
	ExternalAPI ExternalAPIConfig `yaml:"external_apis"`
	WebSocket   WebSocketConfig   `yaml:"websocket"`
	Email       EmailConfig       `yaml:"email"`
}

type ServerConfig struct {
	Port         string        `yaml:"port"`
	Mode         string        `yaml:"mode"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DatabaseConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DBName          string        `yaml:"dbname"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret      string        `yaml:"secret"`
	ExpireHours time.Duration `yaml:"expire_hours"`
}

type CORSConfig struct {
	AllowOrigins     []string `yaml:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

type RateLimitConfig struct {
	Enabled           bool `yaml:"enabled"`
	RequestsPerMinute int  `yaml:"requests_per_minute"`
	Burst             int  `yaml:"burst"`
}

type LogConfig struct {
	Level    string `yaml:"level"`
	Output   string `yaml:"output"`
	FilePath string `yaml:"file_path"`
}

type UploadConfig struct {
	MaxSize      int      `yaml:"max_size"`
	AllowedTypes []string `yaml:"allowed_types"`
	UploadPath   string   `yaml:"upload_path"`
}

type ExternalAPIConfig struct {
	ScheduleAPI struct {
		BaseURL string        `yaml:"base_url"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"schedule_api"`
	LibraryAPI struct {
		BaseURL string        `yaml:"base_url"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"library_api"`
}

type WebSocketConfig struct {
	PingPeriod     time.Duration `yaml:"ping_period"`
	PongWait       time.Duration `yaml:"pong_wait"`
	WriteWait      time.Duration `yaml:"write_wait"`
	MaxMessageSize int64         `yaml:"max_message_size"`
}

type EmailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

var AppConfig *Config

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла конфига: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфига: %w", err)
	}

	// Загружаем переменные окружения (приоритет выше)
	loadEnvOverrides(&config)

	AppConfig = &config
	return &config, nil
}

func loadEnvOverrides(config *Config) {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		config.Server.Mode = mode
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.Database.Password = dbPassword
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.JWT.Secret = jwtSecret
	}
}

// GetDatabaseURL возвращает строку подключения к БД
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Password, c.Database.DBName, c.Database.SSLMode,
	)
}

// GetRedisURL возвращает строку подключения к Redis
func (c *Config) GetRedisURL() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}
