package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (a *appError) Error() string {
	return fmt.Sprintf("code: %d, error: %s", a.Code, a.Message)
}

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		log.Println("Handle APP error")
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *appError
			switch err.(type) {
			case *appError:
				parsedError = err.(*appError)
			default:
				parsedError = &appError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			// Put the error into response
			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			// or c.AbortWithStatusJSON(parsedError.Code, parsedError)
			return
		}

	}
}
