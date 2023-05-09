package sin

import (
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node //trie tree, roots["GET"] is the trie tree root of "GET" req
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{roots: make(map[string]*node),
	handlers: make(map[string]HandlerFunc),}
}

func  parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/") // resolve pattern into  parts

	parts := make([]string, 0)
	for _, item := range vs  {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// add a new router
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok { // if roots[method] is not exist (for example,  method = 'GET' and there is no 'GET' req is registed so far), there create it
		r.roots[method] =  &node{}
	}
	r.roots[method].insert(pattern,parts, 0)// insert pattern into trie tree
	r.handlers[key] = handler
}

func  (r *router)getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] ==  ':' {
				params[part[1:]] = searchParts[index]
			}
			if  part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:],"/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}	
}
