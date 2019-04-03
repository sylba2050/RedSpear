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
        data := new(DB.Article)
        inDB := new(DB.Article)

        if err := c.Bind(data); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        db.Where("id = ?", data.ID).First(&inDB)
        inDB = data
        db.Save(&inDB)

        return c.HTML(http.StatusOK, "ok")
    }
}

func Delete(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("articleid")
        inDB := new(DB.Article)

        db.Where("id = ?", id).First(&inDB)
        db.Delete(&inDB)

        return c.HTML(http.StatusOK, "ok")
    }
}

func GetStocksByArticleId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("articleid")
        stocks := []DB.Stock{}

        db.Where("article_id = ?", id).Find(&stocks)
        return c.JSON(http.StatusOK, stocks)
    }
}

func Stock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        auth := new(DB.Auth)
        articleId := c.Param("articleid")

        if err := c.Bind(auth); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        stock := new(DB.Stock)
        stock.StockedUserId = auth.UserId
        stock.ArticleId = articleId

        db.Create(&stock)

        return c.HTML(http.StatusOK, "ok")
    }
}

func UnStock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        auth := new(DB.Auth)
        articleId := c.Param("articleid")
        stock := new(DB.Stock)

        if err := c.Bind(auth); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        err := db.Where("stocked_user_id = ? AND article_id = ?", auth.UserId, articleId).First(&stock).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusOK, "It is already unstock")
        }

        db.Delete(&stock)

        return c.HTML(http.StatusOK, "ok")
    }
}

func GetLikesByArticleId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("articleid")
        likes := []DB.Like{}

        db.Where("article_id = ?", id).Find(&likes)
        return c.JSON(http.StatusOK, likes)
    }
}

func Like(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        auth := new(DB.Auth)
        articleId := c.Param("articleid")

        if err := c.Bind(auth); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        like := new(DB.Like)
        like.LikedUserId = auth.UserId
        like.ArticleId = articleId

        db.Create(&like)

        return c.HTML(http.StatusOK, "ok")
    }
}

func UnLike(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        auth := new(DB.Auth)
        articleId := c.Param("articleid")
        like := new(DB.Like)

        if err := c.Bind(auth); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        err := db.Where("liked_user_id = ? AND article_id = ?", auth.UserId, articleId).First(&like).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusOK, "It is already unstock")
        }

        db.Delete(&like)

        return c.HTML(http.StatusOK, "ok")
    }
}

func GetCommentByArticleId(db *gorm.DB) echo.HandlerFunc {
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
