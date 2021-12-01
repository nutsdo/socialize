package main

import (
	"errors"
	"github.com/nutsdo/socialize"
	"github.com/nutsdo/socialize/qq"
	"github.com/nutsdo/socialize/wechat"
	"github.com/nutsdo/socialize/weibo"
)

//var set Set

type Third map[string]socialize.Login

type Set struct {
	Wechat Third
	Weibo  Third
	QQ     Third
}

func NewSet(cfg *Config) *Set {
	set := &Set{}
	if cfg.Wechat != nil {
		factory := wechat.LoginFactory{}
		if set.Wechat == nil {
			set.Wechat = make(Third)
		}
		for k, c := range cfg.Wechat {
			set.Wechat[k] = factory.CreateLoginClient(c.Appid, c.Secret)
		}
	}
	if cfg.Weibo != nil {
		factory := weibo.LoginFactory{}
		if set.Weibo == nil {
			set.Weibo = make(Third)
		}
		for k, c := range cfg.Weibo {
			set.Weibo[k] = factory.CreateLoginClient(
				c.Appid,
				c.Secret,
				weibo.WithRedirectUri(c.RedirectUri))
		}
	}
	if cfg.QQ != nil {
		factory := qq.LoginFactory{}
		if set.QQ == nil {
			set.QQ = make(Third)
		}
		for k, c := range cfg.QQ {
			set.QQ[k] = factory.CreateLoginClient(
				c.Appid,
				c.Secret,
				qq.WithRedirectUri(c.RedirectUri))
		}
	}
	return set
}

func (s Set) GetClient(platform, key string) (socialize.Login, error) {

	if len(platform) == 0 {
		return nil, errors.New("platform is required")
	}

	if len(key) == 0 {
		return nil, errors.New("key is required")
	}
	if platform == "wechat" {
		if value, ok := s.Wechat[key]; ok {
			return value, nil
		}
	}
	if platform == "weibo" {
		if value, ok := s.Weibo[key]; ok {
			return value, nil
		}
	}
	if platform == "qq" {
		if value, ok := s.QQ[key]; ok {
			return value, nil
		}
	}
	return nil, errors.New("not found client")
}
