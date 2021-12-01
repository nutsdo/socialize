package qq

import "github.com/nutsdo/socialize"

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"` //仅微信
	OpenId       string `json:"open_id"`       // 仅微信
	Scope        string `json:"scope"`         //仅微信
}

func (r *AccessTokenResponse) GetAccessToken() *socialize.AccessTokenDetail {
	return &socialize.AccessTokenDetail{
		AccessToken:  r.AccessToken,
		ExpiresIn:    r.ExpiresIn,
		RefreshToken: r.RefreshToken,
		OpenId:       r.OpenId,
		Scope:        r.Scope,
	}
}

type UnionIDResponse struct {
	ClientId string `json:"client_id"`
	Openid   string `json:"openid"`
	Unionid  string `json:"unionid"`
}

type QQResponse struct {
	Ret          int    `json:"ret"`
	Msg          string `json:"msg"`
}

type UserResponse struct {
	QQResponse
	Nickname     string `json:"nickname"`
	Figureurl    string `json:"figureurl"`
	Figureurl1   string `json:"figureurl_1"`
	Figureurl2   string `json:"figureurl_2"`
	FigureurlQq1 string `json:"figureurl_qq_1"`
	FigureurlQq2 string `json:"figureurl_qq_2"`
	Gender       string `json:"gender"`
	Province      string `json:"province"`
	City          string `json:"city"`
	FigureurlQq   string `json:"figureurl_qq"`
	FigureurlType string `json:"figureurl_type"`
}

type QQUser struct {
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 40*40头像
	Gender   int    `json:"gender"`   // 0未知，1男，2女
	Openid   string `json:"openid"`   // qq开放平台应用唯一ID
	Unionid  string `json:"unionid"`  // unionid qq开放平台唯一ID
	Province      string `json:"province"`
	City          string `json:"city"`
}

type UserInfoResponse struct {
	Nickname string `json:"nickname"`
}

func (r *UserInfoResponse) GetUser() *socialize.UserDetail {
	return &socialize.UserDetail{Nickname: r.Nickname}
}


