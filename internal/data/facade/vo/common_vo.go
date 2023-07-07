package vo

import (
	"fmt"
)

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Err) Successful() bool {
	if e.Code == 0 {
		return true
	}
	return false
}

func (e Err) Failure() bool {
	return !e.Successful()
}

func (e Err) Error() error {
	if e.Successful() {
		return nil
	}
	return fmt.Errorf("request error code:%v msg:%v", e.Code, e.Message)
}

type VoidErr struct {
	Err
}

type CommonReqVo struct {
	UserId        int64  `json:"userId"`
	OrgId         int64  `json:"orgId"`
	SourceChannel string `json:"sourceChannel"`
}

type Void struct {
	// 主键
	ID int64 `json:"id"`
}

type CommonRespVo struct {
	Err
	Void *Void `json:"data"`
}

type BasicReqVo struct {
	Page uint
	Size uint
}

type BoolRespVo struct {
	Err
	IsTrue bool `json:"data"`
}

type BoolRespVoData struct {
	IsTrue bool `json:"isTrue"`
}

type BasicInfoReqVo struct {
	UserId int64 `json:"userId"`
	OrgId  int64 `json:"orgId"`
}
