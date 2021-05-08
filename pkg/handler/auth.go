package handler

import (
	"github.com/gin-gonic/gin"
	pack "go-serv-arduino"
	"net/http"
)

func (h* Handler) signUp(c *gin.Context) {
	var input pack.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"res": "Вы успешно зарегистрировали аккаунт!",
	})

}

func (h* Handler) signIn(c *gin.Context) {
	var input pack.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"res": "OK",
	})

}
