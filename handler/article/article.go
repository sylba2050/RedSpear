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
