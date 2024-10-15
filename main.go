package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Google SSO login endpoint
	router.GET("/auth/google", func(c *gin.Context) {
		// Redirect to Google SSO
		c.Redirect(http.StatusFound, "https://accounts.google.com/o/oauth2/auth")
	})

	// Google OAuth2 callback endpoint
	router.GET("/auth/google/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code is required"})
			return
		}
		// Handle the OAuth2 callback logic here
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully logged in with Google",
			"token":   "example_token",
			"user_id": "example_user_id",
		})
	})

	// Get user profile
	router.GET("/users/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		// Fetch user profile logic here
		c.JSON(http.StatusOK, gin.H{
			"id":              userId,
			"username":        "example_username",
			"email":           "example_email",
			"derbies_created": []string{"derby1", "derby2"},
			"derbies_joined":  []string{"derby3", "derby4"},
		})
	})

	// Create a new derby
	router.POST("/derbies", func(c *gin.Context) {
		var derby struct {
			Name      string `json:"name" binding:"required"`
			Location  string `json:"location" binding:"required"`
			StartTime string `json:"start_time" binding:"required"`
			EndTime   string `json:"end_time" binding:"required"`
			Rules     string `json:"rules"`
		}
		if err := c.ShouldBindJSON(&derby); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Create derby logic here
		c.JSON(http.StatusOK, gin.H{
			"message":  "Derby created successfully",
			"derby_id": "example_derby_id",
		})
	})

	// Start the server
	router.Run(":8080")
}
