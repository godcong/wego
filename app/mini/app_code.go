package mini

import (
	"github.com/godcong/wego/core"
	"log"

	"github.com/godcong/wego/util"
)

/*AppCode AppCode*/
type AppCode struct {
	*Program
}

func newAppcode(program *Program) interface{} {
	return &AppCode{
		Program: program,
	}
}

// NewAppCode ...
func NewAppCode(config *core.Config) *AppCode {
	return newAppcode(NewMiniProgram(config)).(*AppCode)
}

/*Get  获取小程序二维码
我们推荐生成并使用小程序码，它具有更好的辨识度。目前有两个接口可以生成小程序码，开发者可以根据自己的需要选择合适的接口。
接口A: 适用于需要的码数量较少的业务场景 接口地址:
https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN
获取 access_token 详见文档
POST 参数说明
参数	类型	默认值	说明
path	String		不能为空，最大长度 128 字节
width	Int	430	二维码的宽度
auto_color	Bool	false	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
line_color	Object	{"r":"0","g":"0","b":"0"}	auth_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
is_hyaline	Bool	false	是否需要透明底色， is_hyaline 为true时，生成透明底色的小程序码
注意:通过该接口生成的小程序码，永久有效，数量限制见文末说明，请谨慎使用。用户扫描该码进入小程序后，将直接进入 path 对应的页面。
*/
func (a *AppCode) Get(path string, optionals util.Map) util.Map {
	params := util.Map{"path": path}
	params.Join(optionals)

	j := a.getStream(Link(wxaGetWXACode), params)

	return util.JSONToMap(j)
}

/*GetQrCode 获取小程序二维码
接口C:适用于需要的码数量较少的业务场景
接口地址:
https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN
获取 access_token 详见文档
POST 参数说明
参数	类型	默认值	说明
path	String		不能为空，最大长度 128 字节
width	Int	430	二维码的宽度
注意:通过该接口生成的小程序二维码，永久有效，数量限制见文末说明，请谨慎使用。用户扫描该码进入小程序后，将直接进入 path 对应的页面。
*/
func (a *AppCode) GetQrCode(path string, width int) util.Map {
	params := util.Map{"path": path, "width": width}

	j := a.getStream(Link(wxaappCreatewxaqrcode), params)
	return util.JSONToMap(j)
}

/*GetUnlimit 获取小程序码
我们推荐生成并使用小程序码，它具有更好的辨识度。目前有两个接口可以生成小程序码，开发者可以根据自己的需要选择合适的接口。
接口B:适用于需要的码数量极多的业务场景
接口地址:
https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
获取 access_token 详见文档
POST 参数说明
参数	类型	默认值	说明
scene	String		最大32个可见字符，只支持数字，大小写英文以及部分特殊字符:!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
page	String		必须是已经发布的小程序存在的页面（否则报错），例如 "pages/index/index" ,根路径前不要填加'/',不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
width	Int	430	二维码的宽度
auto_color	Bool	false	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
line_color	Object	{"r":"0","g":"0","b":"0"}	auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
is_hyaline	Bool	false	是否需要透明底色， is_hyaline 为true时，生成透明底色的小程序码
注意:通过该接口生成的小程序码，永久有效，数量暂无限制。用户扫描该码进入小程序后，开发者需在对应页面获取的码中 scene 字段的值，再做处理逻辑。使用如下代码可以获取到二维码中的 scene 字段的值。调试阶段可以使用开发工具的条件编译自定义参数 scene=xxxx 进行模拟，开发工具模拟时的 scene 的参数值需要进行 urlencode
*/
func (a *AppCode) GetUnlimit(scene string, optionals util.Map) util.Map {
	params := util.Map{"scene": scene}
	params.Join(optionals)

	j := a.getStream(Link(wxaGetWXACodeUnlimit), params)
	return util.JSONToMap(j)
}

func (a *AppCode) getStream(url string, m util.Map) []byte {
	log.Println(url, m)
	token0 := a.AccessToken().GetToken()
	token := token0.KeyMap()
	//strings.Join([]string{"access_token", token0.GetKey()}, "=")

	resp := core.RequestRaw(
		url,
		"post",
		util.Map{core.DataTypeQuery: token.URLEncode(),
			core.DataTypeJSON: m})
	panic(resp)
	//TODO
	return nil
}
