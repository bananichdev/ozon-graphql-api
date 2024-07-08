package db

import (
	"fmt"

	"github.com/bananichdev/ozon-graphql-api/internal/errors"
	"github.com/bananichdev/ozon-graphql-api/internal/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func (pr *PostRepo) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := pr.DB.Find(&posts)
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}

func (pr *PostRepo) GetPostByID(id int) (*models.Post, error) {
	var post *models.Post
	err := pr.DB.First(&post, "id = ?", id)
	if err.Error != nil {
		return nil, errors.GenerateError(fmt.Sprintf("Post with id=%d does not exists", id))
	}
	return post, nil
}

func (pr *PostRepo) CreatePost(input models.NewPost) (*models.Post, error) {
	post := models.Post{Author: input.Author, Title: input.Title, Content: input.Content, CommentsDisabled: input.CommentsDisabled}
	err := pr.DB.Create(&post)

	if err.Error != nil {
		return nil, err.Error
	}

	return &post, nil
}
