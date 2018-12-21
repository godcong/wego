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

// TransMessage ...
func (d *Device) TransMessage(deviceID, openID, content string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"device_type": d.Get("device_type"),
		"device_id":   deviceID,
		"open_id":     openID,
		"content":     cipher.Base64Encode([]byte(content)),
	}
	return core.PostJSON(deviceTransmsg, token, maps)

}

// CreateQrCode ...
func (d *Device) CreateQrCode(devices []string) core.Responder {
	token := d.accessToken.KeyMap()
	num := strconv.Itoa(len(devices))
	maps := util.Map{
		"device_num":     num,
		"device_id_list": devices,
	}
	return core.PostJSON(deviceCreateQrcode, token, maps)
}

// Authorize ...
func (d *Device) Authorize(ables []util.MapAble, productID string, optype int) core.Responder {
	token := d.accessToken.KeyMap()
	num := strconv.Itoa(len(ables))
	var tmps []util.Map
	for _, v := range ables {
		tmps = append(tmps, v.ToMap())
	}
	maps := util.Map{
		"device_num":  num,
		"device_list": tmps,
		"product_id":  productID,
		"op_type":     optype,
	}
	return core.PostJSON(deviceAuthorizeDevice, token, maps)
}

// GetQrCode ...
func (d *Device) GetQrCode(productID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"product_id": productID,
	}
	return core.PostJSON(deviceGetqrcode, token, maps)
}

// Bind ...
func (d *Device) Bind(openID, deviceID, ticket string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"openid":    openID,
		"device_id": deviceID,
		"ticket":    ticket,
	}
	return core.PostJSON(deviceBind, token, maps)
}

// Unbind ...
func (d *Device) Unbind(openID, deviceID, ticket string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"openid":    openID,
		"device_id": deviceID,
		"ticket":    ticket,
	}
	return core.PostJSON(deviceUnbind, token, maps)
}

// CompelBind ...
func (d *Device) CompelBind(openID, deviceID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"openid":    openID,
		"device_id": deviceID,
	}
	return core.PostJSON(deviceCompelBind, token, maps)
}

// CompelUnbind ...
func (d *Device) CompelUnbind(openID, deviceID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"openid":    openID,
		"device_id": deviceID,
	}
	return core.PostJSON(deviceCompelUnbind, token, maps)
}

// GetStatus ...
func (d *Device) GetStatus(deviceID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"device_id": deviceID,
	}
	return core.PostJSON(deviceGetStat, token, maps)
}

// VerifyQrCode ...
func (d *Device) VerifyQrCode(ticket string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"ticket": ticket,
	}
	return core.PostJSON(deviceVerifyQrcode, token, maps)
}

// GetOpenid ...
func (d *Device) GetOpenid(deviceID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"device_id": deviceID,
	}
	return core.PostJSON(deviceGetOpenid, token, maps)
}

// GetBindDevice ...
func (d *Device) GetBindDevice(openID string) core.Responder {
	token := d.accessToken.KeyMap()
	maps := util.Map{
		"open_id": openID,
	}
	return core.PostJSON(deviceGetBindDevice, token, maps)
}
