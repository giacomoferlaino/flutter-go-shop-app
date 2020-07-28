package auth

// LoginData is the login request data model
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
