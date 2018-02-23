package official

type OAuth interface {
	PrepareCallbackUrl(application Application)
}

type oauth struct {
	Config
	app Application
}

func (o *oauth) PrepareCallbackUrl(application Application) {
	//$callback = $app['config']->get('oauth.callback');
	//if (0 === stripos($callback, 'http')) {
	//return $callback;
	//}
	//$baseUrl = $app['request']->getSchemeAndHttpHost();
	//
	//return $baseUrl.'/'.ltrim($callback, '/');
}

func NewOAuth(application Application, config Config) OAuth {
	return &oauth{
		Config: config,
		app:    application,
		//client: application.Client(),
	}
}
