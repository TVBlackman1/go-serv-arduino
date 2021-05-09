package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h* Handler) getDevices(c *gin.Context) {
	devices := h.services.Device.GetDevices()

	c.JSON(http.StatusOK, map[string]interface{}{
		"devices": devices,
		"count": len(devices),
	})
}