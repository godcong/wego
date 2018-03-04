package message

import "encoding/xml"

type Text struct {
	Message
	Content CDATA
}

func (t *Text) ToBytes() []byte {
	bytes, e := xml.Marshal(*t)
	if e != nil {
		return nil
	}
	return bytes
}
