package slack

import (
	"github.com/Priyank-Vaidya/ChatOps/database/cassandra"
	// "github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

//Build a model in go with parameters channelId and userId and message to be sent to cassandra
// type Slack struct {
// 	ChannelId string `json:"channelId"`
// 	UserId string `json:"userId"`
// 	Message string `json:"message"`
// }

// // Make an http post request to send the message to slack and then post in cassandra
// func PostMessageToSlack(slack Slack) {
// 	// Set the Slack API URL
// 	url := "https://slack.com/api/chat.postMessage"

// 	// Create the request payload
// 	payload, err := json.Marshal(slack)
// 	if err != nil {
// 		log.Fatal("Failed to marshal request payload")
// 	}

// 	// Create the HTTP request
// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
// 	if err != nil {
// 		log.Fatal("Failed to create request")
// 	}

// 	// Set the request headers
// }





type SlackMessage struct {
	ChannelId string `json:"channelId"`
	UserId string `json:"userId"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func HelloSlackHandler(c *gin.Context) {

	err := godotenv.Load()
	if(err!=nil){
		log.Fatal("Error in loading ENV")
	}

	slackToken := os.Getenv("SLACK_AUTH_TOKEN")
	fmt.Println(slackToken)
	fmt.Println("ENTERED THE SLACK HANDLER")
	var slackMsg SlackMessage

	if err := c.ShouldBindJSON(&slackMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Failed to parse the Body"})
		return
	}

	url := "https://slack.com/api/chat.postMessage"

	payload, err := json.Marshal(slackMsg)
	fmt.Println(slackMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request payload"})
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the request"})
		return 
	}

	authToken := "Bearer " + slackToken 

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to Slack API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to Slack API"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent to Slack API"})

	// Now post the message to cassandra
	
	session, err := cassandra.GetCassandraSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra")
	}

	defer session.Close()

	if err := session.Query("INSERT INTO slack (channel_id, user_id, text) VALUES (?, ?, ?)",
		SlackMessage.ChannelId, SlackMessage.UserId, SlackMessage.Message).Exec(); err != nil {
		log.Fatal("Failed to insert data into Cassandra")
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent to Slack API and inserted into Cassandra"})
}