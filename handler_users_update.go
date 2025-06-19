package main

import (
	"encoding/json"
	"net/http"

	"github.com/db-cooper7/chirpy/internal/auth"
	"github.com/db-cooper7/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not find token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not validate user", err)
		return
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode request body", err)
		return
	}

	password, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not encrypt password", err)
		return
	}

	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:          params.Email,
		HashedPassword: password,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not update user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Email:       user.Email,
			IsChirpyRed: user.IsChirpyRed,
		},
	})
}
