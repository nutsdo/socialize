package main

type Config struct {
	Wechat ThirdApp
	Weibo  ThirdApp
	QQ     ThirdApp
}

type ThirdApp map[string]LoginConfig

type LoginConfig struct {
	Appid       string
	Secret      string
	RedirectUri string
}
