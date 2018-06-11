package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go_tp_db/handlers"
	"net/http/pprof"
	"strings"
)

var (
	Cmdline = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	Profile = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	Symbol  = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	Trace   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	Index   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
)

func PprofHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("text/html")

	if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/cmdline") {
		Cmdline(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/profile") {
		Profile(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/symbol") {
		Symbol(ctx)
	} else if strings.HasPrefix(string(ctx.Path()), "/debug/pprof/trace") {
		Trace(ctx)
	} else {
		Index(ctx)
	}
}

func InitRouter() *fasthttprouter.Router {
	r := fasthttprouter.New()

	r.GET("/debug/pprof/:match", PprofHandler)
	r.POST("/api/forum/:slug", handlers.ForumCreate)
	r.POST("/api/forum/:slug/create", handlers.ForumThreadCreate)
	r.GET("/api/forum/:slug/details", handlers.ForumDetails)
	r.GET("/api/forum/:slug/threads", handlers.GetThreads)
	r.GET("/api/forum/:slug/users", handlers.GetUsers)

	r.GET("/api/post/:id/details", handlers.PostDetails)
	r.POST("/api/post/:id/details", handlers.PostUpdate)

	r.POST("/api/service/clear", handlers.ClearDataBase)
	r.GET("/api/service/status", handlers.StatusDataBase)

	r.POST("/api/thread/:slug_or_id/create", handlers.PostsCreate)
	r.GET("/api/thread/:slug_or_id/details", handlers.ThreadDetails)
	r.POST("/api/thread/:slug_or_id/details", handlers.ThreadUpdateDetails)
	r.GET("/api/thread/:slug_or_id/posts", handlers.ThreadPosts)
	r.POST("/api/thread/:slug_or_id/vote", handlers.Vote)

	r.POST("/api/user/:nickname/create", handlers.UserCreate)
	r.GET("/api/user/:nickname/profile", handlers.UserProfile)
	r.POST("/api/user/:nickname/profile", handlers.UserUpdateProfile)

	return r
}
