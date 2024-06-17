package api

import (
	"github.com/gin-gonic/gin"
	"guangfa-fund/internal/svc/sso"
)

const (
	pathPrefix = "/guangfaFund"
)

func Router(engine *gin.Engine) {
	//SSO
	ssoGroup := engine.Group(pathPrefix + "/sso")
	{
		//获取授权code
		ssoGroup.GET("/authorize", sso.Authorize)
		//获取令牌token
		ssoGroup.GET("/token", sso.Token)
		//获取用户信息
		ssoGroup.GET("/userInfo", sso.UserInfo)
		//回调地址
		ssoGroup.GET("/callback", sso.Callback)
	}
}
