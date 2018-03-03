package core

import "github.com/godcong/wego/core/message"

type Message struct {
	message.Common
	//message.Text
	//message.Image
	//message.Music
	//message.Video
	//attributes   Map
	//properties   []string
	//aliases      Map
}

//type Article struct {
//}
//
//func NewMessage() *Message {
//	return &Message{}
//}

//
//func (m *Message) SetAttribute(key string, val interface{}) *Message {
//	m.properties = append(m.properties, key)
//	m.attributes.Set(key, val)
//	return m
//}
//
//func (m *Message) SetAttributes(m0 Map) *Message {
//	for k, v := range m0 {
//		m.SetAttribute(k, v)
//	}
//	return m
//}
//
//func (m *Message) GetAttribute(key string) interface{} {
//	return m.attributes.Get(key)
//}
//
//func (m *Message) GetAttributes(keys []string) []interface{} {
//	var m0 []interface{}
//	for _, v := range keys {
//		m0 = append(m0, m.attributes.Get(v))
//	}
//	return m0
//}

func (m *Message) SetType(msgType message.MsgType) *Message {
	m.MsgType = msgType
	return m
}

func (m *Message) GetType() message.MsgType {
	return m.MsgType
}
