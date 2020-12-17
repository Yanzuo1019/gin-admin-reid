package api

import (
	"gin-admin-reid/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func AuthCheck(c *gin.Context) {
	token := c.Request.Header.Get("authorization")

	if token == "" {
		log.Println("empty token")
		c.JSON(500, gin.H{
			"status":    500,
			"error_msg": "empty token",
			"data": gin.H{
				"valid": false,
			},
		})
	} else {
		claims, err := utils.ParseToken(token)
		if err != nil {
			log.Println("fail to parse token: ", err)
			c.JSON(500, gin.H{
				"status":    500,
				"error_msg": "fail to parse token",
				"data": gin.H{
					"valid": false,
				},
			})
		} else if time.Now().Unix() > claims.ExpiresAt {
			log.Println("token expired")
			c.JSON(200, gin.H{
				"status":    200,
				"error_msg": "token expired",
				"data": gin.H{
					"valid": false,
				},
			})
		} else {
			log.Println("authorization check successfully")
			c.JSON(200, gin.H{
				"status":    200,
				"error_msg": "",
				"data": gin.H{
					"valid": true,
				},
			})
		}
	}
}
