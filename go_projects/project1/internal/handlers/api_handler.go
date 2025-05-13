package handlers

import (
	"encoding/json"
	"go_project1/internal/database"
	"go_project1/internal/models"
	"log"
	"net/http"
	"strings"
	"time"
)

// Handler struct to hold database interface
type Handler struct {
    db database.DatabaseInterface
}

// NewHandler creates a new handler instance
func NewHandler(db database.DatabaseInterface) *Handler {
    return &Handler{db: db}
}

func (h *Handler) BlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is targeting a specific post or all posts
	path := r.URL.Path
	pathParts := strings.Split(path, "/")
	
	// Route for specific post (has ID): /api/posts/{id}
	if len(pathParts) > 3 && pathParts[len(pathParts)-1] != "" {
		switch r.Method {
		case http.MethodGet:
			h.GetPost(w, r)
		case http.MethodPut:
			h.UpdatePost(w, r)
		case http.MethodDelete:
			h.DeletePost(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	}
	
	// Route for all posts: /api/posts
	switch r.Method {
	case http.MethodGet:
		h.GetAllPosts(w, r)
	case http.MethodPost:
		h.CreatePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetAllPosts returns all blog posts
func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllPosts called")
	posts, err := h.db.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// GetPost returns a specific blog post by ID
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("GetPost called")
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	post, exists, err := h.db.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}

// CreatePost creates a new blog post
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("CreatePost called")
	var post models.BlogPost
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique ID and set timestamps
	newPost := models.NewBlogPost()
	post.ID = newPost.ID
	post.CreatedAt = newPost.CreatedAt
	post.UpdatedAt = newPost.UpdatedAt

	// Create the blog post
	if err := h.db.CreatePost(post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// UpdatePost updates an existing blog post
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdatePost called")
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	// Fetch the existing post
	existingPost, exists, err := database.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Parse the updated post from request body
	var updatedPost models.BlogPost
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    updatedPost.ID = id
    updatedPost.CreatedAt = existingPost.CreatedAt
    updatedPost.UpdatedAt = time.Now().Format(time.RFC3339)

	// Update the post
	if err := h.db.UpdatePost(updatedPost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedPost)
}

// DeletePost deletes a blog post by ID
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	log.Println("DeletePost called")
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	// Delete the post
	deleted, err := h.db.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
    // Create success response
    response := map[string]interface{}{
        "message": "Post deleted successfully",
        "id":      id,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}