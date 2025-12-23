package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"size:255"`
	PostNum     int       `gorm:"default:0"`
	PostList    []Post    `gorm:"foreignkey:UserID"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"size:255"`
	PostContent string `gorm:"size:255"`
	CommentNum  int    `gorm:"default:0"`
	Status      string `gorm:"default:无评论"`
	UserID      uint
	User        User      `gorm:"foreignkey:UserID"`
	CommentList []Comment `gorm:"foreignkey:PostID"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID             uint   `gorm:"primary_key"`
	CommentContent string `gorm:"size:255"`
	PostID         uint
	Post           Post      `gorm:"foreignkey:PostID"`
	CreatedTime    time.Time `gorm:"autoCreateTime"`
	UpdatedTime    time.Time `gorm:"autoUpdateTime"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	user := User{
		ID: p.UserID,
	}
	tx.Take(&user)
	user.PostNum += 1
	tx.Save(&user)
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	post := Post{
		ID: c.PostID,
	}
	tx.Take(&post)
	post.CommentNum += 1
	if post.Status == "无评论" {
		post.Status = "有评论"
	}
	tx.Save(&post)
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	post := Post{
		ID: c.PostID,
	}
	tx.Take(&post)
	post.CommentNum -= 1
	if post.CommentNum < 0 {
		return errors.New("评论已删除,请勿重复删除")
	}
	if post.CommentNum == 0 && post.Status == "有评论" {
		post.Status = "无评论"
	}
	tx.Save(&post)
	return nil
}
