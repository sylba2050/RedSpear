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
        return c.HTML(http.StatusOK, "ok")
    }
}
