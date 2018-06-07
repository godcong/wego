package message

/*Music Music */
type Music struct {
	Message
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	MusicURL     string `xml:"music_url"`
	HQMusicURL   string `xml:"hq_music_url"`
	ThumbMediaID string `xml:"thumb_media_id"`
}
