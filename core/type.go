package core

import "github.com/valyala/fasthttp"

type Context struct {
	ctx *fasthttp.RequestCtx
}

type Next func() error
