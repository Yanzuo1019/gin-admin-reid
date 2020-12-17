package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var Config map[string]interface{}

func ReadConfig(configFileName string) {
	cfg, err := ini.Load(configFileName)
	if err != nil {
		log.Panic("failed to open config file: ", err)
	}

	username := cfg.Section("Administrator").Key("username").String()
	password := cfg.Section("Administrator").Key("password").String()
	admin := make(map[string]string)
	admin["username"] = username
	admin["password"] = password

	addr := cfg.Section("Network").Key("addr").String()
	port := cfg.Section("Network").Key("port").String()
	network := make(map[string]string)
	network["addr"] = addr
	network["port"] = port

	hbaddr := cfg.Section("HeartBeatServer").Key("addr").String()
	hbport := cfg.Section("HeartBeatServer").Key("port").String()
	hbserver := make(map[string]string)
	hbserver["addr"] = hbaddr
	hbserver["port"] = hbport


	Config = make(map[string]interface{})
	Config["admin"] = admin
	Config["network"] = network
	Config["hbserver"] = hbserver

	log.Println("open config file successfully")
}