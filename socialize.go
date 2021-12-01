package socialize

var socialize LoginFactory

// LoginFactory 第三方登录抽象工厂
type LoginFactory interface {
	CreateLoginClient(appid, secret string, option ...interface{}) Login
}

// Login 登录接口
type Login interface {
	AccessToken(code string) (UnifiedAccessToken, error)
	UserInfo(token *AccessTokenDetail) (UnifiedUser, error)
	UserInfoByCode(code string) (UnifiedUser, error)
}

type UnifiedAccessToken interface {
	GetAccessToken() *AccessTokenDetail
}

type UnifiedUser interface {
	GetUser() *UserDetail
}
