package middleware

import (
	"net/http"

	"github.com/DeniesLie/gpstracker/internal/api/controllers/response"
	"github.com/DeniesLie/gpstracker/internal/core/service"
	"github.com/DeniesLie/gpstracker/internal/core/validation"
	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
func ErrorHandler(c *gin.Context) {
	c.Next()
	ginErr := c.Errors.Last()

	if ginErr != nil {
		err := ginErr.Err
		_, isNotFoundErr := err.(service.NotFoundError)
		_, isValidationErr := err.(validation.ValidationError)
		_, isBusinessErr := err.(service.BusinessError)

		switch {
		case isNotFoundErr:
			c.JSON(http.StatusNotFound, response.Error(response.NotFoundMessage))

		case isValidationErr:
			c.JSON(http.StatusBadRequest, response.InvalidData(err.Error()))

		case isBusinessErr:
			c.JSON(http.StatusBadRequest, response.Error(err.Error()))

		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
