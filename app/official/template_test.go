package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core/message"
)

func TestNewTemplate(t *testing.T) {
	t0 := official.NewTemplate(config)

	//testTemplate_SetIndustry(t, t0)
	//testTemplate_GetIndustry(t, t0)
	testTemplate_Add(t, t0)
	//testTemplate_Send(t, t0)
}

func testTemplate_SetIndustry(t *testing.T, template *official.Template) {
	resp := template.SetIndustry("1", "2")
	t.Log(string(resp.Bytes()))
}

func testTemplate_GetIndustry(t *testing.T, template *official.Template) {
	resp := template.GetIndustry()
	t.Log(string(resp.Bytes()))
}

func testTemplate_Add(t *testing.T, template *official.Template) {
	resp := template.Add("TM00015")
	t.Log(string(resp.Bytes()))
}

func testTemplate_DelAllPrivate(t *testing.T, template *official.Template) {
	resp := template.GetAllPrivate()
	t.Log(string(resp.Bytes()))
}

func testTemplate_GetAllPrivate(t *testing.T, template *official.Template) {
	resp := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
	t.Log(string(resp.Bytes()))
}

func testTemplate_Send(t *testing.T, template *official.Template) {
	//resp0 := template.GetAllPrivate()
	//resp1 := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
	//t.Log(resp0.ToString())
	//t.Log(resp1.ToString())
	resp := template.Send(&message.Template{
		ToUser:     "ogJPnwU54xYMetJLjCbh6ycRvJW4",
		TemplateID: "tAsZKUQO0zNkrfvsTi2JexHJ9ZPudXuZSdcurGzE7Yo",
		URL:        "http://baidu.com",
		Data: message.TemplateData{
			"first": &message.ValueColor{
				Value: "恭喜你成功购买奇葩商品一枚！",
				Color: "#173177",
			},
			//KeyNote1: message.ValueColor{
			//	Value: "333",
			//	Color: "#173177",
			//},
			//KeyNote2: message.ValueColor{
			//	Value: "555",
			//	Color: "#173177",
			//},
			//KeyNote3: message.ValueColor{
			//	Value: "2018年",
			//	Color: "#173177",
			//},
			"Remark": &message.ValueColor{
				Value: "欢迎再次买买买！",
			},
			"orderProductName": &message.ValueColor{
				Value: "神奇的奇葩",
			},
			"orderMoneySum": &message.ValueColor{
				Value: "9999.00元",
				Color: "#122177",
			},
		},
	})
	t.Log(string(resp.Bytes()))
}
