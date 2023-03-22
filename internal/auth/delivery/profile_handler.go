package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api UserAuthHandler) getUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	return
}
