package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/integrationninjas/go-app/models"
)

const randomUserAPIURL = "https://randomuser.me/api/"

// HTTP client with timeout for making external API requests
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	userData, err := fetchRandomUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user data: %v", err)
		return
	}

	// Validate that we have at least one result
	if len(userData.Results) == 0 {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "No user data returned from API")
		return
	}

	encodeJSON(w, userData.Results[0]) // Encode and return the first user data
}

func fetchRandomUser() (models.UserData, error) {
	// Execute GET request with timeout
	resp, err := httpClient.Get(randomUserAPIURL)
	if err != nil {
		return models.UserData{}, err
	}
	defer resp.Body.Close()

	// Validate HTTP status code
	if resp.StatusCode != http.StatusOK {
		return models.UserData{}, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

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
