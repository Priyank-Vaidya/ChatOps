package routers

import (
	"../models"
	"net/http"
	"../models/slack"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/createuser", models.CreateUser)
	router.GET("/getuserbyid/:id", models.getUserById)
	router.GET("/getallusers", models.getAllUsers)
	http.HandleFunc("/slackapi", slack.HelloSlackHandler)
	
}

