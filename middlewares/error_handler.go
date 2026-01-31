package middlewares

import (
	"kasir-api/common/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				cutomError := dtos.ErrorResponse{}
				switch err := err.(type) {
				case dtos.ErrorResponse:
					// This section for handling from own panic/error
					cutomError = err
				default:
					// This section for handling unexpected error
					cutomError = dtos.ErrorResponse{
						ErrorCode: http.StatusInternalServerError,
						Message: dtos.Response{
							Status: dtos.BaseResponse{
								Success: false,
								Message: "Internal Server Error",
							},
							Data: err,
						}}
				}
				c.JSON(cutomError.ErrorCode, cutomError.Message)
				c.Abort()
			}
		}()
		c.Next()
	}
}
