package crypt

import "github.com/godcong/wego/core/tool"

type PrpCrypt struct {
	key string
}

func NewPrp(key string) *PrpCrypt {
	return &PrpCrypt{
		key: key,
	}
}

func (*PrpCrypt) Random() string {
	return tool.GenerateRandomString(16, tool.T_RAND_ALL)
}

//func ()  {
//
//}
