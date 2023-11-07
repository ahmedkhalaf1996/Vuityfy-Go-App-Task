package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	CreatorId     int    `json:"creator_id"`
	PostTitle     string `json:"post_title"`
	PostBody      string `json:"post_body"`
	PostFile      string `json:"post_file"`
	IsDirectScore bool   `json:"is_direct_score"`
}

// type Comments struct {
// 	gorm.Model
// 	PostId      int    `gorm:"primaryKey"`
// 	UserId      int    `gorm:"primaryKey"`
// 	CommentBody string `json:"comment_body"`
// 	UserName    string `json:"userName"`
// 	UserPicture string `json:"userPicture"`
// }

// type Likes struct {
// 	gorm.Model
// 	PostId int `gorm:"primaryKey"`
// 	UserId int `gorm:"primaryKey"`
// }
