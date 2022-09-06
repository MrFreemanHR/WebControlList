package WebControlListModule

import (
	"errors"
	"fmt"
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

/**
Main functions
*/

func (w *WebControlList) AddNewRouteRule(Rule RouterRule, Route string) {
	// Splice string by '/' delimiter
	splittedPath := strings.Split(Route, "/")
	node := w.findNodeByPath(splittedPath, w.routeNodes)
	node.addRouterRule(Rule)
}

func (w *WebControlList) AddNewUserRole(Name string, Value UserRoleValue, ParentName string) error {
	var parentRole *UserRole
	Name = strings.ToLower(Name)
	ParentName = strings.ToLower(ParentName)
	if ParentName != "" {
		result := w.findRoleByName(ParentName, nil)
		if result == nil {
			return errors.New("parent role not found")
		}
		parentRole = result
	} else {
		parentRole = w.rolesTree
	}
	// We find new role from tree root because roles are unique
	result := w.findRoleByName(Name, w.rolesTree)
	if result != nil {
		return errors.New("not unique new role")
	}
	var newRole = &UserRole{
		Name:   Name,
		Parent: parentRole,
		Value:  Value,
	}
	parentRole.Children = append(parentRole.Children, newRole)
	return nil
}

func (w *WebControlList) GetUserRole(Name string) *UserRole {
	Name = strings.ToLower(Name)
	return w.findRoleByName(Name, w.rolesTree)
}

/**
Debug functions
*/

func (w *WebControlList) PrintAllRouteNodes(Parent *RouteNode, callLevel ...int) {
	if Parent == nil {
		Parent = w.routeNodes
	}
	var asterLevel = 1
	if len(callLevel) != 0 {
		asterLevel = callLevel[0]
	}
	var asterString = ""
	asterString = asterString + strings.Repeat("*", asterLevel)
	fmt.Printf("- %s: /%s\n", asterString, Parent.Name)
	if len(Parent.routerRules) > 0 {
		fmt.Printf("- %s: |\n", asterString)
		var maxLen = 0
		for _, rule := range Parent.routerRules {
			humanReadable := rule.HumanReadable()
			if len(humanReadable) > maxLen {
				maxLen = len(humanReadable)
			}
			fmt.Printf("- %s: |- %s\n", asterString, humanReadable)
		}
		fmt.Printf("- %s: |%s\n", asterString, strings.Repeat("_", maxLen+2))
	}
	for _, child := range Parent.Children {
		w.PrintAllRouteNodes(child, asterLevel+1)
	}
}

/**
Internal functions
*/

func (w *WebControlList) findNodeByPath(path []string, rootNode *RouteNode) *RouteNode {
	name := path[0]
	// If this is the desired node return this
	if name == rootNode.Name && len(path) == 1 {
		return rootNode
	}
	if name == rootNode.Name {
		newPath := path[1:]
		return w.findNodeByPath(newPath, rootNode)
	}
	// If this is the end of path - create the last one node
	if len(path) == 1 {
		newRootNode := &RouteNode{
			Name:   name,
			Parent: rootNode,
		}
		rootNode.Children = append(rootNode.Children, newRootNode)
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

func (w *WebControlList) findRoleByName(Name string, Parent *UserRole) *UserRole {
	Name = strings.ToLower(Name)
	if Parent == nil {
		Parent = w.rolesTree
	}
	if len(Parent.Children) != 0 {
		for _, child := range Parent.Children {
			if child.Name == Name {
				return child
			} else {
				result := w.findRoleByName(Name, child)
				if result != nil {
					return result
				}
			}
		}
	}
	return nil
}
