package models

type Post struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Topics      []Topic `json:"topics"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	PostType    string  `json:"post_type"`
	Author      string  `json:"author"`
	CreatedAt   string  `json:"created_at"`
}
