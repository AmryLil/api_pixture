package services

import (
	"api/dto"
	"api/handleError"
	"api/helper"
	models "api/models"
	"api/repositories"
)

type Service interface {
	GetAll() ([]models.UserAccounts, error)
	GetByID(id any) (models.UserAccounts, error)
	// Update(id int) (models.UserAccounts, error)
	// Delete(id int) (models.UserAccounts, error)
	Register(req dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	FindBy(field string, valueField string, userAccount models.UserAccounts) (models.UserAccounts, error)
}

type service struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *service {
	return &service{repository}
}

func (s *service) Register(req dto.RegisterRequest) error {
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		return &handleError.BadRequestError{Message: "username already exist"}
	}
	if req.Password != req.ConfirmPassword {
		return &handleError.BadRequestError{Message: "password not match"}
	}
	if len(req.Password) < 8 {
		return &handleError.BadRequestError{Message: "password at least 8 character"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &handleError.BadRequestError{Message: err.Error()}
	}

	user := models.UserAccounts{
		Username: req.Username,
		Fullname: req.Fullname,
		Password: passwordHash,
	}

	err = s.repository.Create(user)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil

}

func (s *service) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var user models.UserAccounts
	models, err := s.repository.FindBy("username", req.Username, user)
	if err != nil {
		return nil, &handleError.NotFoundError{Message: "username not found"}
	}

	if err := helper.VerifyPassword(models.Password, req.Password); err != nil {
		return nil, &handleError.NotFoundError{Message: "your password is wrong"}

	}

	token, err := helper.GenerateToken(models)
	if err != nil {
		return nil, &handleError.InternalServerError{Message: err.Error()}
	}

	userResponse := &dto.LoginResponse{
		Name: models.Username, ID: models.Id,

		Token: token,
	}
	return userResponse, err

}

func (s *service) FindBy(field string, valueField string, userAccount models.UserAccounts) (models.UserAccounts, error) {
	return s.repository.FindBy(field, valueField, userAccount)
}

func (s *service) GetAll() ([]models.UserAccounts, error) {
	useraccount, err := s.repository.GetAll()
	return useraccount, err
}

func (s *service) GetByID(id any) (models.UserAccounts, error) {
	return s.repository.GetByID(id)
}
