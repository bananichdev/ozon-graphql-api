package graph

import (
	"github.com/bananichdev/ozon-graphql-api/db"
	"github.com/bananichdev/ozon-graphql-api/models"
)

type Resolver struct {
	posts       []*models.Post
	comments    []*models.Comment
	PostRepo    db.PostRepo
	CommentRepo db.CommentRepo
}
