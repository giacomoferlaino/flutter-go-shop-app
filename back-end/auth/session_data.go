package auth

// SessionData contains the information needed to send authenticated requests
type SessionData struct {
	IDToken   string `json:"idToken"`
	ExpiresIn uint64 `json:"expiresIn"`
}
