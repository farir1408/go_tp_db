package handlers


import (
	"github.com/valyala/fasthttp"
	"log"
	"go_tp_db/models"
	"go_tp_db/errors"
)

func ThreadDetails(ctx *fasthttp.RequestCtx) {
	log.Println("Thread Details")
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	thread := models.Thread{}

	err := thread.ThreadDetails(slug)

	if err == nil {
		ctx.SetStatusCode(200)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.ThreadNotFound {
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.Write(resErr)
	}
}