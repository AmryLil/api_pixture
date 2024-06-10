package user

import "time"

type UserAccounts struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Fullname  string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}

type UserDetails struct {
	Id          string `gorm:"primaryKey`
	Telp        string `gorm:"type:varchar(255)"`
	Gender      string `gorm:"type:varchar(255)"`
	Location    string `gorm:"type:varchar(255)"`
	Email       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
}
