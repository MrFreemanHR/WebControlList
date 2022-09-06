package WebControlListModule

import "strings"

type OperationType uint8

const (
	AllowRule  = OperationType(1)
	DeniedRule = OperationType(2)
)

type RouterRule struct {
	Operation     OperationType
	UserStatement UserStatement
}

func (r *RouterRule) HumanReadable() string {
	var result []string
	switch r.Operation {
	case AllowRule:
		result = append(result, "ALLOW")
	case DeniedRule:
		result = append(result, "DENY")
	}
	result = append(result, r.UserStatement.humanReadable())
	return strings.Join(result, " ")
}
