package authenticate

type UserAuthRequest struct {
	Name  string `json:"name" binding:"required"`
	Login string `json:"login" binding:"required"`
}
