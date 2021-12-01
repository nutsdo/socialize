package socialize

type UserDetail struct {
	Openid     string   `json:"openid"`
	Unionid    string   `json:"unionid"`
	WeiboUid   string   `json:"weibo_uid"`
	Nickname   string   `json:"nickname"`
	Gender     int      `json:"gender"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Lang       string   `json:"lang"`
}
