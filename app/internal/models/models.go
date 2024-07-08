package models

type Mutation struct {
}

type NewComment struct {
	Author   string `json:"author"`
	PostID   int    `json:"postId"`
	ParentID *int   `json:"parentId,omitempty"`
	Content  string `json:"content"`
}

type NewPost struct {
	Author           string `json:"author"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	CommentsDisabled bool   `json:"commentsDisabled"`
}

type Query struct {
}

type Subscription struct {
}
