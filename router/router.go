package router

import (
	"lecture/go-wallet/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", controller.Health)
	router.POST("/mnemonics", controller.NewMnemonic)
	router.POST("/wallets", controller.NewWallet)
	router.GET("/balance", controller.GetBalance)
	return router
}
