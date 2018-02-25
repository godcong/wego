package official

import "github.com/godcong/wego/core"

type OAuth interface {
	//PrepareCallbackUrl(application Application)
}

type Oauth struct {
	core.Config
	client core.Client
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

//
//func NewOAuth(application Application, config Config) OAuth {
//	return &Oauth{
//		Config: config,
//		app:    application,
//		//client: application.Client(),
//	}
//}
