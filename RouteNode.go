package WebControlListModule

type RouteNode struct {
	Name        string
	Parent      *RouteNode
	Children    []*RouteNode
	routerRules []*RouterRule
}

func (r *RouteNode) addRouterRule(rule RouterRule) {
	r.routerRules = append(r.routerRules, &rule)
}
