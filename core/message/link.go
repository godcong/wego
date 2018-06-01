package message

type Link struct {
	*Message
	Title       string //消息标题
	Url         string //消息链接
	Description string //消息描述
}

func NewLink(msg *Message, title, url, desc string) *Link {
	return &Link{
		Message:     msg,
		Title:       title,
		Url:         url,
		Description: desc,
	}
}
