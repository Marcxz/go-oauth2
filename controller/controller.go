package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Usecase interface {
}

type Controller struct {
	u Usecase
}

func NewController(u Usecase) *Controller {
	return &Controller{u}
}

func (c *Controller) Health(g *gin.Context) {
	type Response struct {
		Message string
	}
	r := Response{
		Message: "The health endpoint is alive" ,
	}
	g.JSON(http.StatusOK, r)
	log.Println("Health endpoint has been executed")
}
