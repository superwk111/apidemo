// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResp struct {
}
