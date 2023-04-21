package main

import (
	"fmt"
	"go-blueshift-api/catalog"
	"go-blueshift-api/journey"
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

func getCatalog(c *gin.Context) {

	id := c.Param("id")

	catalog, err := catalog.GetCatalog(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(catalog)

	c.IndentedJSON(http.StatusCreated, catalog)
}

func main() {
	router := gin.Default()

	router.GET("/catalog/:id", getCatalog)

	router.GET("/journey/:id", journey.GetJourneyFromContentful)
	router.Run("localhost:8000")
}
