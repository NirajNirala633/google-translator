package controller

import (
	"google-translator/helper"
	"google-translator/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Translate(c *gin.Context) {
	var translateObj model.Translate
	err := c.BindJSON(&translateObj)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request can't fulfil",
		})
		return
	}
	resp, err := helper.Translate(translateObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "request can't fulfil by by google",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    string(resp),
	})
}
