package mini_program

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Template struct {
	config.Config
	*MiniProgram
	//client *core.Client
}

func newTemplate(program *MiniProgram) *Template {
	template := Template{
		Config:      defaultConfig,
		MiniProgram: program,
		//client:      program.GetClient(),
	}
	//template.client.SetDomain(core.NewDomain(""))
	return &template
}

func NewTemplate() *Template {
	return newTemplate(program)
}

func (t *Template) List(offset, count int) util.Map {
	return t.GetClient().HttpPostJson(
		t.client.Link(core.TEMPLATE_LIBRARY_LIST_URL_SUFFIX), util.Map{"offset": offset, "count": count}, nil).ToMap()
}

func (t *Template) Get(id string) util.Map {
	return t.GetClient().HttpPostJson(
		t.client.Link(core.TEMPLATE_LIBRARY_GET_URL_SUFFIX), util.Map{"id": id}, nil).ToMap()
}

func (t *Template) Delete(templateId string) util.Map {
	return t.GetClient().HttpPostJson(
		t.client.Link(core.TEMPLATE_DEL_URL_SUFFIX), util.Map{"template_id": templateId}, nil).ToMap()
}

func (t *Template) GetTemplates(offset, count int) util.Map {
	return t.GetClient().HttpPostJson(
		t.client.Link(core.TEMPLATE_LIST_URL_SUFFIX), util.Map{"offset": offset, "count": count}, nil).ToMap()
}

func (t *Template) Add(id string, keyword util.Map) util.Map {
	return t.GetClient().HttpPostJson(
		t.client.Link(core.TEMPLATE_ADD_URL_SUFFIX), util.Map{"id": id, "keyword_id_list": keyword}, nil).ToMap()
}

func (t *Template) Send(template *message.Template) *net.Response {
	token := t.token.GetToken()
	resp := t.client.HttpPostJson(
		t.client.Link(TEMPLATE_SEND_URL_SUFFIX),
		token.KeyMap(),
		template,
	)
	return resp
}
