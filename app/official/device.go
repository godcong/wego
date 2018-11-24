package official

import (
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strconv"
)

/*Device Device */
type Device struct {
	//Config
	*Account
}

func newDevice(acc *Account) *Device {
	return &Device{
		//Config:  defaultConfig,
		Account: acc,
	}
}

/*NewDevice NewDevice*/
func NewDevice(config *core.Config) *Device {
	return newDevice(NewOfficialAccount(config))
}

func (d *Device) TransMessage(deviceID, openID, content string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"device_type": d.Get("device_type"),
		"device_id":   deviceID,
		"open_id":     openID,
		"content":     cipher.Base64Encode([]byte(content)),
	}
	return core.PostJSON("device/transmsg", token, maps)

}

func (d *Device) QrCode(devices []string) core.Responder {
	token := d.accessToken.KeyMap()
	num := strconv.Itoa(len(devices))
	maps := util.Map{
		"device_num":     num,
		"device_id_list": devices,
	}
	return core.PostJSON("'device/create_qrcode", token, maps)
}
