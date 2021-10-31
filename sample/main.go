package main

import (
	"log"
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

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getHeaders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, c.Request.Header)
}

func middleware(c *gin.Context) {
	log.Printf("header of X-Auth-Token: %s", c.GetHeader("X-Auth-Token"))

	c.Next()
}

func main() {
	router := gin.Default()
	router.Use(middleware)

	router.GET("/albums", getAlbums)
	router.GET("/", getHeaders)

	router.Run("0.0.0.0:9001")
}
