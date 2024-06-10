package handler

import (
	"api/dto"
	"api/handleError"
	"api/helper"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userdetailshandler struct {
	userdetails_service services.UserDetailsService
}

func NewUserDetailsHandler(service services.UserDetailsService) *userdetailshandler {
	return &userdetailshandler{service}
}

func (h *userdetailshandler) GetDataUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	data, err := h.userdetails_service.GetDataByID(userID)
	if err != nil {
		handleError.HandleError(c, &handleError.InternalServerError{Message: err.Error()})
		return
	}

	userdetails := &dto.UserDetailsResponse{
		Telp:        data.Telp,
		Gender:      data.Gender,
		Location:    data.Location,
		Email:       data.Email,
		Description: data.Description,
	}

	c.JSON(http.StatusOK, gin.H{"data": userdetails})
}

func (h *userdetailshandler) AddUserDetails(c *gin.Context) {
	var userdetails dto.UserDetailsRequest

	if err := c.ShouldBindJSON(&userdetails); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	if err := h.userdetails_service.AddDetails(userdetails); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "User details success to add",
	})
	c.JSON(http.StatusCreated, response)
}
