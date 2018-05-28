package server

import (
	"log"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func ProfileHandlers() *gin.Engine {
	e := gin.Default()
    e.GET("/", func(context *gin.Context) {
		context.Redirect(301, context.Request.RequestURI+"/pprof")
	})

	g := e.Group("/internal/debug/pprof")
	{
		g.GET("/", gin.WrapF(pprof.Index))
		g.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		g.GET("/profile", gin.WrapF(pprof.Profile))
		g.GET("/symbol", gin.WrapF(pprof.Symbol))
		g.GET("/trace", gin.WrapF(pprof.Trace))
		g.GET("/block", gin.WrapH(pprof.Handler("block")))
		g.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		g.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		g.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}

	return e
}

func LoadProfileServer(port string) {
	go func() {
		log.Fatalln(ProfileHandlers().Run(port))
	}()
}
