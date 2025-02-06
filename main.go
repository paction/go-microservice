package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /api/time-in-halifax
func getCurrentTimeInHalifax(c *gin.Context) {
	loc, err := time.LoadLocation("America/Halifax")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load timezone"})
		return
	}

	currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
	c.JSON(http.StatusOK, gin.H{"time_in_halifax": currentTime})
}

// POST /api/content
func postContent(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomValue := rng.Intn(2) // 0 or 1

	// Random success/failure
	if randomValue == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Content processed successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": "Something went wrong"})
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/time-in-halifax", getCurrentTimeInHalifax)
	r.POST("/api/content", postContent)

	// Start server on port 8095
	log.Println("Server running on http://localhost:8095")
	if err := r.Run(":8095"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
