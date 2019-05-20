package wego

import "github.com/godcong/wego/util"

/*Button Button */
type Button struct {
	util.Map
}

/*MatchRule MatchRule*/
type MatchRule struct {
	TagID              string `json:"tag_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

/*NewClickButton NewClickButton*/
func NewClickButton(name, key string) *Button {
	return newButton(EventTypeClick, util.Map{"name": name, "key": key})

}

/*NewViewButton NewViewButton*/
func NewViewButton(name, url string) *Button {
	return newButton(EventTypeView, util.Map{"name": name, "url": url})
}

/*NewSubButton NewSubButton*/
func NewSubButton(name string, sub []*Button) *Button {
	return newButton("", util.Map{"name": name, "key": "testkey", "sub_button": sub})
}

func newButton(typ EventType, val util.Map) *Button {
	button := NewBaseButton()
	if typ != "" {
		button.Set("type", typ)
	}
	button.Join(val)
	return button
}

/*SetSub SetSub*/
func (b *Button) SetSub(name string, sub []*Button) *Button {
	b.Map = util.Map{}
	b.Set("name", name)
	b.Set("sub_button", sub)
	return b
}

/*NewBaseButton NewBaseButton */
func NewBaseButton() *Button {
	return &Button{
		Map: make(util.Map),
	}
}

/*SetButtons SetButtons*/
func (b *Button) SetButtons(buttons []*Button) *Button {
	b.Set("button", buttons)
	return b
}

/*GetButtons GetButtons */
func (b *Button) GetButtons() []*Button {
	buttons := b.Get("button")
	if v0, b := buttons.([]*Button); b {
		return v0
	}
	return nil
}

/*GetMatchRule GetMatchRule*/
func (b *Button) GetMatchRule() *MatchRule {
	if mr := b.Get("matchrule"); mr != nil {
		return mr.(*MatchRule)
	}
	return nil
}

/*SetMatchRule SetMatchRule*/
func (b *Button) SetMatchRule(rule *MatchRule) *Button {
	b.Set("matchrule", rule)
	return b
}

func (b *Button) mapGet(name string) interface{} {
	return b.Get(name)
}

func (b *Button) mapSet(name string, v interface{}) util.Map {
	return b.Set(name, v)
}

/*AddButton AddButton*/
func (b *Button) AddButton(buttons *Button) *Button {
	if v := b.GetButtons(); v != nil {
		b.SetButtons(append(v, buttons))
	} else {
		b.SetButtons([]*Button{buttons})
	}
	return b
}
