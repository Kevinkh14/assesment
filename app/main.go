package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

func main() {

  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/file", hello)

  e.Logger.Fatal(e.Start(":8081"))
}

func hello(c echo.Context) error {
  return c.yaml(http.StatusOK, "objects1")
}