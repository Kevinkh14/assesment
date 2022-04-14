package main

import (
  "log"
  "io/ioutil"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "gopkg.in/yaml.v2"
  "net/http"
)

type Config struct {
  Name string `json: "name"`
  Description string `json:"description"`
  Attributes []string `jason:"attributes"`
}

func getName(c echo.Context) error {
  // parmenter
  objName :=c.Param("objName")

  //added pararmeter name to search for file
  var file = "objects/" + objName + ".yaml"  

  //search for file
  data, err := ioutil.ReadFile(file)
  if err != nil {
    log.Fatalf("error :%v",err)
  }

  //bind yaml data
  var config Config
  err = yaml.Unmarshal([]byte(data), &config)
  if err != nil{
    log.Fatalf("error: %v", err)
  }

  //put yaml data into result
  result := Config{
    Name: config.Name,
    Description:config.Description,
    Attributes:config.Attributes,
  }

  // return json 
  return c.JSON(http.StatusOK, result)
}

func main() {

  e := echo.New()
  e.Use(middleware.CORS())

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/:objName", getName)

  e.Logger.Fatal(e.Start(":8081"))
}

