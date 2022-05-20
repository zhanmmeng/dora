package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 200, msg, data)
}

func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 400, msg, data)
}
