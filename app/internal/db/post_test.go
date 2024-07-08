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

func TestPostRepo(t *testing.T) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", settings.DBHost, settings.DBUser, settings.DBPass, settings.TestDBName, settings.DBPort)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	r := db.PostRepo{DB: DB}

	t.Run("should create and return Post", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Post{})
		defer DB.Migrator().DropTable(&models.Post{})

		want := models.Post{
			ID:               1,
			Author:           "test_author",
			Title:            "test_author",
			Content:          "test_content",
			CommentsDisabled: false,
		}
		get, err := r.CreatePost(models.NewPost{
			Author:           "test_author",
			Title:            "test_author",
			Content:          "test_content",
			CommentsDisabled: false,
		})

		require.Equal(t, nil, err)
		require.Equal(t, want.ID, get.ID)
		require.Equal(t, want.Author, get.Author)
		require.Equal(t, want.Title, get.Title)
		require.Equal(t, want.Content, get.Content)
		require.Equal(t, want.CommentsDisabled, get.CommentsDisabled)
	})

	t.Run("should return current Post", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Post{})
		defer DB.Migrator().DropTable(&models.Post{})

		want := models.Post{
			ID:               1,
			Author:           "test_author",
			Title:            "test_author",
			Content:          "test_content",
			CommentsDisabled: false,
		}
		DB.Create(&want)
		get, err := r.GetPostByID(want.ID)

		require.Equal(t, nil, err)
		require.Equal(t, want.ID, get.ID)
		require.Equal(t, want.Author, get.Author)
		require.Equal(t, want.Title, get.Title)
		require.Equal(t, want.Content, get.Content)
		require.Equal(t, want.CommentsDisabled, get.CommentsDisabled)
	})

	t.Run("should return current Posts", func(t *testing.T) {
		DB.Migrator().CreateTable(&models.Post{})
		defer DB.Migrator().DropTable(&models.Post{})

		want := []models.Post{
			{
				ID:               1,
				Author:           "test_author",
				Title:            "test_author",
				Content:          "test_content",
				CommentsDisabled: false,
			},
			{
				ID:               2,
				Author:           "test_author_2",
				Title:            "test_author_2",
				Content:          "test_content_2",
				CommentsDisabled: true,
			},
		}
		DB.Create(&want)
		get, err := r.GetPosts()

		for i := range want {
			require.Equal(t, nil, err)
			require.Equal(t, want[i].ID, get[i].ID)
			require.Equal(t, want[i].Author, get[i].Author)
			require.Equal(t, want[i].Title, get[i].Title)
			require.Equal(t, want[i].Content, get[i].Content)
			require.Equal(t, want[i].CommentsDisabled, get[i].CommentsDisabled)
		}
	})
}
