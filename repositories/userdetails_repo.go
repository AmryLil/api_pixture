package repositories

import (
	models "api/models"

	"gorm.io/gorm"
)

type UserDetailsRepo interface {
	AddDetails(userDetails models.UserDetails) error
	GetDataByID(id any) (models.UserDetails, error)
}

type userdetails_repo struct {
	db *gorm.DB
}

func NewUserDetailsRepo(db *gorm.DB) *userdetails_repo {
	return &userdetails_repo{db}
}

func (r *userdetails_repo) AddDetails(userDetails models.UserDetails) error {
	err := r.db.Create(&userDetails).Error
	return err
}

func (r *userdetails_repo) GetDataByID(id any) (models.UserDetails, error) {
	var userdetails models.UserDetails
	err := r.db.Find(&userdetails, id).Error
	return userdetails, err
}
