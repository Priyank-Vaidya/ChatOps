package slack

import (
	// "github.com/gocql/gocql"
	// "github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"net/http"
)

type SlackMessage struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func HelloSlackHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var slackMsg SlackMessage
	err := json.NewDecoder(r.Body).Decode(&slackMsg)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Set the Slack API URL
	url := "https://slack.com/api/chat.postMessage"

	// Create the request payload
	payload, err := json.Marshal(slackMsg)
	if err != nil {
		http.Error(w, "Failed to marshal request payload", http.StatusInternalServerError)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer xoxb-5467279451222-5473827817747-Zx35wT3SPy70GUWXh8sbXIrd")
	req.Header.Set("Content-Type", "application/json")

	// Send the request to Slack API
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to Slack API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to send message to Slack API", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent to Slack API"))
}
