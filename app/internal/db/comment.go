package db

import (
	"fmt"

	"github.com/bananichdev/ozon-graphql-api/internal/errors"
	"github.com/bananichdev/ozon-graphql-api/internal/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	DB *gorm.DB
}

func (cr *CommentRepo) GetComments(postID, first, skip int) ([]*models.Comment, error) {
	pr := PostRepo{DB: cr.DB}
	post, err := pr.GetPostByID(postID)
	if err != nil {
		return nil, errors.InternalError
	}
	if post.CommentsDisabled {
		return nil, errors.GenerateError(fmt.Sprintf("Post with id=%d have comments disabled", postID))
	}

	var comments []*models.Comment
	e := cr.DB.Limit(first).Offset(skip).Find(&comments, "post_id = ?", postID)
	if e.Error != nil {
		return nil, errors.InternalError
	}
	return comments, nil
}

func (cr *CommentRepo) GetAllCommentsByPostID(postID int) ([]*models.Comment, error) {
	var comments []*models.Comment
	e := cr.DB.Find(&comments, "post_id = ?", postID)
	if e.Error != nil {
		return nil, errors.InternalError
	}
	return comments, nil
}

func (cr *CommentRepo) GetReplies(parentId int) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := cr.DB.Find(&comments, "parent_id = ?", parentId)
	if err.Error != nil {
		return nil, errors.InternalError
	}
	return comments, nil
}

func (cr *CommentRepo) CreateComment(input models.NewComment) (*models.Comment, error) {
	pr := PostRepo{DB: cr.DB}
	post, err := pr.GetPostByID(input.PostID)
	if err != nil {
		return nil, errors.InternalError
	}
	if post.CommentsDisabled {
		return nil, errors.GenerateError(fmt.Sprintf("Post with id=%d have comments disabled", input.PostID))
	}
	if input.ParentID != nil {
		var parent *models.Comment
		e := cr.DB.First(&parent, "id = ?", *input.ParentID)
		if e.Error != nil {
			return nil, errors.GenerateError(fmt.Sprintf("Comment with id=%d does not exists", *input.ParentID))
		}
		if parent.PostID != input.PostID {
			return nil, errors.GenerateError("The comment and the reply to it must be in one post")
		}
	}

	comment := models.Comment{Author: input.Author, PostID: input.PostID, ParentID: input.ParentID, Content: input.Content}
	cr.DB.Create(&comment)

	return &comment, nil
}
