package mini_test

import (
	"testing"

	"github.com/godcong/wego/app/mini"
	"github.com/godcong/wego/core/message"
)

func TestNewTemplate(t *testing.T) {
	//t0 := mini.NewTemplate()

	//testTemplate_SetIndustry(t, t0)
	//testTemplate_GetIndustry(t, t0)
	//testTemplate_Add(t, t0)
	//testTemplate_Send(t, t0)
}

//func testTemplate_SetIndustry(t *testing.T, template *mini.Template) {
//	rlt := template.SetIndustry("1", "2")
//	t.Log(rlt.ToString())
//}
//
//func testTemplate_GetIndustry(t *testing.T, template *mini.Template) {
//	rlt := template.GetIndustry()
//	t.Log(rlt.ToString())
//}
//
//func testTemplate_Add(t *testing.T, template *mini.Template) {
//	rlt := template.Add("TM00015")
//	t.Log(rlt.ToString())
//}
//
//func testTemplate_DelAllPrivate(t *testing.T, template *mini.Template) {
//	rlt := template.GetAllPrivate()
//	t.Log(rlt.ToString())
//}
//
//func testTemplate_GetAllPrivate(t *testing.T, template *mini.Template) {
//	rlt := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
//	t.Log(rlt.ToString())
//}

func testTemplate_Send(t *testing.T, template *mini.Template) {
	//rlt0 := template.GetAllPrivate()
	//rlt1 := template.DelAllPrivate("vc2ekfQmEP9qE9eBW9gGWaUrsLvztC9XOeB-cftLroo")
	//t.Log(rlt0.ToString())
	//t.Log(rlt1.ToString())
	rlt := template.Send(&message.Template{
		ToUser:     "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
		TemplateId: "0-A8LciZI4nQpjFnQ_jtykix4rqKlMcqbSILDaJKPhQ",
		Url:        "",
		Data: message.TemplateData{
			"keyword1": &message.ValueColor{
				Value: "恭喜你成功购买奇葩商品一枚！",
				Color: "#173177",
			},
			"keyword2": &message.ValueColor{
				Value: "恭喜你成功购买奇葩商品一枚！",
				Color: "#173177",
			},
			"keyword3": &message.ValueColor{
				Value: "恭喜你成功购买奇葩商品一枚！",
				Color: "#173177",
			},
			"keyword4": &message.ValueColor{
				Value: "恭喜你成功购买奇葩商品一枚！",
				Color: "#173177",
			},
		},
		Page:            "index?value=123",
		FormId:          "1523991474645",
		EmphasisKeyword: "keyword1.DATA",
	})
	t.Log(rlt.ToString())
}
