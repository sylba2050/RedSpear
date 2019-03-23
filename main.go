package main

import (
    "./handler"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())

    e.Static("/img", "img")

    e.File("/", "html/index.html")
    e.File("/write", "html/write.html")

    e.File("/js/index.js", "js/index.js")
    e.File("/js/write.js", "js/write.js")

    e.File("/css/index.css", "css/index.css")
    e.File("/css/write.css", "css/write.css")
    e.File("/css/code.css", "css/code.css")
    e.File("/css/github.css", "css/github.css")

    e.POST("api/md", markdown.Default)

    e.Start(":8080")
}
