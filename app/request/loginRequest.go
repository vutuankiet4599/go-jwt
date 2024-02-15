package request

type LoginRequest struct {
	Email    string `json:"email" format:"email" binding:"required,email"`
	Password string `json:"password" format:"password" binding:"required"`
}
