package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	"log"
)

func ForumCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	forum := models.Forum{}
	resErr := models.Error{}
	forum.UnmarshalJSON(ctx.PostBody())

	resp, err := forum.ForumCreate()

	if err == nil {
		ctx.SetStatusCode(201)
		log.Println("this block is completed AllIsOkey Forum")
		buf, _ := forum.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserNotFound {
		ctx.SetStatusCode(404)
		log.Println("this block is completed UserNotFound Forum")
		ctx.Write(resErr.ErrorMsgJSON(err.Error()))
	}

	if err == errors.ForumIsExist {
		ctx.SetStatusCode(409)
		log.Println("this block is completed ForumIsExist")
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

