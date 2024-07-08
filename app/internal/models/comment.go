package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	PostID    int       `json:"postId" gorm:"column:post_id"`
	ParentID  *int      `json:"parentId,omitempty" gorm:"column:parent_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}
