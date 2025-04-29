package routes

import (
	"go_project1/internal/handlers"
	"net/http"
	"strings"
)

// BlogPostsHandler handles all blog post related requests
func BlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is targeting a specific post or all posts
	path := r.URL.Path
	pathParts := strings.Split(path, "/")
	
	// Route for specific post (has ID): /api/posts/{id}
	if len(pathParts) > 3 && pathParts[len(pathParts)-1] != "" {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPost(w, r)
		case http.MethodPut:
			handlers.UpdatePost(w, r)
		case http.MethodDelete:
			handlers.DeletePost(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	}
	
	// Route for all posts: /api/posts
	switch r.Method {
	case http.MethodGet:
		handlers.GetAllPosts(w, r)
	case http.MethodPost:
		handlers.CreatePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes configures all API routes
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Define API routes
	mux.HandleFunc("/api/posts", BlogPostsHandler)
	mux.HandleFunc("/api/posts/", BlogPostsHandler) // Handle paths with trailing parts
	
	return mux
}