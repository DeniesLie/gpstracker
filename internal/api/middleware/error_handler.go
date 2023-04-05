package middleware

import (
	"errors"
	"net/http"

	"github.com/DeniesLie/gpstracker/internal/api/controllers/response"
	"github.com/DeniesLie/gpstracker/internal/core/service"
	"github.com/DeniesLie/gpstracker/internal/core/validation"
	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
func ErrorHandler(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()

	if err != nil {
		notFoundErr := &service.NotFoundError{}
		validationErr := &validation.ValidationError{}
		businessErr := &service.BusinessError{}

		switch {
		case errors.As(err, notFoundErr):
			c.JSON(http.StatusNotFound, response.Error(response.NotFoundMessage))

		case errors.As(err, validationErr):
			c.JSON(http.StatusBadRequest, response.InvalidData(validationErr.Res))

		case errors.As(err, businessErr):
			c.JSON(http.StatusBadRequest, response.Error(businessErr.Message))

		default:
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}
}
