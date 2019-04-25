package wego

import (
	"context"
	"crypto/tls"
	"github.com/sirupsen/logrus"
)

// PaymentOption ...
type PaymentOption func(obj *Payment)

// PaymentBodyType ...
func PaymentBodyType(s BodyType) PaymentOption {
	return func(obj *Payment) {
		obj.BodyType = s
	}
}

// PaymentSubID ...
func PaymentSubID(mchid, appid string) PaymentOption {
	return func(obj *Payment) {
		obj.subAppID = appid
		obj.subMchID = mchid
	}
}

// PaymentKey ...
func PaymentKey(public, privite string) PaymentOption {
	return func(obj *Payment) {
		obj.publicKey = public
		obj.privateKey = privite
	}
}

// PaymentRemote ...
func PaymentRemote(remote string) PaymentOption {
	return func(obj *Payment) {
		obj.remoteURL = remote
	}
}

// PaymentLocal ...
func PaymentLocal(local string) PaymentOption {
	return func(obj *Payment) {
		obj.localHost = local
	}
}

// PaymentSandbox ...
func PaymentSandbox(sandbox *Sandbox) PaymentOption {
	return func(obj *Payment) {
		obj.sandbox = sandbox
	}
}

// PaymentNotifyURL ...
func PaymentNotifyURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.notifyURL = s
	}
}

// PaymentRefundedURL ...
func PaymentRefundedURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.refundedURL = s
	}
}

// PaymentScannedURL ...
func PaymentScannedURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.scannedURL = s
	}
}

// AccessTokenOption ...
type AccessTokenOption func(obj *AccessToken)

// AccessTokenRemote ...
func AccessTokenRemote(url string) AccessTokenOption {
	return func(obj *AccessToken) {
		obj.remoteURL = url
	}
}

// AccessTokenURL ...
func AccessTokenURL(url string) AccessTokenOption {
	return func(obj *AccessToken) {
		obj.tokenURL = url
	}
}

// AccessTokenKey ...
func AccessTokenKey(key string) AccessTokenOption {
	return func(obj *AccessToken) {
		obj.tokenKey = key
	}
}

// OfficialAccountOption ...
type OfficialAccountOption func(obj *OfficialAccount)

// OfficialAccountOauth ...
func OfficialAccountOauth(oauth *OAuthProperty) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.oauth = *oauth
	}
}

// OfficialAccountAccessTokenProperty ...
func OfficialAccountAccessTokenProperty(property *AccessTokenProperty) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.accessToken = NewAccessToken(property, AccessTokenKey(accessTokenKey), AccessTokenURL(accessTokenURLSuffix))
	}
}

// OfficialAccountAccessToken ...
func OfficialAccountAccessToken(token *AccessToken) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.accessToken = token
	}
}

// OfficialAccountBodyType ...
func OfficialAccountBodyType(bodyType BodyType) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.BodyType = bodyType
	}
}

// OfficialAccountRemote ...
func OfficialAccountRemote(remote string) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.remoteURL = remote
	}
}

// OfficialAccountLocal ...
func OfficialAccountLocal(local string) OfficialAccountOption {
	return func(obj *OfficialAccount) {
		obj.localHost = local
	}
}

// ClientOption ...
type ClientOption func(obj *Client)

// ClientContext ...
func ClientContext(ctx context.Context) ClientOption {
	return func(obj *Client) {
		obj.context = ctx
	}
}

// ClientAccessToken ...
func ClientAccessToken(token *AccessToken) ClientOption {
	return func(obj *Client) {
		obj.accessToken = token
	}
}

// ClientAccessToken ...
func ClientAccessTokenProperty(property *AccessTokenProperty) ClientOption {
	return func(obj *Client) {
		obj.accessToken = NewAccessToken(property)
	}
}

// ClientSafeCert ...
func ClientSafeCert(property *SafeCertProperty) ClientOption {
	return func(obj *Client) {
		cfg, e := property.Config()
		if e != nil {
			logrus.Errorf("ClientSafeCert err:%+v", e)
			return
		}
		obj.TLSConfig = cfg

	}
}

// ClientTLSConfig ...
func ClientTLSConfig(config *tls.Config) ClientOption {
	return func(obj *Client) {
		obj.TLSConfig = config

	}
}

// ClientBodyType ...
func ClientBodyType(bt BodyType) ClientOption {
	return func(obj *Client) {
		obj.BodyType = bt
	}
}

// SandboxOption ...
type SandboxOption func(obj *Sandbox)

// SandboxSubID ...
func SandboxSubID(mch, app string) SandboxOption {
	return func(obj *Sandbox) {
		obj.subAppID = app
		obj.subMchID = mch
	}
}

// JSSDKOption ...
type JSSDKOption func(obj *JSSDK)

// JSSDKAccessTokenProperty ...
func JSSDKAccessTokenProperty(property *AccessTokenProperty) JSSDKOption {
	return func(obj *JSSDK) {
		obj.accessToken = NewAccessToken(property, AccessTokenKey(accessTokenKey), AccessTokenURL(accessTokenURLSuffix))
	}
}

// JSSDKAccessToken ...
func JSSDKAccessToken(token *AccessToken) JSSDKOption {
	return func(obj *JSSDK) {
		obj.accessToken = token
	}
}

// JSSDKSubAppID ...
func JSSDKSubAppID(id string) JSSDKOption {
	return func(obj *JSSDK) {
		obj.subAppID = id
	}
}

// JSSDKURL ...
func JSSDKURL(url string) JSSDKOption {
	return func(obj *JSSDK) {
		obj.url = url
	}
}
