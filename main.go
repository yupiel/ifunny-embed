package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	const baseUrl = "https://ifunny.co/"
	const baseUrlPicture = baseUrl + "picture/"
	const baseUrlVideo = baseUrl + "video/"

	r := gin.Default()

	r.GET("/picture/:id", func(c *gin.Context) {
		c.Redirect(http.StatusFound, getResourceUrl(&ResourceData{
			RequestUrl:           fmt.Sprintf("%s%s", baseUrlPicture, c.Param("id")),
			OriginUrlBase:        "https://imageproxy.ifunny.co/crop:x-20,resize:640x,quality:90x75",
			LinkSearchIndexStart: "/images/",
			LinkSearchIndexEnd:   ".jpg",
		}))
	})

	r.GET("/video/:id", func(c *gin.Context) {
		c.Redirect(http.StatusFound, getResourceUrl(&ResourceData{
			RequestUrl:           fmt.Sprintf("%s%s", baseUrlVideo, c.Param("id")),
			OriginUrlBase:        "https://img.ifunny.co",
			LinkSearchIndexStart: "/videos/",
			LinkSearchIndexEnd:   ".mp4",
		}))
	})

	r.Run()
}

func getResourceUrl(resourceData *ResourceData) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", resourceData.RequestUrl, nil)
	if err != nil {
		log.Println(err)
	}

	//Pretend to be regular browser because request gets rejected otherwise
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error occured when reading body: ", err)
	}
	bodyString := string(body)

	return fmt.Sprintf("%s%s%s",
		resourceData.OriginUrlBase,
		bodyString[strings.Index(bodyString, resourceData.LinkSearchIndexStart):strings.Index(bodyString, resourceData.LinkSearchIndexEnd)],
		resourceData.LinkSearchIndexEnd,
	)
}

type ResourceData struct {
	RequestUrl           string
	OriginUrlBase        string
	LinkSearchIndexStart string
	LinkSearchIndexEnd   string
}
