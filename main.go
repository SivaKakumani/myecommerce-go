package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// https://pkg.go.dev/github.com/gin-gonic/gin

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kvsp_paul")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Initialize the Gin router
	router := gin.Default()

	router.Use(cors.Default())

	// Define route handlers
	router.GET("/", index)

	router.GET("/api/authentication", authentication)

	router.POST("/signup", submitForm)

	// Start the server
	router.Run(":8082")
}

// Handler for the root route
func index(c *gin.Context) {
	c.String(200, "Success")
}

// Handler for the /api/authentication route
func authentication(c *gin.Context) {
	data := map[string]interface{}{
		"Modules": 15,
		"Subject": "Data Structures and Algorithms",
	}
	c.JSON(200, data)
}
func submitForm(c *gin.Context) {
	// Bind JSON data to a struct
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON data")
		return
	}

	// Access user data
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)

	// Insert data into the database
	_, err := db.Exec("INSERT INTO pygo (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error inserting data into the database")
		return
	}

	c.String(http.StatusOK, "Success")
}
