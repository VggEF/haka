package main

import (
	"log"
	"time"

	"student-app/internal/achievements"
	"student-app/internal/auth"
	"student-app/internal/chats"
	"student-app/internal/coins"
	"student-app/internal/events"
	"student-app/internal/games"
	"student-app/internal/middleware"
	"student-app/internal/news"
	"student-app/internal/notifications"
	"student-app/internal/schedule"
	"student-app/internal/users"
	"student-app/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// =========================
	// DATABASE
	// =========================

	dbConfig := &database.Config{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Password:        "postgres",
		DBName:          "student_app",
		SSLMode:         "disable",
		MaxOpenConns:    25,
		MaxIdleConns:    10,
		ConnMaxLifetime: 5 * time.Minute,
	}

	sqlDB, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	log.Println("✅ Подключено к PostgreSQL")

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Ошибка GORM:", err)
	}

	// =========================
	// MIGRATIONS
	// =========================

	if err := gormDB.AutoMigrate(
		&coins.ShopItem{},
		&coins.Transaction{},
		&coins.Purchase{},
	); err != nil {
		log.Println("Ошибка миграции:", err)
	}

	// =========================
	// WEBSOCKET
	// =========================

	wsHub := chats.NewHub()
	go wsHub.Run()

	// =========================
	// REPOSITORIES
	// =========================

	userRepo := users.NewRepository(sqlDB)
	scheduleRepo := schedule.NewRepository(sqlDB)
	eventsRepo := events.NewRepository(sqlDB)
	gamesRepo := games.NewRepository(sqlDB)
	chatsRepo := chats.NewRepository(sqlDB)
	notificationsRepo := notifications.NewRepository(sqlDB)
	newsRepo := news.NewRepository(sqlDB)
	achievementsRepo := achievements.NewRepository(sqlDB)

	// =========================
	// API CLIENTS
	// =========================

	scheduleAPIClient := schedule.NewAPIClient()

	// =========================
	// SERVICES
	// =========================

	authService := auth.NewService(auth.NewRepository(sqlDB))
	userService := users.NewService(userRepo)
	scheduleService := schedule.NewService(
		scheduleRepo,
		scheduleAPIClient,
	)

	eventsService := events.NewService(eventsRepo)
	gamesService := games.NewService(gamesRepo)
	chatsService := chats.NewService(chatsRepo)

	notificationsService := notifications.NewService(
		notificationsRepo,
	)

	coinsService := coins.NewService(gormDB)
	newsService := news.NewService(newsRepo)
	achievementsService := achievements.NewService(achievementsRepo)

	// =========================
	// HANDLERS
	// =========================

	authHandler := auth.NewHandler(authService)
	userHandler := users.NewHandler(userService)
	scheduleHandler := schedule.NewHandler(scheduleService)
	eventsHandler := events.NewHandler(eventsService)
	gamesHandler := games.NewHandler(gamesService)
	chatsHandler := chats.NewHandler(chatsService)
	wsHandler := chats.NewWebSocketHandler(wsHub)

	notificationsHandler := notifications.NewHandler(
		notificationsService,
	)

	coinsHandler := coins.NewHandler(coinsService)
	newsHandler := news.NewHandler(newsService)
	achievementsHandler := achievements.NewHandler(achievementsService)

	// =========================
	// ROUTER
	// =========================

	r := gin.Default()

	// =========================
	// CORS - ПОЛНОСТЬЮ РАЗРЕШАЕМ ДЛЯ ТЕСТА
	// =========================

	// Обработка preflight (OPTIONS) запросов
	r.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.AbortWithStatus(204)
	})

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Временно разрешаем все источники
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	// =========================
	// SWAGGER
	// =========================

	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerfiles.Handler),
	)

	// =========================
	// HEALTH
	// =========================

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// =========================
	// ROOT
	// =========================

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})

	// =========================
	// PUBLIC ROUTES
	// =========================

	r.POST("/api/login", authHandler.Login)
	r.POST("/api/register", authHandler.Register)

	// shop
	r.GET("/api/coins/shop", coinsHandler.GetShopItems)

	// NEWS PUBLIC
	r.GET("/api/news", newsHandler.GetAll)
	r.GET("/api/news/pinned", newsHandler.GetPinned)
	r.GET("/api/news/:id", newsHandler.GetByID)

	// ACHIEVEMENTS PUBLIC
	r.GET("/api/achievements", achievementsHandler.GetAllAchievements)

	// =========================
	// AUTH ROUTES
	// =========================

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	{
		// =========================
		// WEBSOCKET
		// =========================

		api.GET("/ws", wsHandler.HandleWebSocket)

		// =========================
		// USERS
		// =========================

		api.GET("/users/me", userHandler.GetMe)
		api.PUT("/users/me", userHandler.UpdateMe)

		api.GET(
			"/users/profile/student",
			userHandler.GetStudentProfile,
		)

		api.PUT(
			"/users/profile/student",
			userHandler.UpdateStudentProfile,
		)

		api.GET(
			"/users/profile/teacher",
			userHandler.GetTeacherProfile,
		)

		api.PUT(
			"/users/profile/teacher",
			userHandler.UpdateTeacherProfile,
		)

		// =========================
		// ADMIN
		// =========================

		admin := api.Group("/admin")
		admin.Use(middleware.AdminOnlyMiddleware())

		{
			admin.GET("/users", userHandler.GetAll)
			admin.POST("/users", userHandler.CreateUser)
			admin.DELETE("/users/:id", userHandler.DeleteUser)

			admin.POST("/coins/add", coinsHandler.AddCoins)
			admin.POST("/coins/spend", coinsHandler.SpendCoins)

			// ACHIEVEMENTS ADMIN
			admin.POST("/achievements", achievementsHandler.CreateAchievement)
			admin.POST("/achievements/award", achievementsHandler.AwardAchievement)
			admin.DELETE("/achievements/:id", achievementsHandler.DeleteAchievement)
		}

		// =========================
		// NEWS ADMIN
		// =========================

		newsAdmin := api.Group("/news")
		newsAdmin.Use(
			middleware.TeacherOrAdminMiddleware(),
		)

		{
			newsAdmin.POST("", newsHandler.Create)
			newsAdmin.PUT("/:id", newsHandler.Update)
			newsAdmin.DELETE("/:id", newsHandler.Delete)
		}

		// =========================
		// COINS
		// =========================

		api.GET(
			"/coins/balance",
			coinsHandler.GetBalance,
		)

		api.GET(
			"/coins/transactions",
			coinsHandler.GetTransactions,
		)

		api.POST(
			"/coins/shop/:itemId/buy",
			coinsHandler.BuyItem,
		)

		// =========================
		// ACHIEVEMENTS PROTECTED
		// =========================

		api.GET("/achievements/me", achievementsHandler.GetMyAchievements)
		api.GET("/achievements/me/xp", achievementsHandler.GetMyXP)
		api.GET("/achievements/user/:user_id", achievementsHandler.GetUserAchievements)

		// =========================
		// SCHEDULE
		// =========================

		api.GET("/schedule", scheduleHandler.GetSchedule)

		api.GET(
			"/schedule/:id",
			scheduleHandler.GetScheduleByID,
		)

		api.GET(
			"/schedule/groups",
			scheduleHandler.GetGroups,
		)

		scheduleAdmin := api.Group("/schedule")

		scheduleAdmin.Use(
			middleware.TeacherOrAdminMiddleware(),
		)

		{
			scheduleAdmin.POST(
				"/",
				scheduleHandler.CreateSchedule,
			)

			scheduleAdmin.PUT(
				"/:id",
				scheduleHandler.UpdateSchedule,
			)

			scheduleAdmin.DELETE(
				"/:id",
				scheduleHandler.DeleteSchedule,
			)

			scheduleAdmin.POST(
				"/sync",
				scheduleHandler.SyncFromAPI,
			)
		}

		// =========================
		// EVENTS
		// =========================

		api.GET("/events", eventsHandler.GetAll)

		api.GET(
			"/events/:id",
			eventsHandler.GetByID,
		)

		api.POST(
			"/events/:id/register",
			eventsHandler.Register,
		)

		api.DELETE(
			"/events/:id/register",
			eventsHandler.Unregister,
		)

		eventsAdmin := api.Group("/events")

		eventsAdmin.Use(
			middleware.TeacherOrAdminMiddleware(),
		)

		{
			eventsAdmin.POST(
				"/",
				eventsHandler.Create,
			)

			eventsAdmin.PUT(
				"/:id",
				eventsHandler.Update,
			)

			eventsAdmin.DELETE(
				"/:id",
				eventsHandler.Delete,
			)
		}

		// =========================
		// GAMES
		// =========================

		api.GET("/games", gamesHandler.GetAllGames)

		api.GET(
			"/games/:id",
			gamesHandler.GetGameByID,
		)

		api.GET(
			"/games/:id/leaderboard",
			gamesHandler.GetLeaderboard,
		)

		api.POST(
			"/games/play",
			gamesHandler.SubmitResult,
		)

		api.GET(
			"/games/results",
			gamesHandler.GetMyResults,
		)

		api.GET(
			"/games/best/:id",
			gamesHandler.GetMyBestScore,
		)

		// =========================
		// CHATS
		// =========================

		api.GET("/chats", chatsHandler.GetUserChats)

		api.GET(
			"/chats/:id",
			chatsHandler.GetChatByID,
		)

		api.POST(
			"/chats/private/:userId",
			chatsHandler.CreatePrivateChat,
		)

		api.POST(
			"/chats/group",
			chatsHandler.CreateGroupChat,
		)

		api.POST(
			"/chats/:id/leave",
			chatsHandler.LeaveChat,
		)

		api.GET(
			"/chats/:id/messages",
			chatsHandler.GetMessages,
		)

		api.POST(
			"/chats/:id/messages",
			chatsHandler.SendMessage,
		)

		api.PUT(
			"/chats/messages/:id/read",
			chatsHandler.MarkAsRead,
		)

		api.DELETE(
			"/chats/messages/:id",
			chatsHandler.DeleteMessage,
		)

		// =========================
		// NOTIFICATIONS
		// =========================

		api.GET(
			"/notifications",
			notificationsHandler.GetMyNotifications,
		)

		api.GET(
			"/notifications/unread",
			notificationsHandler.GetUnreadCount,
		)

		api.PUT(
			"/notifications/:id/read",
			notificationsHandler.MarkAsRead,
		)

		api.PUT(
			"/notifications/read-all",
			notificationsHandler.MarkAllAsRead,
		)

		api.DELETE(
			"/notifications/:id",
			notificationsHandler.DeleteNotification,
		)
	}

	// =========================
	// START SERVER
	// =========================

	log.Println("🚀 Сервер запущен на :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска:", err)
	}
}
