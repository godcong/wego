package official_account_test

import (
	"testing"

	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/official_account"
)

func TestNewTemplate(t *testing.T) {
	t0 := official_account.NewTemplate()

	//testTemplate_SetIndustry(t, t0)
	//testTemplate_GetIndustry(t, t0)
	//testTemplate_Add(t, t0)
	testTemplate_Send(t, t0)
}

func testTemplate_SetIndustry(t *testing.T, template *official_account.Template) {
	rlt := template.SetIndustry("1", "2")
	t.Log(rlt.ToString())
}

func testTemplate_GetIndustry(t *testing.T, template *official_account.Template) {
	rlt := template.GetIndustry()
	t.Log(rlt.ToString())
}

func testTemplate_Add(t *testing.T, template *official_account.Template) {
	rlt := template.Add("TM00015")
	t.Log(rlt.ToString())
}

func testTemplate_DelAllPrivate(t *testing.T, template *official_account.Template) {
	rlt := template.GetAllPrivate()
	t.Log(rlt.ToString())
}

func testTemplate_GetAllPrivate(t *testing.T, template *official_account.Template) {
	rlt := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
	t.Log(rlt.ToString())
}

func testTemplate_Send(t *testing.T, template *official_account.Template) {
	//rlt0 := template.GetAllPrivate()
	//rlt1 := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
	//t.Log(rlt0.ToString())
	//t.Log(rlt1.ToString())
	rlt := template.Send(&message.Template{
		ToUser:     "oLyBi0hSYhggnD-kOIms0IzZFqrc",
		TemplateId: "tAsZKUQO0zNkrfvsTi2JexHJ9ZPudXuZSdcurGzE7Yo",
		Url:        "http://baidu.com",
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
	t.Log(rlt.ToString())
}
