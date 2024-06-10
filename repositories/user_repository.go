package repositories

import (
	models "api/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.UserAccounts, error)
	GetByID(id any) (models.UserAccounts, error)
	Create(userAccount models.UserAccounts) error
	UsernameExist(username string) bool

	FindBy(field string, valueField string, userAccount models.UserAccounts) (models.UserAccounts, error)
}

type user_repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user_repository {
	return &user_repository{db}
}

func (r *user_repository) UsernameExist(username string) bool {
	var user models.UserAccounts
	err := r.db.First(&user, "username = ?", username).Error
	return err == nil
}

func (r *user_repository) FindBy(field string, valueField string, userAccount models.UserAccounts) (models.UserAccounts, error) {

	strField := fmt.Sprintf("%s = ?", field)
	err := r.db.Where(strField, valueField).First(&userAccount).Error
	return userAccount, err
}

func (r *user_repository) GetAll() ([]models.UserAccounts, error) {
	var user_account []models.UserAccounts
	err := r.db.Find(&user_account).Error
	return user_account, err
}

func (r *user_repository) GetByID(id any) (models.UserAccounts, error) {
	var user_account models.UserAccounts
	err := r.db.Find(&user_account, id).Error
	return user_account, err
}

func (r *user_repository) Create(userAccount models.UserAccounts) error {
	err := r.db.Create(&userAccount).Error
	return err
}
