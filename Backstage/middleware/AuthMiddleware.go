package middleware

import (
	"dora/Backstage/common"
	"dora/Backstage/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "validate权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token解析出错"})
			c.Abort()
			return
		}

		//通过验证，获取claims 的userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户不存在，token无效
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			c.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
