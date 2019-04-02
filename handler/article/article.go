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

func Post(db *gorm.DB) echo.HandlerFunc {
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

func Update(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Delete(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Stock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func UnStock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Like(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func UnLike(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Comment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func UpdateComment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func DeleteComment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Get(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        articles := []DB.Article{}
        db.Limit(20).Find(&articles)

        return c.JSON(http.StatusOK, articles)
    }
}

func GetNewer(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetVisited(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetPopulerPerDay(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetPopulerPerWeek(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetPopulerPerMonth(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetByUser(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userid := c.Param("userid")

        articles := []DB.Article{}
        db.Limit(20).Where("user_id = ?", userid).Find(&articles)

        return c.JSON(http.StatusOK, articles)
    }
}

func GetById(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Cp(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}
