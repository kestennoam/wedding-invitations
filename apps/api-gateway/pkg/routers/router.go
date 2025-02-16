package routers

import "github.com/gin-gonic/gin"

type Router interface {
	setupRoutes(*gin.Engine)
}
type AppRouter interface {
	Run(addr string)
}
type router struct {
	engine *gin.Engine
}

func NewRouter() AppRouter {
	engine := gin.Default()
	r := &router{engine: engine}
	// weddingsRouter
	wr := &WeddingsRouter{}
	r.setupRoutes(wr)

	return r
}

func (r *router) setupRoutes(router Router) {
	router.setupRoutes(r.engine)
}

func (r *router) Run(addr string) {
	r.engine.Run(addr)
}
