package api

import (
	"gin-admin-reid/service"
	"github.com/gin-gonic/gin"
)

func DeviceStatus(c *gin.Context) {
	var dev []map[string]interface{}

	service.RecordLock.Lock()
	for key := range service.Record {
		if service.Record[key].Status == service.On2Off {
			service.Record[key].Status = service.Offline
		} else if service.Record[key].Status == service.Off2On {
			service.Record[key].Status = service.Online
		}

		devItem := gin.H{
			"id":      service.Record[key].Id,
			"address": key,
			"mac":     service.Record[key].Mac,
			"status":  service.Record[key].Status,
		}

		dev = append(dev, devItem)
	}
	service.RecordLock.Unlock()

	dev = append(dev, gin.H{
		"id": 0,
		"address": "172.19.63.106",
		"mac": "90-78-41-1B-8B-13",
		"status": "Offline",
	})

	c.JSON(200, gin.H{
		"status": 200,
		"error_msg": "",
		"data": gin.H{
			"dev": dev,
		},
	})
}
