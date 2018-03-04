package core

type ServerCallback func(message *Message) []byte

type WriteAble interface {
	ToBytes() []byte
}
