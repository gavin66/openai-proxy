package main

import (
	"openai-proxy/internal/handlers"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8890"
	}
	handler := gin.Default()

	handler.Any("/*path", handlers.Proxy)

	endless.ListenAndServe(os.Getenv("HOST")+":"+PORT, handler)
}
