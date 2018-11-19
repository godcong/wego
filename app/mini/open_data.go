package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*OpenData OpenData */
type OpenData struct {
	*Program
}

func newOpenData(program *Program) interface{} {
	return &OpenData{
		Program: program,
	}
}

// NewOpenData ...
func NewOpenData(config *core.Config) *OpenData {
	return newOpenData(NewMiniProgram(config)).(*OpenData)
}

// RemoveUserStorage ...
func (d *OpenData) RemoveUserStorage(openID string, sessionKey string, keys []string) core.Responder {
	maps := util.Map{
		"key": keys,
	}

	query := util.Map{
		"appid":      d.Get("app_id"),
		"secret":     d.Get("secret"),
		"openid":     openID,
		"sig_method": "hmac_sha256",
		"signature":  util.MakeSignHMACSHA256(string(maps.ToJSON()), sessionKey),
	}
	return core.PostJSON(Link(wxaRemoveUserStorage), query, maps)
}

// SetUserStorage ...
func (d *OpenData) SetUserStorage(openID string, sessionKey string, list util.Map) core.Responder {
	maps := util.Map{
		"kv_list": list,
	}

	query := util.Map{
		"appid":      d.Get("app_id"),
		"secret":     d.Get("secret"),
		"openid":     openID,
		"sig_method": "hmac_sha256",
		"signature":  util.MakeSignHMACSHA256(string(maps.ToJSON()), sessionKey),
	}
	return core.PostJSON(Link(wxaSetUserStorage), query, maps)
}
