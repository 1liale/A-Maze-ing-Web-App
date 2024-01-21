package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1liale/maze-backend/handlers"
	"github.com/1liale/maze-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestError_Custom(t *testing.T) {
	record := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(record) // create gin context for record
	engine.Use(middlewares.ErrorHandler())

	engine.GET("/", handlers.CustomError)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	engine.ServeHTTP(record, ctx.Request)

	assert.IsType(t, http.StatusOK, record.Code)
	assert.NotEqual(t, http.StatusOK, record.Code)
}

func TestError_Standard(t *testing.T) {
	record := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(record) // create gin context for record
	engine.Use(middlewares.ErrorHandler())

	engine.GET("/", handlers.StandardError)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	engine.ServeHTTP(record, ctx.Request)

	assert.IsType(t, http.StatusOK, record.Code)
	assert.NotEqual(t, http.StatusOK, record.Code)
}
