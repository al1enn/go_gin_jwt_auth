package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// Profile godoc
// @Summary Валидация токена
// @Description Проверяет действительность JWT токена
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/profile [get]
func Profile(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}