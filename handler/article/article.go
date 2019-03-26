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

type ResponseData struct {
    gorm.Model
    UserId string `json:"userid"`
    Title string `json:"title"`
    Content string `json:"content"`
    Cp uint `json:"cp"`
    NumLikes uint `json:"numlikes"`
    NumStocks uint `json:"numstocks"`
}

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

func ConcatLikeAndStock(db *gorm.DB, articles []DB.Article) []ResponseData {
    res := []ResponseData{}
    for _, article := range articles {
        tmp := ResponseData{ article.Model, article.UserId, article.Title, article.Content, article.Cp, 0, 0 }

        likes := []DB.Like{}
        var numLikes uint = 0
        db.Where("article_id = ?", article.ID).Find(&likes).Count(&numLikes)
        tmp.NumLikes = numLikes

        stocks := []DB.Stock{}
        var numStocks uint = 0
        db.Where("article_id = ?", article.ID).Find(&stocks).Count(&numStocks)

        tmp.NumStocks = numStocks

        res = append(res, tmp)
    }

    return res
}

func GET(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        articles := []DB.Article{}
        db.Limit(20).Find(&articles)

        res := ConcatLikeAndStock(db, articles)

        return c.JSON(http.StatusOK, res)
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
