package service

import (
	"GolangStudy/blog/common/config"
	"GolangStudy/blog/common/middleware"
	"GolangStudy/blog/common/result"
	UserDto "GolangStudy/blog/user/model/dto"
	UserEntity "GolangStudy/blog/user/model/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var regUser UserDto.RegUser
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrParam)
		return
	}
	// 检查用户是否已存在
	var checkUser UserEntity.User
	if err := config.DB.Where("username = ?", regUser.User).First(&checkUser); err == nil {
		c.JSON(http.StatusBadRequest, result.ErrUserNameExists)
		return
	}
	// 检查邮箱是否已存在
	if err := config.DB.Where("email = ?", regUser.Email).First(&checkUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, result.ErrUserEmailExists)
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	regUser.Password = string(hashedPassword)

	user := UserEntity.User{
		UserName: regUser.User,
		Password: string(hashedPassword),
		Email:    regUser.Email,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, result.ErrUserCreate)
		return
	}
	c.JSON(http.StatusOK, result.Success(user))
}

func Login(c *gin.Context) {
	var user UserDto.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrParam)
		return
	}
	var checkUser UserEntity.User
	if err := config.DB.Where("username = ?", user.User).Take(&checkUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, result.ErrUserQuery)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, result.ErrUserOrPwd)
		return
	}
	claims := middleware.Myclaims{
		ID:       checkUser.ID,
		Username: checkUser.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * config.Conf.Jwt.TokenExpire)),
			Issuer:    config.Conf.Jwt.Issuer,
		},
	}
	token, err := middleware.CreateJwt(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, result.Success(gin.H{"token": token}))
}
