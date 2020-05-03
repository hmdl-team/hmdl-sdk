package request

type LoginRequest struct {
	Name string `json:"name" validate:"required"`
	Pass string `json:"pass" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token,omitempty" validate:"required"`
}
