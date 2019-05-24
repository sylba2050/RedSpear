package article

import (
    "github.com/sylba2050/RedSpear/struct/DB"
    _ "github.com/sylba2050/RedSpear/struct/Point"

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
        stocks := []DB.ArticleStock{}

        db.Where("article_id = ?", id).Find(&stocks)
        return c.JSON(http.StatusOK, stocks)
    }
}

func Stock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        stock := new(DB.ArticleStock)

        if err := c.Bind(stock); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

		// TODO: validate nil
        db.Create(&stock)

        return c.HTML(http.StatusOK, "ok")
    }
}

func UnStock(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        stock := new(DB.ArticleStock)

        if err := c.Bind(stock); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

		recode := new(DB.ArticleStock)
        err := db.Where("user_id = ? AND article_id = ?", stock.UserId, stock.ArticleId).First(&recode).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusOK, "It is already unstock")
        }

        db.Delete(&recode)

        return c.HTML(http.StatusOK, "ok")
    }
}

func GetLikesByArticleId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("articleid")
        likes := []DB.ArticleLike{}

        db.Where("articleid = ?", id).Find(&likes)
        return c.JSON(http.StatusOK, likes)
    }
}

func Like(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        like := new(DB.ArticleLike)

        if err := c.Bind(like); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        db.Create(&like)

        return c.HTML(http.StatusOK, "ok")
    }
}

func UnLike(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        like := new(DB.ArticleLike)

        if err := c.Bind(like); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        recode := new(DB.ArticleLike)
        err := db.Where("user_id = ? AND article_id = ?", like.UserId, like.ArticleId).First(&recode).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusOK, "It is already unlike")
        }

        db.Delete(&recode)

        return c.HTML(http.StatusOK, "ok")
    }
}

func GetCommentsByArticleId(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("articleid")
        comments := []DB.ArticleComment{}

        db.Where("article_id = ?", id).Find(&comments)
        return c.JSON(http.StatusOK, comments)
    }
}

func Comment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        comment := new(DB.ArticleComment)

        if err := c.Bind(comment); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        db.Create(&comment)
        return c.HTML(http.StatusOK, "ok")
    }
}

func UpdateComment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        comment := new(DB.ArticleComment)
        if err := c.Bind(comment); err != nil {
            fmt.Fprintln(os.Stderr, err)
            return err
        }

        old := new(DB.ArticleComment)
        err := db.Where("id = ?", comment.ID).First(&old).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusNotFound, "Not Found")
        }

        old.Content = comment.Content
        db.Save(&old)

        return c.HTML(http.StatusOK, "ok")
    }
}

func DeleteComment(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        comment := new(DB.ArticleComment)

        err := db.Where("id = ?", comment.ID).First(&comment).Error
        if gorm.IsRecordNotFoundError(err) {
            return c.HTML(http.StatusNotFound, "Not Found")
        }
        db.Delete(&comment)

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
        articleid := c.Param("articleid")

        article := new(DB.Article)
        db.Where("id = ?", articleid).First(&article)

        return c.JSON(http.StatusOK, article)
    }
}

func Cp(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.HTML(http.StatusOK, "ok")
    }
}
