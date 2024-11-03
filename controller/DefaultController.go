package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DefaultController struct {
}

func (dc DefaultController) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"ti": "ok"})
}

func (dc DefaultController) Status(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now().UnixMilli(),
		"msg":       "Program Is Running",
	})
}
