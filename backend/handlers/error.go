package handlers

import (
	"fmt"
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
)

func CustomError(c *gin.Context) {
	c.Error(&models.InternalError{
		Code: http.StatusBadGateway,
		Msg:  "Testing Error Endpoint",
	})
}

func StandardError(c *gin.Context) {
	c.Error(fmt.Errorf("Testing Error Endpoint"))
}
