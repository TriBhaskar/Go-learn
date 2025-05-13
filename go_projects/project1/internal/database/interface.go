package database

import "go_project1/internal/models"

type DatabaseInterface interface {
    Initialize() error
    GetAllPosts() ([]models.BlogPost, error)
    GetPostByID(id string) (models.BlogPost, bool, error)
    CreatePost(post models.BlogPost) error
    UpdatePost(post models.BlogPost) error
    DeletePost(id string) (bool, error)
}