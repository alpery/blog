package handler

import (
	"encoding/json"
	"net/http"

	"blog/internal/database"
	"blog/internal/model"

	"github.com/google/uuid"
)

type BlogHandler struct {
	DB *database.MongoDB
}

func (h *BlogHandler) HandleBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the request body
		var req struct {
			//TODO: add a real blog post here
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		// Decode the request body into the `req` struct
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create a new blog post
		post := model.BlogPost{
			ID:      uuid.New(),
			Title:   req.Title,
			Content: req.Content,
		}

		// Create a new blog post in the database
		_, err := h.DB.CreateBlog(post)
		if err != nil {
			http.Error(w, "Failed to create blog post", http.StatusInternalServerError)
			return
		}

		// Send back the blog post
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Blog post created successfully"))
	}
}
