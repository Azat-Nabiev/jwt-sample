package dto

type UserResponse struct {
	Name  string `json:"name" binding:"required"`
	Login string `json:"password" binding:"required"`
}
