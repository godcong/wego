package official

type OfficialAccount interface {
	AccessToken() AccessTokenInterface
}

type officialAccount struct {
	Config
	app Application
}

type OfficialAccountAccessToken struct {
}

func NewOfficialAccount(application Application) OfficialAccount {
	return &officialAccount{
		Config: application.GetConfig("official.default"),
		app:    application,
	}
}

func (a *officialAccount) AccessToken() AccessTokenInterface {
	return NewAccessToken(a.app, a.Config)
}
