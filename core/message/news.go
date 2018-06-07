package message

/*News News */
type News struct {
	Message
	Items []*NewItem `xml:"items"`
}

/*NewItem NewItem */
type NewItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PicURL      string `xml:"pic_url"`
	URL         string `xml:"url"`
}
