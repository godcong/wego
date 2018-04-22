package core

import "github.com/godcong/wego/core/message"

type MessageCallback func(message *Message) message.Messager
type PaymentCallback func(p Map) bool

type WriteAble interface {
	ToBytes() []byte
}
