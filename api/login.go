package api

import (
	"gin-admin-reid/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var request LoginReq

	admin, _ := utils.Config["admin"].(map[string]string)

	if err := c.BindJSON(&request); err == nil {
		if request.Username == admin["username"] && request.Password == admin["password"] {
			token, err := utils.GenerateToken(request.Username, request.Password)
			if err != nil {
				log.Println("fail to generate token: ", err)
				c.JSON(500, gin.H{
					"status": 500,
					"error_msg": "fail to generate token",
					"data": gin.H{
						"token": "",
					},
				})
			} else {
				log.Println("generate token successfully")
				c.JSON(200, gin.H{
					"status": 200,
					"error_msg": "",
					"data": gin.H{
						"token": token,
					},
				})
			}
		} else {
			log.Println("invalid username and password")
			c.JSON(500, gin.H{
				"status": 500,
				"error_msg": "invalid username and password",
				"data": gin.H{
					"token": "",
				},
			})
		}
	} else {
		log.Println("fail to execute BindJson: ", err)
		c.JSON(500, gin.H{
			"status": 500,
			"error_msg": "invalid username and password",
			"data": gin.H{
				"token": "",
			},
		})
	}
}