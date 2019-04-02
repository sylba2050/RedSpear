package main

import (
    "./handler/markdown"
    "./handler/article"
    "./handler/user"
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

    e.POST("/article", article.Post(db))
    e.PUT("/article", article.Update(db))
    e.DELETE("/article/:articleid", article.Delete(db))

    e.POST("/article/:articleid/stock", article.Stock(db))
    e.DELETE("/article/:articleid/stock", article.UnStock(db))

    e.POST("/article/:articleid/like", article.Like(db))
    e.DELETE("/article/:articleid/like", article.UnLike(db))

    e.POST("/article/:articleid/comment", article.Comment(db))
    e.PUT("/article/:articleid/comment/:commentid", article.UpdateComment(db))
    e.DELETE("/article/:articleid/comment/:commentid", article.DeleteComment(db))

    e.GET("/articles", article.Get(db))
    e.GET("/articles/newer", article.GetNewer(db))
    e.GET("/articles/visited", article.GetVisited(db))
    e.GET("/articles/populer/day", article.GetPopulerPerDay(db))
    e.GET("/articles/populer/week", article.GetPopulerPerWeek(db))
    e.GET("/articles/populer/month", article.GetPopulerPerMonth(db))
    e.GET("/articles/:userid", article.GetByUser(db))

    e.GET("/article/:articleid", article.GetById(db))

    e.GET("/article/:articleid/cp", article.Cp(db))

    e.POST("/user", user.Create(db))
    e.PUT("/user", user.Update(db))
    e.DELETE("/user", user.Delete(db))

    e.POST("/user/follow/:userid", user.Follow(db))
    e.DELETE("/user/follow/:userid", user.UnFollow(db))

    e.POST("/user/tab/follow", user.TabFollow(db))
    e.DELETE("/user/tab/follow", user.TabUnFollow(db))

    e.GET("/user/:userid/cp", user.Cp(db))

    e.POST("/md", markdown.Default)

    e.Start(":8080")
}
