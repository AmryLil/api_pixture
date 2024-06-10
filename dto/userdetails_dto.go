package dto

import "time"

type UserDetailsRequest struct {
	Id          string    `json:"id" binding:"required"`
	Telp        string    `json:"telp" `
	Gender      string    `json:"gender" `
	Location    string    `json:"location" `
	Email       string    `json:"email" `
	Description string    `json:"description" `
	CreatedAt   time.Time `json:"created_at" `
}
type UserDetailsResponse struct {
	Telp        string    `json:"telp" `
	Gender      string    `json:"gender" `
	Location    string    `json:"location" `
	Email       string    `json:"email" `
	Description string    `json:"description" `
	CreatedAt   time.Time `json:"created_at" `
}
