package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	n := handler
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}
	return n
}