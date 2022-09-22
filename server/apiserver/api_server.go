package apiserver

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()
	logger = zapLogger.Named("apiserver")

	logger.Info("Starting Erupe")
}

func Start() {
	println("Starting api server")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/test", test)
	router.POST("/chat/cast", sendServerChatMessage)
	router.POST("/chat/broadcast")
	router.Run("localhost:8080")
}

// test endpoint
func test(c *gin.Context) {
	logger.Info("Test endpoint")

}

// Sends a message to a specific session
func sendServerChatMessage(c *gin.Context) {

}

// Sends a broadcast message to all sessions.
func sendBroadcastChatMessage(c *gin.Context) {

}
