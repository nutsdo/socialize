package wechat

import (
	"github.com/nutsdo/socialize"
)

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
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func (r *UnionIDResponse) GetUser() *socialize.UserDetail {
	return &socialize.UserDetail{
		Openid:     r.Openid,
		Unionid:    r.Unionid,
		Nickname:   r.Nickname,
		Gender:     r.Sex,
		Province:   r.Province,
		City:       r.City,
		Country:    r.Country,
		Headimgurl: r.Headimgurl,
		Privilege:  r.Privilege,
	}
}
