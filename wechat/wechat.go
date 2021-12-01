package wechat

import (
	"encoding/json"
	"errors"
	"github.com/nutsdo/socialize"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ApiHost         = "https://api.weixin.qq.com/sns/oauth2"
	AccessTokenPath = ApiHost + "/access_token"
	UserInfoPath    = ApiHost + "/userinfo"
	UnionIDPath     = ApiHost + "/userinfo"
)

type LoginFactory struct{}

type Login struct {
	credential *Credential
	client     *http.Client
}

func (lf LoginFactory) CreateLoginClient(appid, secret string, options ...interface{}) socialize.Login {
	return NewLogin(appid, secret)
}

func NewLogin(appid, secret string) socialize.Login {
	//if len(appid) == 0 || len(secret) == 0 {
	//	return nil
	//}
	credential := &Credential{
		appid:  appid,
		secret: secret,
	}
	return &Login{
		credential: credential,
		client: http.DefaultClient}
}

func (l *Login) AccessToken(code string) (socialize.UnifiedAccessToken, error) {
	if len(code) == 0 {
		return nil, errors.New("code is required")
	}

	query := url.Values{}
	query.Add("appid", l.credential.appid)
	query.Add("secret", l.credential.secret)
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")
	dourl := AccessTokenPath + "?" + query.Encode()
	resp, err := l.client.Get(dourl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if isError(bodyBytes) {
		return nil, errors.New(string(bodyBytes))
	}

	atr := &AccessTokenResponse{}
	err = json.Unmarshal(bodyBytes, atr)
	if err != nil {
		return nil, err
	}
	return atr, nil
}

func (l *Login) RefreshAccessToken() {

}

func (l *Login) UserInfo(token *socialize.AccessTokenDetail) (socialize.UnifiedUser, error) {
	return l.GetUnionID(token.AccessToken, token.OpenId)
}

func (l *Login) UserInfoByCode(code string) (socialize.UnifiedUser, error) {
	token, err := l.AccessToken(code)
	if err != nil {
		return nil, err
	}
	return l.UserInfo(token.GetAccessToken())
}

func (l *Login) GetUnionID(token, openid string) (socialize.UnifiedUser, error) {
	query := url.Values{}
	query.Add("access_token", token)
	query.Add("openid", openid)
	query.Add("lang", "zh_CN")
	dourl := UnionIDPath + "?" + query.Encode()
	resp, err := l.client.Get(dourl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if isError(bodyBytes) {
		return nil, errors.New(string(bodyBytes))
	}

	r := &UnionIDResponse{}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
