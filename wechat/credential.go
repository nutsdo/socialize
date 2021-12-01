package wechat

type Credential struct {
	appid  string
	secret string
}

func NewCredential(appid, secret string) *Credential {
	return &Credential{appid: appid, secret: secret}
}
