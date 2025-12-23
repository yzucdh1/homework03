package main

import (
	"errors"
	"fmt"
	"github.com/yzucdh1/homework03/global"
	"github.com/yzucdh1/homework03/model"
	"gorm.io/gorm"
)

func main() {
	global.Connect()
	//global.DB.AutoMigrate(&model.Post{})
	//AddUser()
	//fmt.Println(QueryByUserId(1))
	//AddPost()
	//AddComment()
	//fmt.Println(QueryPostByMostCommentNum())
	deleteComment(5)
}

func AddUser() {
	// 增加用户
	user := model.User{
		Name: "王五",
	}
	global.DB.Debug().Create(&user)
}

func AddPost() {
	user := model.User{ID: 3}
	global.DB.Take(&user)
	post := model.Post{
		Title:       "python语言技术入门到精通",
		PostContent: "python语言技术从入门到精通.....",
		UserID:      user.ID,
	}
	global.DB.Create(&post)
}

func AddComment() {
	post := model.Post{ID: 6}
	global.DB.Take(&post)
	comment := model.Comment{
		CommentContent: "python是编译型还是解释型语言?.....",
		PostID:         post.ID,
	}
	global.DB.Debug().Create(&comment)
}

func QueryByUserId(Id uint) model.User {
	user := model.User{
		ID: Id,
	}
	global.DB.Debug().Preload("PostList.CommentList").Take(&user)
	return user
}

func QueryPostByMostCommentNum() model.Post {
	post := model.Post{}
	global.DB.Debug().Order("comment_num desc").Take(&post)
	return post
}

func deleteComment(Id uint) {
	comment := &model.Comment{
		ID: Id,
	}
	result := global.DB.Take(&comment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("评论不存在", result.Error)
		} else {
			fmt.Println("查询错误", result.Error)
		}
		return
	}
	global.DB.Debug().Delete(&comment)
}
