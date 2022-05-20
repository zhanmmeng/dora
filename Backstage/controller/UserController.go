package controller

import (
	"dora/Backstage/common"
	"dora/Backstage/model"
	"dora/Backstage/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

	//密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"加密错误"})
		return
	}

	//创建
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	c.JSON(200, gin.H{"code": 200, "msg": "注册成功"})
}

func Login(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
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

	//手机号是否存在
	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "用户不存在"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "msg": "密码错误"})
		return
	}

	//发放token
	token := "11"

	c.JSON(200, gin.H{"code": 200, "msg": "登录成功","data":gin.H{"token":token}})
}

func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
