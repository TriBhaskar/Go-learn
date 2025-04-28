package models

type BlogPost struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	Tags      []string `json:"tags"`
	Published bool     `json:"published"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}