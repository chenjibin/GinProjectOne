package controller

import (
	"GinProjectOne/common"
	"GinProjectOne/model"
	"GinProjectOne/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetUserInfo(ctx *gin.Context) {
	u2 := uuid.NewV4()
	log.Printf("generated Version 4 UUID %v", u2)
	response.Success(ctx, gin.H{
		"name": "小孩",
		"phone": "15061997812",
		"address": "顺安新城",
	}, "用户信息获取成功!")
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	err := ctx.ShouldBindJSON(&requestUser)
	if err != nil {
		response.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "操作失败！")
		return
	}
	name := requestUser.Name
	phone := requestUser.Phone
	password := requestUser.Password
	if isTelephoneExist(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经存在")
		return
	}
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name: name,
		Phone: phone,
		Password:  string(bcryptPassword),
	}
	DB.Create(&newUser)
	response.Success(ctx,  gin.H{"token": "123"}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.UserLogin{}
	err := ctx.ShouldBindJSON(&requestUser)
	if err != nil {
		response.Fail(ctx, gin.H{
			"error": err.Error(),
		}, "操作失败！")
		return
	}
	var user model.User
	DB.Where("phone = ?", requestUser.Phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)) ; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	response.Success(ctx,  gin.H{"token": "123"}, "登录成功!")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	result := db.Where("phone = ?", telephone).First(&user)
	if result.RowsAffected != 0 {
		return true
	}
	return false
}