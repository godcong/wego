package message

type News struct {
	Message
	Items []*NewItem
}

type NewItem struct {
	Title       string
	Description string
	PicURL      string
	URL         string
}
