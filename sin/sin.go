package sin

import (
	"log"
	"net/http"
)

// 
type HandlerFunc func(*Context)

//
type Engine struct{
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

/*
	Reguest route	

	request_type: HTTP Request, such as "GET", "POST"
	route_address: Route Address, such as "/hello"
	handler: handler method
*/
func (e *Engine) addRoute(request_type string, route_address string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", request_type, route_address)
	e.router.addRoute(request_type, route_address,handler)
}

// GET request
func (e  *Engine) GET(route_address string, handler HandlerFunc) {
	e.addRoute("GET", route_address, handler)
}

// POST request
func  (e *Engine) POST(route_address string, handler HandlerFunc) {
	e.addRoute("POST",route_address,handler)
}

//  Start a http server

func (e *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr, e)
}

// implement http.Handler inferface
func (e *Engine)ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
