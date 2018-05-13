package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
)

func UserCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	user := models.User{}
	nickname := ctx.UserValue("nickname").(string)
	user.UnmarshalJSON(ctx.PostBody())
	user.NickName = nickname

	resp, err := user.UserCreate()

	switch err {
	case nil:
		ctx.SetStatusCode(201)
		buf, _ := user.MarshalJSON()
		ctx.Write(buf)
	case errors.UserIsExist:
		ctx.SetStatusCode(409)
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func UserProfile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	nickname := ctx.UserValue("nickname").(string)

	result := models.User{}
	err := result.UserProfile(nickname)

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := result.MarshalJSON()
		ctx.Write(buf)
	case errors.UserNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}

func UserUpdateProfile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	nickname := ctx.UserValue("nickname").(string)
	user := models.User{}
	user.UnmarshalJSON(ctx.PostBody())
	user.NickName = nickname

	err := user.UpdateUserProfile()

	switch err {
	case nil:
		ctx.SetStatusCode(200)
		buf, _ := user.MarshalJSON()
		ctx.Write(buf)
	case errors.UserUpdateConflict:
		ctx.SetStatusCode(409)
		ctx.Write([]byte(err.Error()))
	case errors.UserNotFound:
		ctx.SetStatusCode(404)
		ctx.Write([]byte(err.Error()))
	}
}
