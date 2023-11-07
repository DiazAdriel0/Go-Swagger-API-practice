package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// user represents data about a user.
type user struct {
    ID     			string  `json:"id"`
    Username  		string  `json:"username"`
    Email			string  `json:"email"`
    Age 			int 	`json:"age"`
	Transactions	int		`json:"transactions"`
}

// users slice to seed record user data.
var users = []user{
    {ID: "1", Username: "Leanne Graham", Email: "Sincere@april.biz", Age: 56, Transactions: 1520},
    {ID: "2", Username: "Ervin Howell", Email: "Shanna@melissa.tv", Age: 27, Transactions: 207},
    {ID: "3", Username: "Clementine Bauch", Email: "Nathan@yesenia.net", Age: 39, Transactions: 45},
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
	router.POST("/users", postUsers)

    router.Run("localhost:8080")
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

// postUsers adds an user from JSON received in the request body.
func postUsers(c *gin.Context) {
    var newUser user

    // Call BindJSON to bind the received JSON to newUser.
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    // Add the new user to the slice.
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}
