package wego

/*QrCodeScene QrCodeScene*/
type QrCodeScene struct {
	SceneID  int    `json:"scene_id,omitempty"`
	SceneStr string `json:"scene_str,omitempty"`
}

/*QrCodeCard QrCodeCard*/
type QrCodeCard struct {
	CardID       string `json:"card_id,omitempty"`        // "card_id": "pFS7Fjg8kV1IdDz01r4SQwMkuCKc",
	Code         string `json:"code"`                     // "code": "198374613512",
	OpenID       string `json:"openid,omitempty"`         // "openid": "oFS7Fjl0WsZ9AMZqrI80nbIq8xrA",
	IsUniqueCode bool   `json:"is_unique_code,omitempty"` // "is_unique_code": false,
	OuterStr     string `json:"outer_str,omitempty"`      // "outer_str":"12b"
}

/*QrCodeCardList QrCodeCardList*/
type QrCodeCardList struct {
	CardID   string `json:"card_id,omitempty"`   // "card_id": "p1Pj9jgj3BcomSgtuW8B1wl-wo88",
	Code     string `json:"code"`                // "code": "198374613512",
	OuterStr string `json:"outer_str,omitempty"` // "outer_str":"12b"
}

/*QrCodeMultipleCard QrCodeMultipleCard*/
type QrCodeMultipleCard struct {
	CardList []QrCodeCardList `json:"card_list,omitempty"`
}

/*QrCodeActionInfo QrCodeActionInfo*/
type QrCodeActionInfo struct {
	Scene        *QrCodeScene        `json:"scene,omitempty"`
	Card         *QrCodeCard         `json:"card,omitempty"`
	MultipleCard *QrCodeMultipleCard `json:"multiple_card,omitempty"`
}

/*QrCodeAction QrCodeAction*/
type QrCodeAction struct {
	ExpireSeconds int              `json:"expire_seconds,omitempty"`
	ActionName    QrCodeActionName `json:"action_name"`
	ActionInfo    QrCodeActionInfo `json:"action_info"`
}

/*QrCodeActionName QrCodeActionName*/
type QrCodeActionName string

// QrMultipleCard ...
const (
	QrMultipleCard  QrCodeActionName = "QR_MULTIPLE_CARD"   //QrMultipleCard QrMultipleCard
	QrCard          QrCodeActionName = "QR_CARD"            //QrCard QrCard
	QrScene         QrCodeActionName = "QR_SCENE"           //QrScene QrScene
	QrLimitStrScene QrCodeActionName = "QR_LIMIT_STR_SCENE" //QrLimitStrScene QrLimitStrScene
)
