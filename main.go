package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

)

func main() {
	// Create a new Gin router
	// using := to have go take care of the type
	gin.SetMode(gin.DebugMode) // Set Gin to debug mode
	router := gin.Default()
 	// Define a simple GET route
	// Set the trusted platform to allow requests from specific platforms
	router.SetTrustedProxies(nil)
	
	//  func(c *gin.Context).
	//  Here, c is a pointer to Gin’s Context type, which holds the incoming *http.Request, response writer, path parameters, query values, and more.
	//  You use c inside your function to read request data and write responses.

	router.GET("/",func(c *gin.Context){
		  // Respond with a JSON object
  c.JSON(200, gin.H{ "message": "Hello, World!", "status": "success" })
	 })
	err := router.Run(":8080")
	if err != nil {
		panic("########## server error ########\n")
		fmt.Printf("Failed to start server: %v", err)
  // Handle any errors that occur when starting the server
  // The server will listen on port 8080
  // If the port is already in use or there is another issue, it will panic with an error message
	}

}