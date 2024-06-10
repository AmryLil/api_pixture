package services

import (
	"api/dto"
	"api/handleError"
	models "api/models"
	"api/repositories"
)

type UserDetailsService interface {
	AddDetails(req dto.UserDetailsRequest) error
	GetDataByID(id any) (models.UserDetails, error)
}

type userdetails_service struct {
	repository repositories.UserDetailsRepo
}

func NewUserDetailsService(repository repositories.UserDetailsRepo) *userdetails_service {
	return &userdetails_service{repository}
}

func (s *userdetails_service) AddDetails(req dto.UserDetailsRequest) error {

	userDetails := models.UserDetails{
		Id:          req.Id,
		Telp:        req.Telp,
		Gender:      req.Gender,
		Location:    req.Location,
		Email:       req.Email,
		Description: req.Description,
		CreatedAt:   req.CreatedAt,
	}

	err := s.repository.AddDetails(userDetails)
	if err != nil {
		return &handleError.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *userdetails_service) GetDataByID(id any) (models.UserDetails, error) {
	return s.repository.GetDataByID(id)
}
