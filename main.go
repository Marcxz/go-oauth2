package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"go-oauth2/controller"
	"go-oauth2/middleware"
	"go-oauth2/router"
	"go-oauth2/usecase"
	
)

func main() {
	// create the different layers
	u := usecase.NewUsecase()
	c := controller.NewController(u)
	m := middleware.NewMiddleware()
	r := router.NewRouter(c, m)

	s := &http.Server{
		Addr: ":3030",
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Fatal error while the api was starting: %s", err.Error())
	}
}
