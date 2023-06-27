package main

import (
	"log"
	// "github.com/gin-gonic/gin"
	// "github.com/gocql/gocql"
	// "github.com/joho/godotenv"
	// "github.com/Priyank-Vaidya/Slack-ChatOps/tree/master/routers"
	"fmt"
	"os"
	""
)

func main() {
	
	// err := database.initCassandra();
	// if err != nil {
	// 	log.Fatal("Failed to initialize Cassandra:", err)
	// }

	// err := godotenv.Load()
	// if(err!=nil){
	// 	log.Fatal("Error in loading ENV")
	// }

	// slackToken := os.Getenv("SLACK_AUTH_TOKEN")
	// fmt.Println(slackToken)

	// router := gin.Default()

	// routers.InitializeRoutes(router)

	

	// Start the server
	log.Println("Server started on http://localhost:8000")
	log.Fatal(router.Run(":8000"))
}