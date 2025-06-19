package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)


var ctx = context.Background() // Create a background context for Redis operations
var rdb *redis.Client // Declare a global variable for the Redis client

func main() {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // default Redis port
	})

	// Test Redis connection
	_, err := rdb.Ping(ctx).Result() // Ping the Redis server to check if it's reachable
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
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
  c.JSON(200, gin.H{ "message": "Hello, World!" })
	 })
	// POST route to shorten URLs
	router.POST("/shorten", func(c *gin.Context) {
		var requestBody struct {
			URL string `json:"url"`
		}

		// Bind JSON input to requestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		// Generate random short code
		shortCode := generateShortCode(6)

		// Save mapping to Redis
		err := rdb.Set(ctx, shortCode, requestBody.URL, 24*time.Hour).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save to Redis"})
			return
		}

		// Return the shortened URL
		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
		c.JSON(200, gin.H{"short_url": shortURL})
	})

	router.GET("/:shortcode", func(c *gin.Context) {
		shortCode := c.Param("shortcode")
	
		// Look up the original URL from Redis
		originalURL, err := rdb.Get(ctx, shortCode).Result()
		if err == redis.Nil {
			// Key not found
			c.JSON(404, gin.H{"error": "Short URL not found"})
			return
		} else if err != nil {
			// Redis error
			c.JSON(500, gin.H{"error": "Failed to query Redis"})
			return
		}
	
		// Redirect to the original URL
		c.Redirect(302, originalURL)
	})
	
	router.Run(":8080") // Start the server on port 8080
}

func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	return string(b)
}