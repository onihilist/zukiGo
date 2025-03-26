package handlers

import (
	"net/http"
	"zukiapi/services"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Messages []map[string]string `json:"messages"`
	Model    string              `json:"model"`
}

func ChatHandler(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	resp, err := services.ChatCall(req.Messages, req.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "chat request failed"})
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}
