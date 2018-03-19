package mini_program

import "github.com/godcong/wego/core"

type Template struct {
	core.Config
	*MiniProgram
}

func (t *Template) List(offset, count int) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_LIBRARY_LIST_URL_SUFFIX), core.Map{"offset": offset, "count": count}, nil).ToMap()
}

func (t *Template) Get(id string) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_LIBRARY_GET_URL_SUFFIX), core.Map{"id": id}, nil).ToMap()
}

func (t *Template) Delete(templateId string) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_DEL_URL_SUFFIX), core.Map{"template_id": templateId}, nil).ToMap()
}

func (t *Template) GetTemplates(offset, count int) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_LIST_URL_SUFFIX), core.Map{"offset": offset, "count": count}, nil).ToMap()
}

func (t *Template) Add(id string, keyword core.Map) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_ADD_URL_SUFFIX), core.Map{"id": id, "keyword_id_list": keyword}, nil).ToMap()
}

func (t *Template) Send(data core.Map) core.Map {
	return t.GetClient().HttpPostJson(t.client.Link(core.TEMPLATE_SEND_URL_SUFFIX), data, nil).ToMap()
}
