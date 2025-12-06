package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Master of Puppets", Artist: "Metallica", Price: 30.00},
    {ID: "2", Title: "All Hope is gone", Artist: "Slipknot", Price: 30.00},
    {ID: "3", Title: "Multitude", Artist: "Stromae", Price: 30.00},
	{ID: "4", Title: "Mutter", Artist: "Rammstein", Price: 30.00},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// 6. Hardcoded secrets and debug information
func hardcodedSecrets(c *gin.Context) {
    // Hardcoded credentials (SAST detectable)
    apiKey := "AKIAIOSFODNN7EXAMPLE" // AWS key pattern
    password := "SuperSecret123!"    // Hardcoded password
    jwtSecret := "secretkey"         // Weak JWT secret
    
    // Debug information exposure
    debugInfo := map[string]string{
        "api_key":    apiKey,
        "password":   password,
        "jwt_secret": jwtSecret,
        "version":    "1.0.0-debug",
    }
    
    c.JSON(http.StatusOK, debugInfo)
}