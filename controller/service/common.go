package service

type Common struct {
	UserAPI *User `inject:""`
	RoleAPI *Role `inject:""`
	MenuAPI *Menu `inject:""`
}
