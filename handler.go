package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

var pathHandler = map[string]func(ctx *fasthttp.RequestCtx){
	"/load": func(ctx *fasthttp.RequestCtx) {
		if !ctx.IsPost() {
			ctx.WriteString(fmt.Sprintf("Unexpected request method: %s!", string(ctx.Request.Header.Method())))
		}
		updateDictionary(ctx)
	},
	"/get": func(ctx *fasthttp.RequestCtx) {
		if !ctx.IsGet() {
			ctx.WriteString(fmt.Sprintf("Unexpected request method: %s!", string(ctx.Request.Header.Method())))
		}
		getAnagrams(ctx)
	},
}

// Тут не стал тянуть какую-то либу-роутер, написал свой(в проде юзал бы готовые и оптимизированные либы).
func handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}
	if fn, ok := pathHandler[path]; ok {
		fn(ctx)
	} else {
		ctx.WriteString("Unsupported path!")
	}
}
