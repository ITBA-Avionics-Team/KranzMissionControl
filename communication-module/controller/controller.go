package controller

import (
	"encoding/json"
	"fmt"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
	"log"
	"net/http"
	"sync"

	"slices"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	extSerial "go.bug.st/serial" // TODO: find a prettier way to do this
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		acceptedOrigins := []string{"http://localhost:5173"}
		return slices.Contains(acceptedOrigins, r.Header.Get("Origin"))
	},
}

func NewCommunicationModuleRouter(systemStatusBroadcast broadcast.Broadcast[model.SystemStatus], serialPort extSerial.Port, serialPortMutex *sync.Mutex) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/system_status", GetSystemStatusHandler(systemStatusBroadcast, wsUpgrader))
	router.POST("/command", GetCommandHandler(serialPort, serialPortMutex))
	return router
}

func GetSystemStatusHandler(systemStatusBroadcast broadcast.Broadcast[model.SystemStatus], wsUpgrader websocket.Upgrader) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer conn.Close()
		systemStatusSubscriptionId, systemStatusSubscription := systemStatusBroadcast.Subscribe()
		defer systemStatusBroadcast.Unubscribe(systemStatusSubscriptionId)
		for {
			select {
			case systemStatusMessage := <-systemStatusSubscription:
				jsonData, err := json.Marshal(systemStatusMessage)
				fmt.Printf("[Controller] Sending message: %s\n", jsonData)
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

func GetCommandHandler(serialPort extSerial.Port, mutex *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		var command model.Command

		// Bind the received JSON to newUser.
		if err := c.ShouldBindJSON(&command); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Process the newUser data (e.g., save to database, etc.)
		fmt.Printf("Received command: %v", command)
		serial.SendCommand(serialPort, mutex, command)

		// Send a response
		c.JSON(http.StatusOK, gin.H{
			"message": "Command sent to LC",
			"command": command,
		})
	}
}
