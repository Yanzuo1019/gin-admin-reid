package main

import (
	"fmt"
	"gin-admin-reid/router"
	"gin-admin-reid/service"
	"gin-admin-reid/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", 0755)
	}
	logFile, err := os.OpenFile("log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("failed to open log file: ", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Println("set default logger successfully")

	_, err = os.Stat("files")
	if err != nil {
		os.Mkdir("files", 0755)
	}

	utils.ReadConfig("config/config.ini")

	service.HeartBeatServerCrash = make(chan int, 1)
	go service.HeartBeatServer()
	go func() {
		for {
			select {
			case <-service.HeartBeatServerCrash:
				log.Println("HeartBeatServer reboot")
				time.Sleep(5 * time.Second)
				go service.HeartBeatServer()
			}
		}
	}()

	router.InitRouter()
}