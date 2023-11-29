package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve the HTML page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Website Status Checker",
		})
	})

	// Handle form submission
	r.POST("/check", func(c *gin.Context) {
		websiteURL := c.PostForm("url") // Get URL from the form

		// Prepare a JSON response
		responseMessage := gin.H{
			"status":  "success",
			"message": "Website is up and running!",
		}

		if !checkWebsite(websiteURL) {
			responseMessage["status"] = "error"
			responseMessage["message"] = "Website is down!"
			c.JSON(http.StatusInternalServerError, responseMessage) // Send JSON response
			return
		}

		c.JSON(http.StatusOK, responseMessage) // Send JSON response
	})
	r.LoadHTMLGlob("templates/*")
	r.Run(":8080")
}

func checkWebsite(url string) bool {
	// Add protocol if missing
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Custom HTTP client with a timeout
	client := &http.Client{}

	// Making a GET request
	response, err := client.Get(url)
	if err != nil {
		log.Printf("Request failed: %v\n", err)
		return false
	}
	defer response.Body.Close()

	// Log the status code for debugging purposes
	log.Printf("HTTP Status Code: %d\n", response.StatusCode)

	return response.StatusCode == http.StatusOK
}
