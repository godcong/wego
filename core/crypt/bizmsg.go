package crypt

import "github.com/godcong/wego/core"

type BizMsg struct {
	token          string
	encodingAESKey string
	appId          string
}

func NewBizMsg(token, key, id string) *BizMsg {
	return &BizMsg{
		token:          token,
		encodingAESKey: key,
		appId:          id,
	}
}

func (m *BizMsg) Encrypt(p core.Map) core.Map {

}
