package requests

type UserStoreRequest struct {
	FirstName       string `json:"first_name" binding:"required,alpha"`
	LastName        string `json:"last_name" binding:"required,alpha"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name" binding:"omitempty,alpha"`
	LastName  string `json:"last_name" binding:"omitempty,alpha"`
	Email     string `json:"email" binding:"omitempty,email"`
}
