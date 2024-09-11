package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"go-oauth2/config"
	"go-oauth2/controller"
	"go-oauth2/middleware"
	"go-oauth2/router"
	"go-oauth2/usecase"
)

func main() {
	config, err := config.LoadEnvVars()
	if err != nil {
		log.Fatal("Error while reading config file")
	}
	// create the different layers
	u := usecase.NewUsecase()
	c := controller.NewController(u)
	m := middleware.NewMiddleware()
	r := router.NewRouter(c, m)

	s := &http.Server{
		Addr: fmt.Sprintf(":%s", config.Port),
		Handler: r,
	}
	
	log.Printf("Server running on port: %s \n", config.Port)
	err = s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Fatal error while the api was starting: %s", err.Error())
	}

	
}
