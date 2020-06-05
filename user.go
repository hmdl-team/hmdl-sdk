package sdk

type IUser interface {
	GetUid() int
	GetRole() string
	GetDisplayName() string
}
