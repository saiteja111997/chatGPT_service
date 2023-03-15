package server

import "github.com/gin-gonic/gin"

// Enable CORS to set the header as origin
func EnableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Header("Access-Control-Allow-Origin", "https://projectai-ba364.web.app")
}
