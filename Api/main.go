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
    Age 			int8 	`json:"age"`
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
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)
	router.PUT("/users/:id", updateUser)

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

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of users, looking for
    // an user whose ID value matches the parameter.
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser user

    // Call BindJSON to bind the received JSON to newUser.
    if err := c.BindJSON(&updatedUser); err != nil {
        return
    }


	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			a = updatedUser
			c.IndentedJSON(http.StatusOK, updatedUser)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}