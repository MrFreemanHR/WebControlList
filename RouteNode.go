package WebControlListModule

type RouteNode struct {
	Name        string
	Parent      *RouteNode
	Children    []*RouteNode
	routerRules []*RouterRule
}
