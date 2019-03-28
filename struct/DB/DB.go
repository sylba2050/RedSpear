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
    tag string `json:"tag" form:"tag" query:"tag"`
    privacy string `json:"privacy" form:"privacy" query:"privacy"`
}

type ArticleStatus struct {
    gorm.Model
    NumComment uint `json:"numcomment" form:"numcomment" query:"numcomment"`
    NumStock uint `json:"numstock" form:"numstock" query:"numstock"`
    NumLike uint `json:"numlike" form:"numlike" query:"numlike"`
}

type Like struct {
    gorm.Model
    LikedUserId string `json:"likeduserid" form:"likeduserid" query:"likeduserid"`
    ArticleId string `json:"articleid" form:"articleid" query:"articleid"`
}

type Stock struct {
    gorm.Model
    StockedUserId string `json:"stockuserid" form:"stockuserid" query:"stockuserid"`
    ArticleId string `json:"articleid" form:"articleid" query:"articleid"`
}

type Comment struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    Content string `json:"content" form:"content" query:"content"`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
}

type CpForArticle struct {
    ArticleId string `json:"articleid" form:"articleid" query:"articleid"`
    //TODO 時系列データをいい感じに
    DummyCp string `json:"cp" form:"cp" query:"cp"`
}

type User struct {
    gorm.Model
    UserId string `json:"userid" form:"userid" query:"userid"`
    UserName string `json:"username" form:"usern"me query:"usern"me`
    LikeList string `json:"likelist" form:"likelist" query:"likelist"`
    StockList string `json:"stocklist" form:"stocklist" query:"stocklist"`
    Mail string `json:"mail" form:"mail" query:"mail"`
    Icon string `json:"icon" form:"icon" query:"icon"`
}

type UserFollowTag struct {
    UserId string `json:"userid" form:"userid" query:"userid"`
    Tags string `json:"tags" form:"tags" query:"tags"`
}

type FF struct {
    Follow string `json:"follow" form:"follow" query:"follow"`
    Follower string `json:"follower" form:"follower" query:"follower"`
}

type CpForUser struct {
    UserId string `json:"userid" form:"userid" query:"userid"`
    //TODO 時系列データをいい感じに
    DummyCp string `json:"cp" form:"cp" query:"cp"`
}
