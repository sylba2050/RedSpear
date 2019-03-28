package main

import (
    "./handler"
    "./handler/article"
    "./struct/DB"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.CORS())

    db, err := gorm.Open("sqlite3", "DB/main.sqlite3")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&DB.Auth{})
    db.AutoMigrate(&DB.Article{})
    db.AutoMigrate(&DB.ArticleStatus{})
    db.AutoMigrate(&DB.Like{})
    db.AutoMigrate(&DB.Stock{})
    db.AutoMigrate(&DB.Comment{})
    db.AutoMigrate(&DB.CpForArticle{})
    db.AutoMigrate(&DB.User{})
    db.AutoMigrate(&DB.UserFollowTag{})
    db.AutoMigrate(&DB.FF{})
    db.AutoMigrate(&DB.CpForUser{})

    e.Static("/img", "img")
    e.Static("/js", "js")
    e.Static("/css", "css")

    e.File("/", "html/index.html")
    e.File("/write", "html/write.html")

    e.POST("/article/post", article.POST(db))
    e.GET("/articles", article.GET(db))
    e.GET("/articles/:userid", article.GetByUser(db))

    e.POST("/md", markdown.Default)

    e.Start(":8080")
}
