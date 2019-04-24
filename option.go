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

// ClientOption ...
type ClientOption func(obj *Client)

// ClientContext ...
func ClientContext(ctx context.Context) ClientOption {
	return func(obj *Client) {
		obj.context = ctx
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
