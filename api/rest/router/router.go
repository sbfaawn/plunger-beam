package router

import (
	"plunger-beam/internal/transport/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
	*http.BasicHandler
	*http.AccountHandler
	*http.MessageHandler
}

func NewRouter(basic *http.BasicHandler, account *http.AccountHandler, message *http.MessageHandler) *router {
	return &router{
		Engine:       gin.Default(),
		BasicHandler: basic,
		AccountHandler: account,
		MessageHandler: message,
	}
}

func (route *router) SetupRouter() {
	engine := route.Engine

	engine.NoRoute(route.NoRouteHandler)
	engine.NoMethod(route.NoMethodAllowed)

	group := engine.Group("/api/plungerapi")

	group.GET("/health", route.HealthCheck)

	account := group.Group("/account")
	message := group.Group("/message", route.CheckSession)

	// account/credential
	account.POST("/register", route.RegistrationHandler)
	account.POST("/login", route.LoginHandler)
	account.GET("/logout", route.LogoutHandler)
	account.GET("/refresh", route.RefreshTokenHandler)

	// message
	message.POST("", route.SendMessages)
	message.GET("", route.GetConversation)
	message.DELETE("", route.DeleteMessage)
}

func (router *router) Run(port string) {
	router.Engine.Run(":" + port)
}
