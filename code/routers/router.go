package routers

import (
	"github.com/Priyank-Vaidya/ChatOps/model"
	// "net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	
)

func InitializeRoutes(router *gin.Engine) {
	fmt.Println("ENTERED THE ROUTERS.GO")
	// router.POST("/createuser", models.CreateUser)
	// router.GET("/getuserbyid/:id", models.getUserById)
	// router.GET("/getallusers", models.getAllUsers)
	router.POST("/slackapi", slack.HelloSlackHandler)

	//router.POST for to deploy the function on AWS Lambda and API Gateway
	
}

