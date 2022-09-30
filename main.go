package main

import (
	"simpleQueue/service"

	"github.com/gin-gonic/gin"
)

func main() {
	queueList := make(map[string][]string)

	r := gin.Default()

	service.InitQueueService(queueList, r)

	r.Run(":8080")
}