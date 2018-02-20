package wego

type AccessTokenInterface interface {
	GetToken() string
	Refresh() AccessTokenInterface
	ApplyToRequest(RequestInterface, Map) RequestInterface
}

type AccessToken struct {
	Config
	app   Application
	token string
}

var acc AccessTokenInterface

func NewAccessToken(application Application, config Config) AccessTokenInterface {
	return &AccessToken{
		Config: config,
		app:    application,
	}
}

func (a *AccessToken) GetToken() string {
	panic("implement me")
}

func (a *AccessToken) getToken(b bool) string {
	panic("implement me")
	//$cacheKey = $this->getCacheKey();
	//$cache = $this->getCache();
	//
	//if (!$refresh && $cache->has($cacheKey)) {
	//return $cache->get($cacheKey);
	//}
	//
	//$token = $this->requestToken($this->getCredentials(), true);
	//
	//$this->setToken($token[$this->tokenKey], $token['expires_in'] ?? 7200);
	//
	//return token;
}

func (a *AccessToken) Refresh() AccessTokenInterface {
	panic("implement me")
}

func (a *AccessToken) ApplyToRequest(RequestInterface, Map) RequestInterface {
	panic("implement me")
}

func (a *AccessToken) GetRefreshedToken(RequestInterface, Map) RequestInterface {
	panic("implement me")
}

func (a *AccessToken) getCredentials() Map {
	return Map{
		"grant_type": "client_credential",
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
	}

}
