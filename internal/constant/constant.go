package constant

const (
	//广发基金-异构系统接入OAuth2认证接口
	AuthorizeUrl   = "http://%s/sso/oauth2.0/authorize?client_id=%s&response_type=%s&redirect_uri=%s"
	GetTokenUrl    = "http://%s/sso/oauth2.0/accessToken?client_id=%s&client_secret=%s&grant_type=%s&code=%s&redirect_uri=%s"
	GetUserInfoUrl = "http://%s/sso/oauth2.0/profile?access_token=%s"

	HOST              = "HOST"
	GUANGFA_FUND_HOST = "GUANGFA_FUND_HOST"
	CLIENT_ID         = "CLIENT_ID"
	CLIENT_SECRET     = "CLIENT_SECRET"
)
