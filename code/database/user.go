package slack

import (
	"github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID 			gocql.UUID 	`json: "id"`
	Username 	string 		`json: "username"`
	Email 		string 		`json: "email"`
}

func CreateUser(c *gin.Context) {

	var user User

	if err:=c.ShouldBindJSON(&user); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := session.Query(query, user.ID, user.Username, user.Email).Exec(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func getUserById(id gocql.UUID) (User, error) {
	var user User
	query := "SELECT * FROM USER WHERE id=?"
	err := session.Query(query, id).Scan(&user.ID, &user.Username, &user.Email)
	return user, err
}

func getAllUsers() ([]User, error) {
	var users []User
	query := "SELECT * FROM USER"
	iteration := session.Query(query).Iter()

	var user User
	for iter.Scan($user.ID, &user.Username, &user.Email){
		users = append(users, user)
	}

	err := iter.Close()
	return users, err
}