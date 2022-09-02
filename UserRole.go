package WebControlListModule

type UserRoleValue uint

type UserRole struct {
	Name     string
	Parent   *UserRole
	Children []*UserRole
	Value    UserRoleValue
}
