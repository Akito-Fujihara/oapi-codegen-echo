package main

import (
	"net/http"

	api "github.com/Akito-Fujihara/oapi-codegen-echo/oapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiV1Group := e.Group("")

	server := &Server{}

	api.RegisterHandlers(apiV1Group, server)

	e.Logger.Fatal(e.Start(":8086"))
}

type Server struct{}

func (s *Server) GetApiV1Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}
