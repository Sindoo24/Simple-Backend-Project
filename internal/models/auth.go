package models


type SignupRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Dob      string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type SignupResponse struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	User    struct {
		ID    int32  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	} `json:"user"`
}

type AuthUser struct {
	ID   int32  `json:"id"`
	Role string `json:"role"`
}
