package main

import (
	"log"
	"net/http"
	"user_service/internal/api"
	"user_service/internal/db"
	api_desc "user_service/internal/pkg/api"
	"user_service/internal/service"
)

func main() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := db.NewUserRepository(conn)

	svc := service.NewUserService(repo)

	apiHandler := api.NewServer(svc)
	apiAuth := api.NewSecurity(svc)

	server, err := api_desc.NewServer(apiHandler, apiAuth)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":8081", server); err != nil {
		log.Fatal(err.Error())
	}
}
