package controller

import (
	"Go-000/Week02/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	var (
		res = gin.H{}
	)
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res["message"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		// 处理error
		res["message"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res["data"] = user
	res["success"] = true
	c.JSON(http.StatusOK, res)
}