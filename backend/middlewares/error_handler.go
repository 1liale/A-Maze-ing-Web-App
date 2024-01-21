package middlewares

import (
	"net/http"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // passing request

		// after request
		for _, e := range c.Errors {
			err := e.Err
			if privte_err, valid := err.(*models.InternalError); valid {
				c.JSON(privte_err.Code, gin.H{
					"code":    privte_err.Code,
					"message": privte_err.Msg,
				})
			} else {
				// always return with status code 500 otherwise
				code := http.StatusInternalServerError
				c.JSON(code,
					models.InternalError{
						Code:   code,
						Msg:    "Internal server error",
						Reason: err.Error(),
					},
				)
			}
		}
	}
}
