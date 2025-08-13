package handler

import (
	"github.com/BOBAvov/sub_track"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) totalSum(c *gin.Context) {
	var input sub_track.SumResponse
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//if sum, err := h.services.totalSum

}
