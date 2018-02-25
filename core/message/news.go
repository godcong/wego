package message

type News struct {
	Items []*NewItem
}

type NewItem struct {
	Title       string
	Description string
	PicURL      string
	URL         string
}
