package WebControlListModule

import "strings"

type UserStatement struct {
	RoleStatement UserRole
	Strict        bool
}

func (u *UserStatement) humanReadable() string {
	if u.Strict {
		return u.RoleStatement.Name
	} else {
		return strings.Join(u.getAllWithParents(&u.RoleStatement, []string{}), " ")
	}
}

func (u *UserStatement) getAllWithParents(Role *UserRole, previous []string) []string {
	if Role.Parent != nil {
		previous = append(previous, Role.Name)
		return u.getAllWithParents(Role.Parent, previous)
	} else {
		if Role.Name == "*" {
			return []string{"*"}
		}
		return previous
	}
}
