package router

import "github.com/gin-gonic/gin"

type Controller interface {
	Health(g *gin.Context)
}

// Middleware contains the middleware methods
type Middleware interface {
	CORS() gin.HandlerFunc
}

func NewRouter(c Controller, middleware Middleware) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())

	root := r.Group("/")
	root.GET("/health", c.Health)

	return r
}