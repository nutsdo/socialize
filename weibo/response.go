package weibo

import (
	"github.com/nutsdo/socialize"
)


type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RemindIn     string `json:"remind_in"`
	Uid          string `json:"uid"`
}

func (r *AccessTokenResponse) GetAccessToken() *socialize.AccessTokenDetail {
	return &socialize.AccessTokenDetail{
		AccessToken:  r.AccessToken,
		ExpiresIn:    r.ExpiresIn,
		WeiboUid: r.Uid,
	}
}

type UIDResponse struct {
	Uid int64 `json:"uid"`
}

type UserResponse struct {
	Id                int64  `json:"id"`
	Idstr             string `json:"idstr"`
	Class             int    `json:"class"`
	ScreenName        string `json:"screen_name"`
	Name              string `json:"name"`
	Province          string `json:"province"`
	City              string `json:"city"`
	Location          string `json:"location"`
	Description       string `json:"description"`
	Url               string `json:"url"`
	ProfileImageUrl   string `json:"profile_image_url"`
	ProfileUrl        string `json:"profile_url"`
	Domain            string `json:"domain"`
	Weihao            string `json:"weihao"`
	Gender            string `json:"gender"`
	FollowersCount    int    `json:"followers_count"`
	FollowersCountStr string `json:"followers_count_str"`
	FriendsCount      int    `json:"friends_count"`
	PagefriendsCount  int    `json:"pagefriends_count"`
	StatusesCount     int    `json:"statuses_count"`
	VideoStatusCount  int    `json:"video_status_count"`
	VideoPlayCount    int    `json:"video_play_count"`
	FavouritesCount   int    `json:"favourites_count"`
	CreatedAt         string `json:"created_at"`
	Following         bool   `json:"following"`
	AllowAllActMsg    bool   `json:"allow_all_act_msg"`
	GeoEnabled        bool   `json:"geo_enabled"`
	Verified          bool   `json:"verified"`
	VerifiedType      int    `json:"verified_type"`
	Remark            string `json:"remark"`
	Insecurity        struct {
		SexualContent bool `json:"sexual_content"`
	} `json:"insecurity"`
	Ptype              int    `json:"ptype"`
	AllowAllComment    bool   `json:"allow_all_comment"`
	AvatarLarge        string `json:"avatar_large"`
	AvatarHd           string `json:"avatar_hd"`
	VerifiedReason     string `json:"verified_reason"`
	VerifiedTrade      string `json:"verified_trade"`
	VerifiedReasonUrl  string `json:"verified_reason_url"`
	VerifiedSource     string `json:"verified_source"`
	VerifiedSourceUrl  string `json:"verified_source_url"`
	FollowMe           bool   `json:"follow_me"`
	Like               bool   `json:"like"`
	LikeMe             bool   `json:"like_me"`
	OnlineStatus       int    `json:"online_status"`
	BiFollowersCount   int    `json:"bi_followers_count"`
	Lang               string `json:"lang"`
	Star               int    `json:"star"`
	Mbtype             int    `json:"mbtype"`
	Mbrank             int    `json:"mbrank"`
	Svip               int    `json:"svip"`
	BlockWord          int    `json:"block_word"`
	BlockApp           int    `json:"block_app"`
	CreditScore        int    `json:"credit_score"`
	UserAbility        int    `json:"user_ability"`
	Urank              int    `json:"urank"`
	StoryReadState     int    `json:"story_read_state"`
	VclubMember        int    `json:"vclub_member"`
	IsTeenager         int    `json:"is_teenager"`
	IsGuardian         int    `json:"is_guardian"`
	IsTeenagerList     int    `json:"is_teenager_list"`
	PcNew              int    `json:"pc_new"`
	SpecialFollow      bool   `json:"special_follow"`
	PlanetVideo        int    `json:"planet_video"`
	VideoMark          int    `json:"video_mark"`
	LiveStatus         int    `json:"live_status"`
	UserAbilityExtend  int    `json:"user_ability_extend"`
	StatusTotalCounter struct {
		TotalCnt       int `json:"total_cnt"`
		RepostCnt      int `json:"repost_cnt"`
		CommentCnt     int `json:"comment_cnt"`
		LikeCnt        int `json:"like_cnt"`
		CommentLikeCnt int `json:"comment_like_cnt"`
	} `json:"status_total_counter"`
	VideoTotalCounter struct {
		PlayCnt int `json:"play_cnt"`
	} `json:"video_total_counter"`
	BrandAccount int `json:"brand_account"`
}

func (r *UserResponse) GetUser() *socialize.UserDetail {
	gender := 0
	if r.Gender == "m" {
		gender = 1
	}else if r.Gender == "f" {
		gender = 2
	}
	return &socialize.UserDetail{
		WeiboUid:   r.Idstr,
		Nickname:   r.ScreenName,
		Gender:     gender,
		Province:   r.Province,
		City:       r.City,
		Country:    r.Location,
		Headimgurl: r.AvatarHd,
		Privilege:  nil,
		Lang: r.Lang,
	}
}
