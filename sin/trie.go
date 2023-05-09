package sin

import "strings"


type node struct {
	pattern string // route, such as /path/:cty
	part string // a part of route, such as :cty
	children []*node
	isWild bool // fuzzy match
}

// find the first match node
func (n *node) matchChild(part string) *node  {
	for _, child := range n.children {  // search in n.children
		if child.part == part || child.isWild { // if match or support fuzzy match
			return  child
		}
	}
	return nil
}

// find all match nodes
func (n* node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part ==  part || child.isWild {
			nodes = append(nodes,child)
		}
	}
	return nodes
}

// insert a pattern into  roots
/*
	pattern: /users/cty/note
	parts: {users, cty, note}
	height: insert from layer 0
*/
func (n *node) insert(pattern string, parts []string, height int)  {
	if len(parts) == height { // insert finish (all of the parts are inserted)
		n.pattern = pattern
		return
	}

	part := parts[height] // the part will be insert (into current node's children)right now
	child := n.matchChild(part) // search in current node's children
	if child == nil { // if no match, create a new one
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} //if pattern start with ':' or '*', it will be a fuzzy match node
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height + 1) // insert the next part in parts
}

// search 
func (n *node) search(parts []string, height int) *node {
	// if all parts are searched or current node isWild then search finish
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {// no match
			return nil
		}
		return n // find
	}

	part := parts[height]
	children := n.matchChildren(part)

	for  _, child := range children {
		result := child.search(parts, height + 1)
		if result !=  nil  { //  find
			return result
		}
	}
	return nil // not find
}
