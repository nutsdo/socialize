package socialize

type AccessTokenDetail struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"` // 仅微信
	OpenId       string `json:"open_id"`       // 仅微信
	Scope        string `json:"scope"`         // 仅微信
	WeiboUid     string `json:"weibo_uid"`     // 仅微博
}
