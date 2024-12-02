package middlewares

import (
	"simplefast/core"
)

func LoggingMiddleware(ctx *core.Context, next core.Next) error {
	core.Logger("middleware1", string(ctx.Path()))
	return next()
}

func LoggingMiddlewareFaile(ctx *core.Context, next core.Next) error {
	core.Logger("middleware2", string(ctx.Path()))
	return core.ResponseUnAuthor(ctx, "Invalid User")

}
