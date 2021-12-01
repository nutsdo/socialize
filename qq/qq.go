package qq

import (
	"encoding/json"
	"errors"
	"github.com/nutsdo/socialize"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ApiHost         = "https://graph.qq.com"
	AccessTokenPath = ApiHost + "/oauth2.0/token"
	UserInfoPath    = ApiHost + "/user/get_user_info"
	UnionIDPath     = ApiHost + "/oauth2.0/me"
)

type LoginFactory struct{}

type Login struct {
	credential  *Credential
	client      *http.Client
	redirectUri string
}

func (lf LoginFactory) CreateLoginClient(appid, secret string, options ...interface{}) socialize.Login {
	var opts []Option
	for _, opt := range options {
		if o, ok := opt.(Option); ok {
			opts = append(opts, o)
		}
	}
	return NewLogin(appid, secret, opts...)
}

type Option func(*Login)

func WithRedirectUri(uri string) Option {
	return func(a *Login) {
		a.redirectUri = uri
	}
}

func NewLogin(appid, secret string, options ...Option) socialize.Login {
	credential := &Credential{
		appid:  appid,
		secret: secret,
	}
	l := &Login{
		credential: credential,
		client:     http.DefaultClient,
	}

	for _, o := range options {
		o(l)
	}
	return l
}

func (l *Login) AccessToken(code string) (socialize.UnifiedAccessToken, error) {

	query := url.Values{}
	query.Add("client_id", l.credential.appid)
	query.Add("client_secret", l.credential.secret)
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")
	query.Add("fmt", "json")
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
	panic("no support!")
}

func (l *Login) UserInfo(token *socialize.AccessTokenDetail) (socialize.UnifiedUser, error) {
	idInfo, err := l.GetUnionID(token.AccessToken)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Add("access_token", token.AccessToken)
	query.Add("oauth_consumer_key", l.credential.appid)
	query.Add("openid", idInfo.Openid)
	dourl := UserInfoPath + "?" + query.Encode()
	resp, err := l.client.Get(dourl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if isError(bodyBytes) {
		return nil, errors.New(string(bodyBytes))
	}

	r := &UserInfoResponse{}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (l *Login) UserInfoByCode(code string) (socialize.UnifiedUser, error) {
	token, err := l.AccessToken(code)
	if err != nil {
		return nil, err
	}
	return l.UserInfo(token.GetAccessToken())
}

func (l *Login) GetUnionID(token string) (*UnionIDResponse, error) {
	query := url.Values{}
	query.Add("access_token", token)
	query.Add("unionid", "true")
	query.Add("fmt", "json")
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
