package media

/*Article Article */
type Article struct {
	Title              string `json:"title"`                           // 标题
	ThumbMediaID       string `json:"thumb_media_id"`                  // 图文消息的封面图片素材id（必须是永久mediaID）
	Author             string `json:"author,omitempty"`                // 作者
	Digest             string `json:"digest,omitempty"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	ShowCoverPic       string `json:"show_cover_pic"`                  // 	是否显示封面，0为false，即不显示，1为true，即显示
	Content            string `json:"content"`                         // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
	ContentSourceURL   string `json:"content_source_url"`              // 图文消息的原文地址，即点击“阅读原文”后的URL
	NeedOpenComment    uint32 `json:"need_open_comment,omitempty"`     // (新增字段）	否	Uint32	是否打开评论，0不打开，1打开
	OnlyFansCanComment uint32 `json:"only_fans_can_comment,omitempty"` // （新增字段）	否	Uint32	是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}
