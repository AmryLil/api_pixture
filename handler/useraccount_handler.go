package handler

import (
	"api/dto"
	"api/handleError"
	"api/helper"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type useraccounthandler struct {
	service services.Service
}

func NewUserAccount(service services.Service) *useraccounthandler {
	return &useraccounthandler{service}
}
func (h *useraccounthandler) RegisterHandler(c *gin.Context) {
	var user dto.RegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}
	if err := h.service.Register(user); err != nil {
		handleError.HandleError(c, err)
		return
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Success, please login",
	})
	c.JSON(http.StatusCreated, response)

}

func (h *useraccounthandler) GetDataUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	data, err := h.service.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        data.Id,
		"fullname":  data.Fullname,
		"username":  data.Username,
		"create_at": data.CreatedAt,
	})
}

func (h *useraccounthandler) LoginHandler(c *gin.Context) {
	var user dto.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError.HandleError(c, &handleError.BadRequestError{Message: err.Error()})
		return
	}

	res, err := h.service.Login(&user)

	if err != nil {
		handleError.HandleError(c, err)
		return
	}
	response := dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Login",
		Data:       res,
	}
	c.JSON(http.StatusOK, response)
}

func (h *useraccounthandler) GetAllUser(c *gin.Context) {
	userAccount, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"useraccount": userAccount})

}
