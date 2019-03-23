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
    Content string `json:"content" form:"content" query:"content"`
    Cp uint `json:"cp" form:"cp" query:"cp"`
    Like string `json:"like" form:"like" query:"like"`
    Comment string `json:"comment" form:"comment" query:"comment"`
}

type Draft struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Content string `json:"content" form:"content" query:"content"`
}

type User struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Cp uint `json:"cp" form:"cp" query:"cp"`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
    StockList string `json:"stocklist" form:"stocklist" query:"stocklist"`
}

type Comment struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Content string `json:"content" form:"content" query:"content"`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
}
