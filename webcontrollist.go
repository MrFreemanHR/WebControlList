package WebControlListModule

import (
	"log"
	"strings"
)

type WebControlList struct {
	rolesTree  *UserRole
	routeNodes *RouteNode
}

func New() *WebControlList {
	allRolesDefault := UserRole{
		Name:  "*",
		Value: UserRoleValue(0),
	}
	rootRoute := RouteNode{
		Name: "",
	}
	var wcl = WebControlList{
		rolesTree:  &allRolesDefault,
		routeNodes: &rootRoute,
	}
	return &wcl
}

func (w *WebControlList) AddNewRouteRule(Rule RouterRule, Route string) {
	// Splice string by '/' delimiter
	splittedPath := strings.Split(Route, "/")
	w.findNodeByPath(splittedPath, w.routeNodes)
	log.Printf("%#+v", w.routeNodes)
}

func (w *WebControlList) findNodeByPath(path []string, rootNode *RouteNode) *RouteNode {
	name := path[0]
	// If this is the desired node return this
	if name == rootNode.Name && len(path) == 1 {
		return rootNode
	}
	// If this is the end of path - create the last one node
	if len(path) == 1 {
		newRootNode := &RouteNode{
			Name:   name,
			Parent: rootNode,
		}
		rootNode.Children = append(rootNode.Children, newRootNode)
		log.Printf("Create last one Node: %#+v", newRootNode)
		return rootNode
	}
	// If no child nodes - create new with path's part and continue to search
	if len(rootNode.Children) == 0 {
		newRootNode := &RouteNode{
			Name:   name,
			Parent: rootNode,
		}
		newPath := path[1:]
		rootNode.Children = append(rootNode.Children, newRootNode)
		log.Printf("Created child: %#+v", newRootNode)
		return w.findNodeByPath(newPath, newRootNode)
	}
	// If there are child nodes - search in it
	for _, child := range rootNode.Children {
		if name == child.Name {
			newPath := path[1:]
			return w.findNodeByPath(newPath, child)
		}
	}
	// If no one from child nodes is valid - create new and continue search
	newRootNode := &RouteNode{
		Name:   name,
		Parent: rootNode,
	}
	newPath := path[1:]
	rootNode.Children = append(rootNode.Children, newRootNode)
	return w.findNodeByPath(newPath, newRootNode)
}
