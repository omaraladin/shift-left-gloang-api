package main

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}{
    {ID: "1", Title: "Master of Puppets", Artist: "Metallica", Price: 30.00},
    {ID: "2", Title: "All Hope is gone", Artist: "Slipknot", Price: 30.00},
    {ID: "3", Title: "Multitude", Artist: "Stromae", Price: 39.99},
    {ID: "4", Title: "Mutter", Artist: "Rammstein", Price: 39.99},
}

func main() {
    r := gin.Default()
    r.GET("/albums", func(c *gin.Context) {
        c.JSON(http.StatusOK, albums)
    })

    addr := listenAddress()
    r.Run(addr)
}

// listenAddress picks a listen address based on APP_ENV and optional PORT.
// If PORT is set it will be used; otherwise Dev -> 8080, Staging -> 8081, Prod -> 8085.
func listenAddress() string {
    if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
        if strings.Contains(p, ":") {
            return p
        }
        return ":" + p
    }
    env := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))
    switch env {
    case "dev", "development":
        return "localhost:8080"
    case "staging":
        return "localhost:8081"
    case "prod", "production":
        return "0.0.0.0:8085"
    default:
        return "localhost:8080"
    }
}