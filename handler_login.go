package main

import (
	"encoding/json"
	"net/http"

	"github.com/db-cooper7/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	if params.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email required", nil)
		return
	}

	dbUser, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	if err = auth.CheckPasswordHash(dbUser.HashedPassword, params.Password); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", nil) // should be returning nil to avoid giving bad actors relevant information
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:    dbUser.ID,
			Email: dbUser.Email,
		},
	})
}
