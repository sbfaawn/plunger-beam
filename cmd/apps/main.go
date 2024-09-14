package main

import (
	"plunger-beam/api/rest/router"
	"plunger-beam/internal/cache"
	"plunger-beam/internal/config"
	"plunger-beam/internal/database"
	"plunger-beam/internal/repository"
	"plunger-beam/internal/service"
	util "plunger-beam/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	group := server.Group("/api/v1")
	group.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"status":  "Success",
			"data":    "",
			"message": "",
		})
	})

	server.Run(":" + cfg.Server.Port)
}

func setUp() {
	config.Read()
	cfg := config.GetConfig()

	passwordEncryptor := util.NewPasswordEncryptor()

	mysqlOpt := database.MySqlOption{
		Address:     cfg.Database.Address,
		Username:    cfg.Database.Username,
		Password:    cfg.Database.Password,
		Port:        cfg.Database.Port,
		Database:    cfg.Database.Database,
		IsPopulated: cfg.Database.IsPopulated,
		IsMigrate:   cfg.Database.IsMigrate,
	}

	mysqlConn := database.NewMySqlConnection(mysqlOpt)
	mysqlConn.ConnectToDB()
	db := mysqlConn.GetDB()

	redisOpt := cache.RedisOption{
		Address:  cfg.Cache.Address,
		Port:     cfg.Cache.Port,
		Password: cfg.Cache.Password,
		DbNum:    cfg.Cache.DbNum,
	}

	redisConn := cache.NewRedisConnection(redisOpt)
	redisConn.ConnectToRedis()
	cli := redisConn.GetClient()

	accountRepository := repository.NewAccountRepository(db)
	messageRepository := repository.NewMessageRepository(db)
	sessionRepository := repository.NewSessionRepository(db, cli)

	accountService := service.NewAccountService(accountRepository, passwordEncryptor)
	messageService := service.NewMessageService(messageRepository)
	sessionService := service.NewSessionService(sessionRepository)

	r := router.NewRouter()
}
