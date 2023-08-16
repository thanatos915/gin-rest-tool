package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Tags 公共接口
// @Summary Ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
