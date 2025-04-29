package handlers

import (
	"encoding/json"
	"go_project1/internal/database"
	"go_project1/internal/models"
	"net/http"
	"strings"
	"time"
)

// GetAllPosts returns all blog posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllBlogPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// GetPost returns a specific blog post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	post, exists, err := database.GetBlogPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// CreatePost creates a new blog post
func CreatePost(w http.ResponseWriter, r *http.Request) {
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
	if err := database.CreateBlogPost(post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// UpdatePost updates an existing blog post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	// Fetch the existing post
	existingPost, exists, err := database.GetBlogPostByID(id)
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

	// Make sure we're using the ID from the URL and preserving created_at
	updatedPost.ID = id
	updatedPost.CreatedAt = existingPost.CreatedAt
	updatedPost.UpdatedAt = time.Now().Format(time.RFC3339)

	// Update the post
	if err := database.UpdateBlogPost(updatedPost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

// DeletePost deletes a blog post by ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	// URL format: /api/posts/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := pathParts[len(pathParts)-1]

	// Delete the post
	deleted, err := database.DeleteBlogPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Return success
	w.WriteHeader(http.StatusNoContent)
}