package wego

type AccessTokenInterface interface {
	GetToken() string
	Refresh() AccessTokenInterface
	ApplyToRequest(RequestInterface, Map) RequestInterface
	//getCredentials() Map
	//getQuery() Map
	//sendRequest() []byte
}

type AccessToken struct {
	Config
	app   Application
	token string
}

func (a *AccessToken) getQuery() Map {
	panic("implement me")
}

func (a *AccessToken) sendRequest() []byte {
	panic("implement me")
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

func (a *AccessToken) GetRefreshToken() string {
	panic("implement me")
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
	panic("implement me")
}
