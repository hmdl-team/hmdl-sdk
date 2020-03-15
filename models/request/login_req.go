package request

type LoginRequest struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type RefreshTokenRequest struct {
	Token string `json:"token,omitempty"`
}
