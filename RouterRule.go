package WebControlListModule

type OperationType uint8

const (
	AllowRule  = OperationType(1)
	DeniedRule = OperationType(2)
)

type RouterRule struct {
	Operation     OperationType
	UserStatement UserStatement
}
