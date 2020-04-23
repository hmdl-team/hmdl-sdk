package request

type LoginRequest struct {
	Name string `json:"name" validate:"required"`
	Pass string `json:"pass" validate:"required"`
}

type RefreshTokenRequest struct {
	Token string `json:"token,omitempty" validate:"required"`
}
