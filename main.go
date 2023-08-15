package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"httpserverdummy/model"
	"net/http"
)

func main() {

	/*	forkPtr := flag.Bool("fork", false, "a bool")*/
	host := flag.String("h", "", "主机名")
	port := flag.Int("p", 80, "目标端口号，默认80")

	var listenAddr string
	if *host != "" {
		listenAddr = fmt.Sprintf("%s:%d", host, port)
	} else {
		listenAddr = fmt.Sprintf("0.0.0.0:%d", port)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", creaeAlbums)

	router.Run(listenAddr)

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func creaeAlbums(c *gin.Context) {
	data, _ := c.GetRawData()
	c.JSON(http.StatusOK, data)
}
