package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h* Handler) getNews(c *gin.Context) {
	news, err := h.services.News.GetNews()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"news": news,
		"count": len(news),
	})
}