package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
)

func ThreadDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	thread := models.Thread{}

	err := thread.ThreadDetails(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	case errors.ThreadNotFound:
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.Write(resErr)
	}
}

func ThreadUpdateDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug_or_id").(string)
	thread := models.ThreadUpdate{}
	thread.UnmarshalJSON(ctx.PostBody())

	result, err := thread.ThreadUpdate(slug)

	if err == nil {
		ctx.SetStatusCode(200)
		buf, _ := result.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.ThreadNotFound {
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.Write(resErr)
	}
}

func Vote(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	vote := models.Vote{}
	vote.UnmarshalJSON(ctx.PostBody())
	slug := ctx.UserValue("slug_or_id").(string)

	err := vote.Vote(slug)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		thread := models.Thread{}
		thread.ThreadDetails(slug)
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	case errors.ThreadNotFound:
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		ctx.Write(resErr)
	}
}
