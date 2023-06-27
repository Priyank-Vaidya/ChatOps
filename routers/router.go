package routers

import (
	"github.com/Priyank-Vaidya/ChatOps/model"
	"net/http"
	
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/createuser", models.CreateUser)
	router.GET("/getuserbyid/:id", models.getUserById)
	router.GET("/getallusers", models.getAllUsers)
	http.HandleFunc("/slackapi", model.slack.HelloSlackHandler)
	
}

