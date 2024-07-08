package db_test

import (
	"fmt"
	"testing"

	"github.com/bananichdev/ozon-graphql-api/internal/db"
	"github.com/bananichdev/ozon-graphql-api/internal/models"
	"github.com/bananichdev/ozon-graphql-api/internal/settings"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCommentRepo(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", settings.DBHost, settings.DBUser, settings.DBPass, settings.TestDBName, settings.DBPort)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	DB.Migrator().CreateTable(&models.Post{})
	defer DB.Migrator().DropTable(&models.Post{})

	p := models.Post{
		ID:               1,
		Author:           "test_author",
		Title:            "test_author",
		Content:          "test_content",
		CommentsDisabled: false,
	}
	DB.Create(&p)

	r := db.CommentRepo{DB: DB}

	t.Run("should create and return Comment", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Comment{})
		defer DB.Migrator().DropTable(&models.Comment{})

		want := models.Comment{
			ID:       1,
			Author:   "test",
			PostID:   1,
			ParentID: nil,
			Content:  "test",
		}
		get, err := r.CreateComment(models.NewComment{
			Author:   "test",
			PostID:   1,
			ParentID: nil,
			Content:  "test",
		})

		require.Equal(t, nil, err)
		require.Equal(t, want.ID, get.ID)
		require.Equal(t, want.Author, get.Author)
		require.Equal(t, want.PostID, get.PostID)
		require.Equal(t, want.ParentID, get.ParentID)
		require.Equal(t, want.Content, get.Content)
	})

	t.Run("should return Replies", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Comment{})
		defer DB.Migrator().DropTable(&models.Comment{})

		c := models.Comment{
			ID:       1,
			Author:   "test",
			PostID:   1,
			ParentID: nil,
			Content:  "test",
		}
		DB.Create(&c)

		ptr := func(i int) *int { return &i }(1)

		want := []models.Comment{
			{
				ID:       1,
				Author:   "test",
				PostID:   1,
				ParentID: ptr,
				Content:  "test",
			},
			{
				ID:       2,
				Author:   "test",
				PostID:   1,
				ParentID: ptr,
				Content:  "test",
			},
		}
		DB.Create(&want)
		get, err := r.GetReplies(p.ID)

		require.Equal(t, nil, err)
		for i := range get {
			require.Equal(t, want[i].ID, get[i].ID)
			require.Equal(t, want[i].Author, get[i].Author)
			require.Equal(t, want[i].PostID, get[i].PostID)
			require.Equal(t, want[i].ParentID, get[i].ParentID)
			require.Equal(t, want[i].Content, get[i].Content)
		}
	})

	t.Run("should return all Comments by postID", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Comment{})
		defer DB.Migrator().DropTable(&models.Comment{})

		want := []models.Comment{
			{
				ID:       1,
				Author:   "test",
				PostID:   1,
				ParentID: nil,
				Content:  "test",
			},
			{
				ID:       2,
				Author:   "test",
				PostID:   1,
				ParentID: nil,
				Content:  "test",
			},
		}
		DB.Create(&want)
		get, err := r.GetAllCommentsByPostID(p.ID)

		require.Equal(t, nil, err)
		for i := range get {
			require.Equal(t, want[i].ID, get[i].ID)
			require.Equal(t, want[i].Author, get[i].Author)
			require.Equal(t, want[i].PostID, get[i].PostID)
			require.Equal(t, want[i].ParentID, get[i].ParentID)
			require.Equal(t, want[i].Content, get[i].Content)
		}
	})

	t.Run("should return Comments", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Comment{})
		defer DB.Migrator().DropTable(&models.Comment{})

		want := []models.Comment{
			{
				ID:       1,
				Author:   "test",
				PostID:   1,
				ParentID: nil,
				Content:  "test",
			},
			{
				ID:       2,
				Author:   "test",
				PostID:   1,
				ParentID: nil,
				Content:  "test",
			},
		}
		DB.Create(&want)
		get, err := r.GetComments(p.ID, 1, 0)

		require.Equal(t, nil, err)
		require.Equal(t, want[0].ID, get[0].ID)
		require.Equal(t, want[0].Author, get[0].Author)
		require.Equal(t, want[0].PostID, get[0].PostID)
		require.Equal(t, want[0].ParentID, get[0].ParentID)
		require.Equal(t, want[0].Content, get[0].Content)
	})
}
