package core

type MessageCallback func(message *Message) []byte
type PaymentCallback func(p Map) bool

type WriteAble interface {
	ToBytes() []byte
}
