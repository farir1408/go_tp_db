package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/models"
)

func StatusDataBase(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	status := models.Status{}
	status.StatusDataBase()
	ctx.SetStatusCode(200)
	buf, _ := status.MarshalJSON()
	ctx.Write(buf)
}

func ClearDataBase(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	models.ClearDataBase()
	ctx.SetStatusCode(200)
}