package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/integrationninjas/go-app/models"
)

const randomUserAPIURL = "https://randomuser.me/api/"

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	userData, err := fetchRandomUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user data: %v", err)
		return
	}

	encodeJSON(w, userData.Results[0]) // Encode and return the first user data
}

func fetchRandomUser() (models.UserData, error) {
	// Create a context with timeout to prevent hanging requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, randomUserAPIURL, nil)
	if err != nil {
		return models.UserData{}, err
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.UserData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.UserData{}, err
	}

	var userData models.UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return models.UserData{}, err
	}

	return userData, nil
}
