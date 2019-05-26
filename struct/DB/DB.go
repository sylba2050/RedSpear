package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	gorm.Model
	UserId   string `json:"userid" form:"userid" query:"userid"`
	Pw       string `json:"pw" form:"pw" query:"pw"`
	AuthCode string `json:"auth_code" form:"auth_code" query:"auth_code"`
	UserName string `json:"user_name" form:"user_name" query:"user_name"`
	IsLogin  bool   `json:"is_login" form:"is_login" query:"is_login"`
	Mail     string `json:"mail" form:"mail" query:"mail"`
	Icon     string `json:"icon" form:"icon" query:"icon"`
	Cp       uint16 `json:"cp" form:"cp" query:"cp"`
}

type Article struct {
	gorm.Model
	UserId  string `json:"userid" form:"userid" query:"userid"`
	Title   string `json:"title" form:"title" query:"title"`
	Content string `json:"content" form:"content" query:"content"`
	Tag     string `json:"tag" form:"tag" query:"tag"`
	Privacy string `json:"privacy" form:"privacy" query:"privacy"`
	Cp      uint16 `json:"cp" form:"cp" query:"cp"`
}

type ArticleLike struct {
	gorm.Model
	UserId    string `json:"userid" form:"userid" query:"userid"`
	ArticleId uint   `json:"articleid" form:"articleid" query:"articleid"`
}

type ArticleStock struct {
	gorm.Model
	UserId    string `json:"userid" form:"userid" query:"userid"`
	ArticleId uint   `json:"articleid" form:"articleid" query:"articleid"`
}

type ArticleComment struct {
	gorm.Model
	UserId    string `json:"userid" form:"userid" query:"userid"`
	ArticleId uint   `json:"articleid" form:"articleid" query:"articleid"`
	Content   string `json:"content" form:"content" query:"content"`
}

type CommentLike struct {
	gorm.Model
	UserId    string `json:"userid" form:"userid" query:"userid"`
	CommentId uint   `json:"commentid" form:"commentid" query:"commentid"`
}

type Tag struct {
	gorm.Model
	ArticleId uint `json:"articleid" form:"articleid" query:"articleid"`
	TagId     uint `json:"tag" form:"tag" query:"tag"`
}

type FollowTag struct {
	gorm.Model
	UserId string `json:"userid" form:"userid" query:"userid"`
	Tag    string `json:"tag" form:"tag" query:"tag"`
}

type FF struct {
	gorm.Model
	Follow   string `json:"follow" form:"follow" query:"follow"`
	Follower string `json:"follower" form:"follower" query:"follower"`
}
