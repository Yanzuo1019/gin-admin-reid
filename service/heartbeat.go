package service

import (
	"encoding/json"
	"gin-admin-reid/utils"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type DevInfo struct {
	Id int
	Mac string
	Status int
}

type HBSignal struct {
	Ip string
	Mac string
}

const (
	Offline int = 0
	Online int = 1
	On2Off int = 2
	Off2On int = 3
)

var Record map[string]*DevInfo
var DevNum int
var RecordLock sync.Mutex
var NumLock sync.Mutex
var HeartBeatServerCrash chan int

func HeartBeatServer() {
	defer func() {
		HeartBeatServerCrash <- 1
	}()

	Record = make(map[string]*DevInfo)
	DevNum = 0

	hbserver, _ := utils.Config["hbserver"].(map[string]string)
	addr := hbserver["addr"]
	port := hbserver["port"]

	listenSocket, err := net.Listen("tcp", addr + ":" + port)
	if err != nil {
		log.Println("fail to create HeartBeatServer listen socket")
		return
	} else{
		log.Println("create HeartBeatServer listen socket successfully")
		for {
			conn, err := listenSocket.Accept()
			if err != nil {
				log.Println("fail to connect ", conn.RemoteAddr().String(), ": ", err)
				continue
			} else {
				log.Println("connect ", conn.RemoteAddr().String(), " successfully")
				go handleConnection(conn)
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(6 * time.Minute))
	for {
		n, err := conn.Read(buffer)
		addr := conn.RemoteAddr().String()
		addr = addr[:strings.Index(addr, ":")]

		if err != nil {
			log.Println("fail to read from ", addr, ": ", err)

			RecordLock.Lock()
			if _, ok := Record[addr]; ok {
				if Record[addr].Status != Offline {
					Record[addr].Status = On2Off
				}
			}
			RecordLock.Unlock()

			log.Println(addr, " offline")
			return
		} else {
			RecordLock.Lock()
			if _, ok := Record[addr]; ok {
				if Record[addr].Status != Online {
					Record[addr].Status = Off2On
				}
			} else {
				data := buffer[:n]
				var info HBSignal
				json.Unmarshal(data, &info)

				NumLock.Lock()
				Record[addr] = &DevInfo{
					Id:     DevNum,
					Mac:    info.Mac,
					Status: Off2On,
				}
				DevNum += 1
				NumLock.Unlock()

				log.Println("Ip:", info.Ip, " Mac:", info.Mac, " connected")
			}
			RecordLock.Unlock()
			conn.SetDeadline(time.Now().Add(6 * time.Minute))
		}
	}
}