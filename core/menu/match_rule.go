package menu

//type MatchRule string
//
//const (
//	MatchRuleTagId              MatchRule = "tag_id"
//	MatchRuleSex                MatchRule = "sex"
//	MatchRuleCountry            MatchRule = "country"
//	MatchRuleProvince           MatchRule = "province"
//	MatchRuleCity               MatchRule = "city"
//	MatchRuleClientPlatformType MatchRule = "client_platform_type"
//	MatchRuleLanguage           MatchRule = "language"
//)
//
//func (m MatchRule) String() string {
//	return string(m)
//}

type MatchRule struct {
	TagId              string `json:"tag_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}
