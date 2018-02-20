package wego

type OfficialAccount interface {
	AccessToken() AccessTokenInterface
}

type officialAccount struct {
	Config
	app Application
}

func NewOfficialAccount(application Application) OfficialAccount {
	return &officialAccount{
		Config: application.GetConfig("official_account.default"),
		app:    application,
	}
}

func (a *officialAccount) AccessToken() AccessTokenInterface {
	return NewAccessToken(a.app, a.Config)
}
