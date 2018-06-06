package core

/*MediaType MediaType */
type MediaType string

/*media types */
const (
	MediaTypeImage MediaType = "image"
	MediaTypeVoice MediaType = "voice"
	MediaTypeVideo MediaType = "video"
	MediaTypeThumb MediaType = "thumb"
)

/*String transfer MediaType to string */
func (m MediaType) String() string {
	return string(m)
}
