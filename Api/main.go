package main

import (
	"fmt"
	// "net/http"
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
	fmt.Println(users)
	/* http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my API")
	})) */
}
