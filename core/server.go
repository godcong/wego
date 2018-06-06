package core

import (
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
)

/*MessageCallback 消息回调函数定义 */
type MessageCallback func(message *Message) message.Messager

/*PaymentCallback 支付回调函数定义 */
type PaymentCallback func(p util.Map) bool

/*WriteAble WriteAble*/
type WriteAble interface {
	ToBytes() []byte
}
