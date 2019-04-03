package user

import (
    // "../../struct/DB"

    "net/http"

    "github.com/labstack/echo"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func Create(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
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

func Follow(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func UnFollow(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func TabFollow(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func TabUnFollow(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func Cp(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetStockByUserId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetLikeByUserId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}

func GetCommentByUserId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}
