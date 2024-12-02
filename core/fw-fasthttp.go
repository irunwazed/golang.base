package core

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/valyala/fasthttp"
)

func (c *Context) Query(key string) string {
	return string(c.ctx.QueryArgs().Peek(key))
}

func (c *Context) JSONBody(key string) interface{} {
	body := c.ctx.Request.Body()
	if len(body) == 0 {
		return nil
	}

	var jsonBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &jsonBody); err != nil {
		return nil
	}
	return jsonBody[key]
}

func (c *Context) Body(key string) interface{} {
	body := c.ctx.Request.Body()
	if len(body) == 0 {
		return nil
	}
	pairs := strings.Split(string(body), "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 && kv[0] == key {
			return kv[1]
		}
	}
	return nil
}

func (c *Context) Form(key string) string {
	return string(c.ctx.FormValue(key))
}

func (c *Context) File(key string) *multipart.FileHeader {
	file, err := c.ctx.FormFile(key)
	if err != nil {
		return nil
	}
	return file
}

func (c *Context) SetContentType(contentType string) {
	c.ctx.SetContentType(contentType)
}

func (c *Context) SetStatusCode(statusCode int) {
	c.ctx.SetStatusCode(statusCode)
}

func (ctx *Context) WriteString(s string) (int, error) {
	ctx.ctx.Response.AppendBodyString(s)
	return len(s), nil
}

func (ctx *Context) Path() []byte {
	return ctx.ctx.URI().Path()
}

func (ctx *Context) Method() []byte {
	return ctx.ctx.Method()
}

func (ctx *Context) LocalIP() []byte {
	return ctx.ctx.LocalIP()
}

func CoreCtx(ctx *fasthttp.RequestCtx) *Context {
	return &Context{ctx: ctx}
}

func (r *Router) ServeHTTP(ctx *fasthttp.RequestCtx) {
	c := CoreCtx(ctx)
	method := c.ctx.Method()
	path := c.ctx.Path()

	if handler, ok := r.routes[string(method)][string(path)]; ok {

		middlewares := r.middlewares[string(path)]

		next := func() error {
			err := handler(c)
			if err != nil {
				ResponseError(c, "ERROR Handler")
				return fmt.Errorf("ERROR Handler")
			}
			return nil
		}

		if len(middlewares) == 0 {
			next()
			return

		}

		isNext := false
		for _, middleware := range middlewares {
			isNext = false
			middleware(c, func() error {
				isNext = true
				return nil
			})
			if !isNext {
				break
			}
		}

		if isNext {
			next()
			return
		}
		ResponseUnAuthor(c, "Invalid User")
		return
	} else {
		ResponseNotFound(c, "Page Not Found")
		return
	}
}

func RunServer(router *Router, port int) {
	server := &fasthttp.Server{
		Handler: router.ServeHTTP,
	}

	fmt.Printf("Server running at http://localhost:%v\n", port)
	if err := server.ListenAndServe(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
