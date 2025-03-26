package handlers

import (
	"net/http"
	"zukiapi/services"

	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Prompt  string `json:"prompt"`
	Model   string `json:"model"`
	Size    string `json:"size"`
	Quality string `json:"quality"`
	N       int    `json:"n"`
}

func ImageHandler(c *gin.Context) {
	var req ImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	resp, err := services.ImageCall(req.Model, req.Prompt, req.Size, req.Quality, req.N)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "image request failed"})
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}
