package wego

type AccessTokenInterface interface {
	GetToken() string
	Refresh() AccessTokenInterface
	ApplyToRequest(RequestInterface, Map) RequestInterface
}

type accessToken struct {
	Config
	app   Application
	token string
}

var acc AccessTokenInterface

func NewAccessToken(application Application, config Config) AccessTokenInterface {
	return &accessToken{
		Config: config,
		app:    application,
	}
}

func (a *accessToken) GetToken() string {
	panic("implement me")
}

func (a *accessToken) getToken(b bool) string {
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

func (a *accessToken) Refresh() AccessTokenInterface {
	panic("implement me")
}

func (a *accessToken) ApplyToRequest(RequestInterface, Map) RequestInterface {
	panic("implement me")
}

func (a *accessToken) GetRefreshedToken(RequestInterface, Map) RequestInterface {
	panic("implement me")
}

func (a *accessToken) getCredentials() Map {
	return Map{
		"grant_type": "client_credential",
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
	}

}
