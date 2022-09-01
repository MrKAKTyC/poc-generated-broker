package main

import (
	"broker/pkg/server"
	"broker/pkg/service"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	router := echo.New()
	broker := service.NewServiceBroker()
	server.RegisterHandlers(router, broker)
	router.Use(middleware.Logger())
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println(err)
	}
}
