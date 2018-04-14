package main

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/config"
	"go_tp_db/router"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func accessLogMiddleware(router fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		log.Println("Access Log Middleware: ", ctx.URI())
		start := time.Now()
		router(ctx)
		log.Printf("[%s] %s, %s\n", string(ctx.Method()), ctx.URI(), time.Since(start))
	})
}

func panicMiddleware(router fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic Middleware", ctx.URI())
				log.Println("recover", err)
				ctx.Error("Internal Server Error", 500)
			}
		}()
		router(ctx)
	})
}

func main() {
	config.InitDB()

	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-syscallChan
		log.Println("Shutdown")
		config.Disconnect()
		os.Exit(0)
	}()

	router := router.InitRouter()
	accessLog := accessLogMiddleware(router.Handler)

	log.Fatal(fasthttp.ListenAndServe(":5000", panicMiddleware(accessLog)))
}
