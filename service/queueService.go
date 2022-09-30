package service

import (
	"net/http"
	"simpleQueue/model/request"

	"github.com/gin-gonic/gin"
)

var q map[string][]string

func InitQueueService(queueList map[string][]string, r *gin.Engine) {
	q = queueList
	r.POST("/queue", enqueue)
	r.GET("/dequeue", dequeue)
}

func enqueue(c *gin.Context) {
	var qRequest request.Q

	err := c.BindJSON(&qRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	q[qRequest.Queue] = append(q[qRequest.Queue], qRequest.Message)

	c.JSON(http.StatusCreated, gin.H{"content": q[qRequest.Queue]})
}

func dequeue(c *gin.Context) {
	queueName := c.Query("queueName")

	if len(q[queueName]) == 0 {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}

	qTemp := q[queueName][0]
	q[queueName] = q[queueName][1:]

	c.JSON(http.StatusOK, gin.H{"content": qTemp})
}
