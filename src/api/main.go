package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Restaurant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	restaurant := Restaurant{
		Id: 3,
		Name: "サイゼリヤ",
	}

	r.GET("/restaurant", func(c *gin.Context) {
		c.JSON(200, restaurant)
	})

	bytes, err := ioutil.ReadFile("restaurants.json")
	if err != nil {
		log.Fatal(err)
	}

	var restaurants []Restaurant
	if err := json.Unmarshal(bytes, &restaurants); err != nil {
		log.Fatal(err)
	}

	r.GET("/restaurants", func(c *gin.Context) {
		c.JSON(200, restaurants)
	})

	r.Run(":8080")
}
