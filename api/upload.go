package api

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func Upload(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")

	if err != nil {
		log.Println("fail to get formfile")
		c.JSON(500, gin.H{
			"status":    500,
			"error_msg": "fail to get formfile",
		})
	} else {
		filename := handler.Filename
		log.Println("receive file ", filename)
		out, err := os.OpenFile("files/" + filename, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Println("fail to create file")
			c.JSON(500, gin.H{
				"status":    500,
				"error_msg": "fail to create file",
			})
		} else {
			defer out.Close()
			_, err := io.Copy(out, file)
			if err != nil {
				log.Println("fail to copy file")
				c.JSON(500, gin.H{
					"status": 500,
					"error_msg": "fail to copy file",
				})
			} else {
				log.Println("process file successfully")
				c.JSON(200, gin.H{
					"status": 200,
					"error_msg": "",
				})
			}
		}
	}
}
