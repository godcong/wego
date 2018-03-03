package wego

type OfficialAccount interface {
}

func GetOfficialAccount() OfficialAccount {
	obj := GetApp().Get("official_account").(OfficialAccount)
	return obj
}
