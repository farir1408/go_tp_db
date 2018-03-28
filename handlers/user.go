package handlers

import (
	"github.com/valyala/fasthttp"
	"go_tp_db/errors"
	"go_tp_db/models"
	"log"
)

func UserCreate(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	user := models.User{}
	nickname := ctx.UserValue("nickname").(string)
	user.UnmarshalJSON(ctx.PostBody())
	user.NickName = nickname

	resp, err := user.UserCreate()

	if err == nil {
		ctx.SetStatusCode(201)
		log.Println("user created successfull")
		buf, _ := user.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserIsExist {
		ctx.SetStatusCode(409)
		log.Println("user was created earlier")
		buf, _ := resp.MarshalJSON()
		ctx.Write(buf)
	}
}

func UserProfile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	nickname := ctx.UserValue("nickname").(string)

	result := models.User{}
	err := result.UserProfile(nickname)

	if err == nil {
		ctx.SetStatusCode(200)
		buf, _ := result.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserNotFound {
		ctx.SetStatusCode(404)
		ctx.Write([]byte(`{"message":"Can't find user"}`))
	}
}

func UserUpdateProfile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	nickname := ctx.UserValue("nickname").(string)
	user := models.User{}
	user.UnmarshalJSON(ctx.PostBody())
	user.NickName = nickname
	log.Println(user.NickName)

	err := user.UpdateUserProfile()

	if err == nil {
		log.Println("All is OK")
		ctx.SetStatusCode(200)
		buf, _ := user.MarshalJSON()
		ctx.Write(buf)
	}

	if err == errors.UserUpdateConflict {
		log.Println("ERROR is Exist")
		ctx.SetStatusCode(409)
		ctx.Write([]byte(`{"message":"Conflict with existing users"}`))
	}

	if err == errors.UserNotFound {
		log.Println("ERROR is Exist")
		ctx.SetStatusCode(404)
		ctx.Write([]byte(`{"message":"Can't find user"}`))
	}
}
