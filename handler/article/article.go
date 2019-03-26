package article

import (
    "../../struct/DB"

    "os"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func POST(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        article := new(DB.Article)
        if err := c.Bind(article); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        db.Create(&article)
        return c.HTML(http.StatusOK, "ok")
    }
}

func GET(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        article := []DB.Article{}
        db.Limit(20).Find(&article)

        return c.JSON(http.StatusOK, article)
    }
}

func GetByUser(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        article := []DB.Article{}
        db.Limit(20).Where("user_id = ?", userid).Find(&article)

        return c.JSON(http.StatusOK, article)
    }
}

func GetLikeCount(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        article_id := c.Param("articleid")

        likes := []DB.Like{}
        numLikes := 0
        db.Where("article_id = ?", article_id).Find(&likes).Count(&numLikes)

        return c.JSON(http.StatusOK, numLikes)
    }
}

func GetStockCount(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        article_id := c.Param("articleid")

        stocks := []DB.Stock{}
        numStocks := 0
        db.Where("article_id = ?", article_id).Find(&stocks).Count(&numStocks)

        return c.JSON(http.StatusOK, numStocks)
    }
}
