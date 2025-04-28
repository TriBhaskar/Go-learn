package models

import (
	"fmt"
	"time"
)

type BlogPost struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Published bool      `json:"published"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewBlogPost() BlogPost {
	now := time.Now().Format(time.RFC3339)
	return BlogPost{
		ID:        fmt.Sprintf("post-%d", time.Now().UnixNano()),
		CreatedAt: now,
		UpdatedAt: now,
	}
}