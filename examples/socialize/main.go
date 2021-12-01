package main

import (
	"fmt"
)

func main() {
	ExampleSet()
}

//func TestFactory()  {
//	var factory socialize.LoginFactory
//	factory = wechat.LoginFactory{}
//	weappL := factory.CreateLoginClient("appid", "secret")
//
//	user, err := weappL.UserInfoByCode("code")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(user)
//
//	factory = weibo.LoginFactory{}
//	weibologin := factory.CreateLoginClient("appid", "secret",
//		weibo.WithRedirectUri("https://weibo.com"))
//	wbuser, _ := weibologin.UserInfoByCode("code")
//	fmt.Println(wbuser)
//
//	factory = qq.LoginFactory{}
//	qqlogin := factory.CreateLoginClient("appid", "secret",
//		weibo.WithRedirectUri("https://qq.com"))
//	quser, _ := qqlogin.UserInfoByCode("code")
//	fmt.Println(quser)
//}

func ExampleSet() {
	//clientId = "1291311483"
	//clientSecret = "abcabf3f3ee86568f6158012d6679704"
	cfg := &Config{
		Wechat: ThirdApp{
			"weapp": LoginConfig{
				Appid:       "appid",
				Secret:      "secret",
				RedirectUri: "redirect_uri",
			},
		},
		Weibo: ThirdApp{
			"mobile": LoginConfig{
				Appid:       "",
				Secret:      "",
				RedirectUri: "",
			},
			"web": LoginConfig{
				Appid:       "",
				Secret:      "",
				RedirectUri: "",
			},
		},
		QQ: ThirdApp{
			"mobile": LoginConfig{
				Appid:       "",
				Secret:      "",
				RedirectUri: "",
			},
			"web": LoginConfig{
				Appid:       "",
				Secret:      "",
				RedirectUri: "",
			},
		},
	}
	var socialset *Set
	socialset = NewSet(cfg)
	l, err := socialset.GetClient("weibo", "web")
	if err != nil {
		fmt.Println("social client is error :", err)
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", l)
	user, err := l.UserInfoByCode("9ebc59108022bea2348b73f9ccd9d7f9")
	if err != nil {
		fmt.Println("get user info error :", err)
		return
	}
	fmt.Printf("%#v \n", user.GetUser())
	fmt.Println("user info is:")
	fmt.Println(user.GetUser())
}