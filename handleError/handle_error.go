package handleError

import (
	"api/dto"
	"api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnathorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:   err.Error(),
	})
	c.JSON(statusCode, response)
}
