package controller

import (
	"dora/Backstage/common"
	"dora/Backstage/model"
	"dora/Backstage/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "math/rand"
	"net/http"
	_ "time"
)

func Register(c *gin.Context) {

	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	//数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "密码不能少于6位"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(DB, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "该手机号已注册"})
		return
	}

	//创建
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	DB.Create(&newUser)

	c.JSON(200, gin.H{"msg": "注册成功"})
}



func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
