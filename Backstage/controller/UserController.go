package controller

import (
	"dora/Backstage/common"
	"dora/Backstage/dto"
	"dora/Backstage/model"
	"dora/Backstage/response"
	"dora/Backstage/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	_ "math/rand"
	"net/http"
	_ "time"
)

// Register 注册
func Register(c *gin.Context) {

	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	//数据验证
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "手机号必须为11位", nil)
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, "密码不能少于6位", nil)
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExist(DB, phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, "该手机号已注册", nil)
		return
	}

	//密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, "加密错误", nil)
		return
	}

	//创建
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)

	response.Success(c, "注册成功", nil)
}

// Login 登录
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "msg": "密码错误"})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "masg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	response.Success(c, "登录成功", gin.H{"token": token})
}

// AuthInfo 获取用户信息
func AuthInfo(c *gin.Context) {
	user, _ := c.Get("user")

	response.Success(c, "success", gin.H{"user_info": dto.ToUserDTO(user.(model.User))})
}

//验证电话号码是否存在
func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
