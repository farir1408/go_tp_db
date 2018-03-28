package main

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/config"
	"go_tp_db/router"
)

func main() {
	config.InitDB()

	fasthttp.ListenAndServe(":5000", nil)

}
