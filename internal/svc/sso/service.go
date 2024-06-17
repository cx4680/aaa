package sso

import (
	"fmt"
	"github.com/acmestack/godkits/gox/stringsx"
	"github.com/gin-gonic/gin"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"
	"guangfa-fund/internal/constant"
	"guangfa-fund/internal/util"
	"guangfa-fund/internal/web/result"
	"net/http"
	"net/url"
	"os"
)

func Authorize(c *gin.Context) {
	request := &Request{}
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Errorf("AuthorizationCode error: %v", err)
		result.Failure(c, http.StatusBadRequest, err.Error())
		return
	}
	//将云文档回调地址存入cookie
	c.SetCookie("guangfa_fund_cloud_documents_redirect_uri", request.RedirectUri, 0, "", "", false, true)
	//重定向到客户地址，获取授权码
	host := stringsx.DefaultIfEmpty(os.Getenv(constant.HOST), c.Request.Host)
	//todo 客户地址域名
	//todo ClientId 客户端应用注册ID，从认证中心获取
	callback := fmt.Sprintf("http://%s/guangfaFund/sso/callback", host)
	authorizeUrl := fmt.Sprintf(constant.AuthorizeUrl, os.Getenv(constant.GUANGFA_FUND_HOST), os.Getenv(constant.CLIENT_ID), "code", url.QueryEscape(callback))
	//重定向到客户，请求获取code
	c.Redirect(http.StatusFound, authorizeUrl)
}

func Token(c *gin.Context) {
	code := c.Query("code")
	if util.IsBlank(code) {
		log.Error("code为空")
		result.Failure(c, http.StatusBadRequest, "code为空")
		return
	}
	//调用客户接口，传入临时授权code，返回token
	host := stringsx.DefaultIfEmpty(os.Getenv(constant.HOST), c.Request.Host)
	callback := fmt.Sprintf("http://%s/guangfaFund/sso/callback", host)
	response := &GetTokenResponse{}
	//todo 客户地址域名
	//todo ClientId 客户端应用注册ID，从认证中心获取
	//todo 密钥 客户端应用注册密钥,从认证中心获取
	getTokenUrl := fmt.Sprintf(constant.GetTokenUrl, os.Getenv(constant.GUANGFA_FUND_HOST), os.Getenv(constant.CLIENT_ID), os.Getenv(constant.CLIENT_SECRET), code, "authorization_code", callback)
	if err := util.HttpGet(getTokenUrl, response); err != nil {
		log.Errorf("调用客户接口错误：%v", err)
		result.Failure(c, http.StatusInternalServerError, "调用客户接口错误")
	}
	//从cookie中获取云文档redirectUri
	redirectUri, err := c.Cookie("guangfa_fund_cloud_documents_redirect_uri")
	if err != nil {
		log.Errorf("cookie error: %v", err)
		result.Failure(c, http.StatusInternalServerError, err.Error())
	}
	//重定向云文档接口，传入token
	c.Redirect(http.StatusFound, redirectUri+"?token="+response.AccessToken)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if util.IsBlank(token) {
		log.Error("token为空")
		result.Failure(c, http.StatusBadRequest, "token为空")
		return
	}
	//调用客户接口，传入token
	response := &getUserInfoResponse{}
	//todo 客户地址域名
	getUserInfoUrl := fmt.Sprintf(constant.GetUserInfoUrl, os.Getenv(constant.GUANGFA_FUND_HOST), token)
	if err := util.HttpGet(getUserInfoUrl, response); err != nil {
		log.Errorf("调用客户接口错误：%v", err)
		result.Failure(c, http.StatusInternalServerError, "调用客户接口错误")
	}
	//从cookie中获取云文档redirectUri
	redirectUri, err := c.Cookie("guangfa_fund_cloud_documents_redirect_uri")
	if err != nil {
		log.Errorf("cookie error: %v", err)
		result.Failure(c, http.StatusInternalServerError, err.Error())
	}
	//重定向云文档接口，传入token
	c.Redirect(http.StatusFound, fmt.Sprintf("%s?workcode=%s&lastname=%s", redirectUri, response.Attributes.WorkCode, response.Attributes.Lastname))
}

func Callback(c *gin.Context) {
	//客户回调，获取临时授权码code（ticket）
	code := c.Query("ticket")
	if util.IsBlank(code) {
		log.Error("ticket为空")
		result.Failure(c, http.StatusBadRequest, "ticket为空")
		return
	}

	//从cookie中获取云文档redirectUri
	redirectUri, err := c.Cookie("guangfa_fund_cloud_documents_redirect_uri")
	if err != nil {
		log.Errorf("cookie error: %v", err)
		result.Failure(c, http.StatusInternalServerError, err.Error())
	}
	//重定向云文档接口，把临时授权code（ticket）传过去
	c.Redirect(http.StatusFound, redirectUri+"?code="+code)
}

/*需求逻辑
1、客户请问访问云文档

2、云文档调用sso服务，携带重定向地址redirectUri参数

3、sso服务重定向到客户,把redirectUri参数保存到cookie

4、客户回调sso服务地址，返回code

5、sso服务将code返回给云文档，通过之前携带的重定向地址redirectUri参数

6、云文档携带code请用sso服务

7、sso服务用code请求客户接口获取token，返回给云文档

8、云文档携带token请用sso服务

9、sso服务用token请求客户接口获取用户信息，返回给云文档*/
