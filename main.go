package main

import (
    "crypto/md5"
    "crypto/tls"
    "database/sql"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "os/exec"
    "time"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

// Hardcoded configuration values (for demo purposes only)
const apiKey = "sk_live_DEMO_123456"
const dbUser = "demo_user"
const dbPass = "P@ssword!23"
const dbDSN  = dbUser + ":" + dbPass + "@tcp(localhost:3306)/demo"

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func main() {
    // Predictable randomness (SAST should flag)
    rand.Seed(42)

    router := gin.Default()
    router.GET("/albums", getAlbums)

    // SQL injection endpoint
    router.GET("/search", searchAlbums) // ?q=Metallica' OR '1'='1

    // Command injection endpoint
    router.GET("/ping", pingHost) // ?host=127.0.0.1 ; rm -rf /

    // Path traversal endpoint
    router.GET("/file", readFile) // ?name=../main.go

    // Weak password hashing endpoint
    router.POST("/hash", hashPassword)

    // Insecure TLS client endpoint
    router.GET("/fetch", insecureFetch) // ?url=https://expired.badssl.com/

	// Run the server
    router.Run("localhost:8080")
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Master of Puppets", Artist: "Metallica", Price: 30.00},
    {ID: "2", Title: "All Hope is gone", Artist: "Slipknot", Price: 30.00},
    {ID: "3", Title: "Multitude", Artist: "Stromae", Price: 39.99},
    {ID: "4", Title: "Mutter", Artist: "Rammstein", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// SQL injection: unsafe string concat with user input
func searchAlbums(c *gin.Context) {
    q := c.Query("q")
    db, err := sql.Open("mysql", dbDSN)
    if err != nil {
        c.String(http.StatusInternalServerError, "DB error")
        return
    }
    defer db.Close()

    // Vulnerable query (use prepared statements in real code)
    query := fmt.Sprintf("SELECT id, title, artist, price FROM albums WHERE title LIKE '%%%s%%' OR artist LIKE '%%%s%%'", q, q)
    rows, err := db.Query(query)
    if err != nil {
        c.String(http.StatusInternalServerError, "Query error")
        return
    }
    defer rows.Close()

    var results []album
    for rows.Next() {
        var a album
        if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
            continue
        }
        results = append(results, a)
    }
    c.JSON(http.StatusOK, gin.H{"query": q, "results": results})
}

// Command injection: passes user input to shell
func pingHost(c *gin.Context) {
    host := c.Query("host")
    cmd := exec.Command("sh", "-c", "ping -c 1 "+host)
    out, err := cmd.CombinedOutput()
    if err != nil {
        c.String(http.StatusInternalServerError, "Command error: %v", err)
        return
    }
    c.String(http.StatusOK, string(out))
}

// Path traversal: reads arbitrary files based on input
func readFile(c *gin.Context) {
    name := c.Query("name")
    // Vulnerable: no sanitization or allowlist
    data, err := ioutil.ReadFile("./data/" + name)
    if err != nil {
        c.String(http.StatusNotFound, "Read error")
        return
    }
    c.Data(http.StatusOK, "text/plain", data)
}

// Weak crypto: MD5 for password hashing
func hashPassword(c *gin.Context) {
    var body struct {
        Password string `json:"password"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    h := md5.Sum([]byte(body.Password))
    c.JSON(http.StatusOK, gin.H{"md5": hex.EncodeToString(h[:])})
}

// Insecure TLS: skip certificate verification
func insecureFetch(c *gin.Context) {
    url := c.Query("url")
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Vulnerable
    }
    client := &http.Client{
        Transport: tr,
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(url)
    if err != nil {
        c.String(http.StatusBadGateway, "Fetch error")
        return
    }
    defer resp.Body.Close()
    b, _ := ioutil.ReadAll(resp.Body)
    c.Data(http.StatusOK, "text/plain", b)
}