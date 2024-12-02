package core

type Router struct {
	routes      map[string]map[string]func(c *Context) error
	middlewares map[string][]func(c *Context, next Next) error
}

func NewRouter() *Router {
	return &Router{
		routes:      make(map[string]map[string]func(c *Context) error),
		middlewares: make(map[string][]func(c *Context, next Next) error),
	}
}

func (r *Router) Handle(method, path string, handler func(c *Context) error) {
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = make(map[string]func(c *Context) error)
	}
	r.routes[method][path] = handler
}

func (r *Router) GET(path string, handler func(c *Context) error) {
	r.Handle("GET", path, handler)
}

func (r *Router) POST(path string, handler func(c *Context) error) {
	r.Handle("POST", path, handler)
}

func (r *Router) PUT(path string, handler func(c *Context) error) {
	r.Handle("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler func(c *Context) error) {
	r.Handle("DELETE", path, handler)
}

func (r *Router) Use(path string, middleware func(c *Context, next Next) error) {
	r.middlewares[path] = append(r.middlewares[path], middleware)
}
