package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"

  "net/http"
)

func getName(c echo.Context) error {
  name :=c.param("name")
  return c.JSON(http.StatusOK, name)
}

func main() {

  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/objects/object1/:name", getName)

  e.Logger.Fatal(e.Start(":8081"))
}

