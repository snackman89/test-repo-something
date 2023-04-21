package journey

// @TODO build out struct
// type journey struct {
// 	JOURNEYURLNICKNAME string `json:"Journey URL Nickname"`
// PRIMARYITEM string `json:"Primary Item"`
// JOURNEYNAME string `json:"Journey Name"`
// WMCJOURNEYURL string `json:"WMC Journey URL"`
// big idea Banner Image string `json:"Big idea "`
// item Category
// item Franchise
// Journey Publish Status
// Journey Published Date
// Journey Unpublished Date
// Journey Tags
// }

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJourneyFromContentful(c *gin.Context) {

	id := c.Param("id")

	// access token is read only
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/yiil34ggxxe1/environments/client-catalog-dev/entries/%v?access_token=%v", id, "RqfHs1di8kEu9pqZCQF2_m0Ehmlv2zob3QvgqeEN-gQ")

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)

	if error != nil {
		log.Fatal(err)
	}

	c.Data(200, "application/json", body)
}
