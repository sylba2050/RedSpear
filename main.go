package main

import (
    "github.com/sylba2050/RedSpear/handler/markdown"
    "github.com/sylba2050/RedSpear/handler/article"
    "github.com/sylba2050/RedSpear/handler/user"
    "github.com/sylba2050/RedSpear/struct/DB"

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

    db.AutoMigrate(&DB.User{})
    db.AutoMigrate(&DB.Article{})
    db.AutoMigrate(&DB.ArticleLike{})
    db.AutoMigrate(&DB.ArticleStock{})
    db.AutoMigrate(&DB.ArticleComment{})
    db.AutoMigrate(&DB.CommentLike{})
    db.AutoMigrate(&DB.User{})
    db.AutoMigrate(&DB.FollowTag{})
    db.AutoMigrate(&DB.FF{})

    e.Static("/img", "img")
    e.Static("/js", "js")
    e.Static("/css", "css")

    e.File("/", "html/index.html")
    e.File("/write", "html/write.html")

    e.GET("/article/:articleid", article.GetById(db))
    e.POST("/article", article.Post(db))
    e.PUT("/article", article.Update(db))
    e.DELETE("/article/:articleid", article.Delete(db))

    e.GET("/article/stock/:articleid", article.GetStocksByArticleId(db))
    e.POST("/article/stock/:articleid", article.Stock(db))
    e.DELETE("/article/stock/:articleid", article.UnStock(db))

    e.GET("/article/like/:articleid", article.GetLikesByArticleId(db))
    e.POST("/article/like/:articleid", article.Like(db))
    e.DELETE("/article/like/:articleid", article.UnLike(db))

    e.GET("/article/comments/:articleid", article.GetCommentsByArticleId(db))
    e.POST("/article/comment/:articleid", article.Comment(db))
    e.PUT("/article/comment/update/:commentid", article.UpdateComment(db))
    e.DELETE("/article/comment/delete/:commentid", article.DeleteComment(db))

    e.GET("/articles", article.Get(db))
    e.GET("/articles/newer", article.GetNewer(db))
    e.GET("/articles/visited", article.GetVisited(db))
    e.GET("/articles/populer/day", article.GetPopulerPerDay(db))
    e.GET("/articles/populer/week", article.GetPopulerPerWeek(db))
    e.GET("/articles/populer/month", article.GetPopulerPerMonth(db))
    e.GET("/articles/:userid", article.GetByUser(db))

    e.GET("/article/cp/:articleid", article.Cp(db))

    e.POST("/user", user.Create(db))
    e.PUT("/user", user.Update(db))
    e.DELETE("/user", user.Delete(db))

    e.POST("/user/follow/:userid", user.Follow(db))
    e.DELETE("/user/follow/:userid", user.UnFollow(db))

    e.POST("/user/tab/follow", user.TabFollow(db))
    e.DELETE("/user/tab/follow", user.TabUnFollow(db))

    e.GET("/user/cp/:userid", user.Cp(db))
    e.GET("/user/stock/:userid", user.GetStockByUserId(db))
    e.GET("/user/like/:userid", user.GetLikeByUserId(db))
    e.GET("/user/comment/:userid", user.GetCommentByUserId(db))

    e.POST("/md", markdown.Default)

    e.Start(":8080")
}
