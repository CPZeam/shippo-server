package model

type WxOffiaccount struct {
	Appid      int    `json:"appid"`
	Avatar_url string `json:"avatar_url"`
	Nickname   string `json:"nickname"`
	Username   string `json:"username"`
	Created_at string `json:"created_at"`
}
