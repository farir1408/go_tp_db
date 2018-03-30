package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	//"log"
	"strings"
)

func ForumCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	forum := models.Forum{}
	forum.UnmarshalJSON(ctx.PostBody())

	resp, err := forum.ForumCreate()

	if err == nil {
		ctx.SetStatusCode(201)
		//log.Println("this block is completed AllIsOkey Forum")
		buf, _ := forum.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserNotFound {
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		//log.Println("this block is completed UserNotFound Forum")
		ctx.Write(resErr)
	}

	if err == errors.ForumIsExist {
		ctx.SetStatusCode(409)
		//log.Println("this block is completed ForumIsExist")
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func ForumDetails(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	slug := ctx.UserValue("slug").(string)
	strings.ToLower(slug)
	forum := models.Forum{}
	err := forum.ForumDetails(slug)

	if err == nil {
		ctx.SetStatusCode(200)
		//log.Println("this block is completed AllIsOkey Forum")
		buf, _ := forum.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserNotFound {
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		//log.Println("this block is completed UserNotFound in GetForum")
		ctx.Write(resErr)
	}
}

func ForumThreadCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	thread := models.Thread{}
	thread.UnmarshalJSON(ctx.PostBody())

	slug := ctx.UserValue("slug").(string)
	strings.ToLower(slug)

	//log.Println(slug)
	thread.ForumId = slug

	resp, err := thread.ForumThreadCreate()

	//log.Println(err)

	if err == nil {
		ctx.SetStatusCode(201)
		//log.Println("this block is completed AllIsOkey Thread")
		buf, _ := thread.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.ForumNotFound {
		ctx.SetStatusCode(404)
		resErr, _ := models.Error{err.Error()}.MarshalJSON()
		//log.Println("this block is completed UserNotFound Thread")
		ctx.Write(resErr)
	}

	if err == errors.ThreadIsExist {
		ctx.SetStatusCode(409)
		//log.Println("this block is completed ThreadIsExist")
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}
