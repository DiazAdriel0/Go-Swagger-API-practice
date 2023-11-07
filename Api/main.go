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

    router.Run("localhost:8080")
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}


