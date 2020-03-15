package request

type ReqSignIn struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
