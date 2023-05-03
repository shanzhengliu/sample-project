package controllers

import (
	"fmt"
	database "modules/gin/Databases"

	"github.com/gin-gonic/gin"
)

func TestInsert(c *gin.Context) {
	fmt.Println("test insert")
	database.DB.Create(&database.User{Name: "test", Age: 18})
	sample_channel := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			sample_channel <- i
		}
		close(sample_channel)
	}()

	for i := range sample_channel {
		// fmt.Println(i)

		fmt.Println(i)

	}
	c.JSON(200, gin.H{
		"message": "test insert",
	})
}
