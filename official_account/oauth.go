package official_account

import "github.com/godcong/wego/core"

type OAuth interface {
	//PrepareCallbackUrl(application Application)
}

type Oauth struct {
	core.Config
	client   core.Client
	scopes   string
	callback string
}

func (o *Oauth) PrepareCallbackUrl() {
	//$callback = $app['config']->get('oauth.callback');
	//if (0 === stripos($callback, 'http')) {
	//return $callback;
	//}
	//$baseUrl = $app['request']->getSchemeAndHttpHost();
	//
	//return $baseUrl.'/'.ltrim($callback, '/');
}

func (o *Oauth) User() {

}

//
//func NewOAuth(application Application, config Config) OAuth {
//	return &Oauth{
//		Config: config,
//		app:    application,
//		//client: application.Client(),
//	}
//}

//qq回调配置
//https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=310198347&redirect_uri=http%3A%2F%2Fwww.right.com.cn%2Fforum%2Fconnect.php%3Fmod%3Dlogin%26op%3Dcallback%26referer%3Dhttp%253A%252F%252Fwww.right.com.cn%252Fforum%252Fthread-147109-1-1.html&state=72a5eb8ae2eba26edc851175955d5094&scope=get_user_info%2Cadd_share%2Cadd_t%2Cadd_pic_t%2Cget_repost_list
