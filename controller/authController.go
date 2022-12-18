package controller

import (
	"casbinginxorm/models"
	"casbinginxorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// CreateToken godoc
//
//	@Summary		CreateToken
//	@Description	CreateToken
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			createtokenrequest	body   models.CreateTokenRequest  true  "CreateToken"
//	@Success		200		{string}    string "OK"
//	@Router			/api/v1/auth/createtoken [post]
func (r *AuthController) CreateToken(c *gin.Context) {
	data := make(map[string]interface{})
	code := utils.INVALID_PARAMS
	var createTokenRequest models.CreateTokenRequest
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&createTokenRequest); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if createTokenRequest.Username != "admin" || createTokenRequest.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}

	jwt := utils.Jwt{}
	token, err := jwt.CreateToken(createTokenRequest)
	if err != nil {
		code = utils.ERROR
	} else {
		data["token"] = token
		code = utils.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMsg(code),
		"data": data,
	})
}

// ValidToken godoc
//
// @Summary ValidToken
// @Tags auth
// @Description ValidToken
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string "OK"
// @Router /api/v1/auth/validtoken/{token} [get]
func (r *AuthController) ValidToken(c *gin.Context) {
	data := make(map[string]interface{})
	code := utils.INVALID_PARAMS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMsg(code),
		"data": data,
	})
}

// CreateRefreshToken godoc
//
// @Summary CreateRefreshToken
// @Tags auth
// @Description CreateRefreshToken
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /api/v1/auth/createrefreshtoken [post]
func (r *AuthController) CreateRefreshToken(c *gin.Context) {
	data := make(map[string]interface{})
	code := utils.INVALID_PARAMS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMsg(code),
		"data": data,
	})
}

// ValidRefreshToken godoc
//
// @Summary ValidRefreshToken
// @Tags auth
// @Description ValidRefreshToken
// @Accept  json
// @Produce  json
// @Param refreshtoken query string true "refreshtoken"
// @Success 200 {string} string "OK"
// @Router /api/v1/auth/validrefreshtoken/{refreshtoken} [get]
func (r *AuthController) ValidRefreshToken(c *gin.Context) {
	data := make(map[string]interface{})
	code := utils.INVALID_PARAMS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMsg(code),
		"data": data,
	})
}
