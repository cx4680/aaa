package sso

type Request struct {
	AppId       string `form:"appId"`
	RedirectUri string `form:"redirectUri"`
	State       string `form:"state"`
}

type AuthorizeRequest struct {
	ClientId     string `json:"client_id"`     //客户端应用注册ID，从认证中心获取
	ResponseType string `json:"response_type"` //code，固定值
	RedirectUri  string `json:"redirect_uri"`  //成功授权后的回调地址，需要urlencode
}

type AuthorizeResponse struct {
	Msg    string `json:"msg"`  //具体的错误信息
	Code   string `json:"code"` //具体的错误码
	Status int    `json:"status"`
}

type GetTokenRequest struct {
	ClientId     string `json:"client_id"`     //客户端应用注册ID，从认证中心获取
	ClientSecret string `json:"client_secret"` //客户端应用注册密钥,从认证中心获取
	Code         string `json:"code"`          //调用认证接口（authorize）获得ticket
	GrantType    string `json:"grant_type"`    //请求类型，默认authorization_code
	RedirectUri  string `json:"redirect_uri"`  //且必须与调用authorize中的该参数值保持一致
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"` //状态码：获取的授权码token值
	Expires     int    `json:"expires"`      //授权码有效时间（秒）：授权服务器返回给应用的访问票据的有效期。注意：有效期以秒为单位。
	Msg         string `json:"msg"`          //提示信息：查看错误代码表
	Code        string `json:"code"`         //错误代码：查看错误代码表
	Status      int    `json:"status"`       //状态码：成功200，失败400
}

type getUserInfoResponse struct {
	Attributes struct {
		Id             int    `json:"id"`
		SubCompanyId   int    `json:"subcompanyid"`
		WorkCode       string `json:"workcode"`
		Sex            string `json:"sex"`
		DepartmentId   int    `json:"departmentid"`
		Mobile         string `json:"mobile"`
		SystemLanguage int    `json:"systemlanguage"`
		Telephone      string `json:"telephone"`
		ManagerId      int    `json:"managerid"`
		CountryId      int    `json:"countryid"`
		Lastname       string `json:"lastname"`
		AssistantId    int    `json:"assistantid"`
		RelatedAccount string `json:"relatedAccount"`
		CertificateNum string `json:"certificatenum"`
		Email          string `json:"email"`
		Status         int    `json:"status"`
	} `json:"attributes"`
	Id     string `json:"id"`
	Msg    string `json:"msg"`
	Code   string `json:"code"`
	Status int    `json:"status"`
}
