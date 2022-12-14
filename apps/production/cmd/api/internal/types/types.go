// Code generated by goctl. DO NOT EDIT.
package types

type Car struct {
	ID    uint64 `json:"id"`
	Brand string `json:"brand"`
	Serie string `json:"serie"`
}

type CarEstbReq struct {
	Brand string `json:"brand"`
	Serie string `json:"serie"`
}

type CarEstbResp struct {
	Code    uint64 `json:"code"`
	Message string `json:"msg"`
	Data    Car    `json:"data"`
}

type CarInfoReq struct {
	ID uint64 `json:"id"`
}

type CarInfoResp struct {
	Code    uint64 `json:"code"`
	Message string `json:"msg"`
	Data    Car    `json:"data"`
}

type Brand struct {
	ID       uint64 `json:"id"`
	Category uint8  `json:"category"`
	Name     string `json:"name"`
}

type BrandEstbReq struct {
	Category uint8  `json:"category"`
	Brand    string `json:"brand"`
}

type BrandEstbResp struct {
	Code    uint64 `json:"code"`
	Message string `json:"msg"`
	Data    Brand  `json:"data"`
}

type BrandListReq struct {
	Page uint16 `json:"page"`
	Size uint16 `json:"size" binding:"required"`
}

type BrandListResp struct {
	Code    uint64  `json:"code"`
	Message string  `json:"msg"`
	Data    []Brand `json:"data"`
}
