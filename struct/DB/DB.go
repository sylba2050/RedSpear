package DB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

type Auth struct {
    UserId string `json:"userid" form:"userid" query:"userid"`
    PW string `json:"pw" form:"pw" query:"pw"`
}

type Article struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Title string `json:"title" form:"title" query:"title"`
    Content string `json:"content" form:"content" query:"content"`
    Cp uint `json:"cp" form:"cp" query:"cp"`
}

type Draft struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Title string `json:"title" form:"title" query:"title"`
    Content string `json:"content" form:"content" query:"content"`
}

type User struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Cp uint `json:"cp" form:"cp" query:"cp"`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
    StockList string `json:"stocklist" form:"stocklist" query:"stocklist"`
}

type Like struct {
    gorm.Model
    LikedUserId string `json:"userid" form:"userid" query:"userid"`
    ArticleId string `json:"articleid" form:"articleid" query:"articleid"`
}

type Stock struct {
    gorm.Model
    StockedUserId string `json:"userid" form:"userid" query:"userid"`
    ArticleId string `json:"articleid" form:"articleid" query:"articleid"`
}

type Comment struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Content string `json:"content" form:"content" query:"content"`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
}
