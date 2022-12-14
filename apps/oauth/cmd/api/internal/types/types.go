// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Code    uint64 `json:"code"`
	Message string `json:"msg"`
}

type Token struct {
	Access string `json:"access"`
}

type AccessReq struct {
	Email  string `json:"email"`
	Secret string `json:"secret"`
}

type AccessResp struct {
	Response
	Data Token `json:"data"`
}
