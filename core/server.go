package core

import (
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
)

type MessageCallback func(message *Message) message.Messager
type PaymentCallback func(p util.Map) bool

type WriteAble interface {
	ToBytes() []byte
}
