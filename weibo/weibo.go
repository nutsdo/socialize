package weibo

import (
	"encoding/json"
	"errors"
	"github.com/nutsdo/socialize"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ApiHost         = "https://api.weibo.com"
	AccessTokenPath = ApiHost + "/oauth2/access_token"
	UserInfoPath    = ApiHost + "/2/users/show.json"
	UserUidPath     = ApiHost + "/2/account/get_uid.json"
)

type LoginFactory struct{}

func (lf LoginFactory) CreateLoginClient(appid, secret string, options ...interface{}) socialize.Login {
	var opts []Option
	for _, opt := range options {
		if o, ok := opt.(Option); ok {
			opts = append(opts, o)
		}
	}
	return NewLogin(appid, secret, opts...)
}

type Login struct {
	credential  *Credential
	client      *http.Client
	redirectUri string
}

type Option func(*Login)

func WithRedirectUri(uri string) Option {
	return func(a *Login) {
		a.redirectUri = uri
	}
}

func NewLogin(appid, secret string, opts ...Option) socialize.Login {
	credential := &Credential{
		clientId:     appid,
		clientSecret: secret,
	}
	l := &Login{
		credential: credential,
		client:     http.DefaultClient,
	}
	for _, o := range opts {
		o(l)
	}
	return l
}

func (c *Login) AccessToken(code string) (socialize.UnifiedAccessToken, error) {

	query := url.Values{}
	query.Add("client_id", c.credential.clientId)
	query.Add("client_secret", c.credential.clientSecret)
	query.Add("code", code)
	query.Add("grant_type", "authorization_code")
	query.Add("redirect_uri", c.redirectUri)

	resp, err := c.client.PostForm(AccessTokenPath, query)
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

func (c *Login) RefreshAccessToken() {

}

func (c *Login) UserInfo(token *socialize.AccessTokenDetail) (socialize.UnifiedUser, error) {
	uid, err := c.GetUID(token.AccessToken)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add("access_token", token.AccessToken)
	query.Add("uid", strconv.FormatInt(uid.Uid, 10))
	dourl := UserInfoPath + "?" + query.Encode()
	resp, err := c.client.Get(dourl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if isError(bodyBytes) {
		return nil, errors.New(string(bodyBytes))
	}

	r := &UserResponse{}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Login) UserInfoByCode(code string) (socialize.UnifiedUser, error) {
	token, err := c.AccessToken(code)
	if err != nil {
		return nil, err
	}
	return c.UserInfo(token.GetAccessToken())
}

func (c *Login) GetUID(token string) (*UIDResponse, error) {

	query := url.Values{}
	query.Add("access_token", token)
	dourl := UserUidPath + "?" + query.Encode()
	resp, err := c.client.Get(dourl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if isError(bodyBytes) {
		return nil, errors.New(string(bodyBytes))
	}

	r := &UIDResponse{}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
