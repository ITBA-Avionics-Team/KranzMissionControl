package controller

import (
	"encoding/json"
	"fmt"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewCommunicationModuleRouter(systemStatusBroadcast broadcast.Broadcast[model.SystemStatus]) *gin.Engine {
	router := gin.Default()
	router.GET("/system_status", SystemStatusHandler(systemStatusBroadcast, wsUpgrader))
	return router
}

func SystemStatusHandler(systemStatusBroadcast broadcast.Broadcast[model.SystemStatus], wsUpgrader websocket.Upgrader) gin.HandlerFunc{
	return func(c *gin.Context) {
		conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		systemStatusSubscriptionId, systemStatusSubscription := systemStatusBroadcast.Subscribe()
		defer systemStatusBroadcast.Unubscribe(systemStatusSubscriptionId)
		for {
			select {
			case systemStatusMessage := <- systemStatusSubscription:
				fmt.Printf("[Controller] Sending message: %+v\n", systemStatusMessage)
				jsonData, err := json.Marshal(systemStatusMessage)
				if err != nil {
					log.Fatalf("Error marshaling to JSON: %v", err)
				}
				if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Printf("WebSocket write error: %v", err)
					return 
			}
				break
			}
		}
	}
}