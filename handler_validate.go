package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const MAX_CHIRP_LEN = 140

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	if len(params.Body) > MAX_CHIRP_LEN {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	profane := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(params.Body, " ")
	for i, word := range words {
		if _, ok := profane[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}

	cleanedBody := strings.Join(words, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBody,
	})
}
