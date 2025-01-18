package handler

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"blog/internal/database"
	"blog/pkg/auth"
)

type AuthHandler struct {
	DB  *database.MongoDB
	JWT *auth.JWTHandler
}

func (h *AuthHandler) HandleAuth() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse the request body for user authentication
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&req); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // Get the user from the database based on email
        user, err := h.DB.GetUserByEmail(req.Email)
        if err != nil {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        // Verify the password against the user's password hash
        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
            http.Error(w, "Invalid password", http.StatusUnauthorized)
            return
        }

        // Generate and send back the JWT
        token, err := h.JWT.Generate(user)
        if err != nil {
            http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
            return
        }

        // Create a response structure
        response := struct {
            Token string `json:"token"`
        }{
            Token: token,
        }

        // Send the response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
