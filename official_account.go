package wego

// OfficialAccount ...
type OfficialAccount struct {
	*OfficialAccountProperty
	client   *Client
	property *Property
	option   *OfficialAccountOption
}

// OfficialAccountOption ...
type OfficialAccountOption struct {
	RemoteAddress string
	LocalHost     string
	UseSandbox    bool
	Sandbox       *SandboxProperty
	NotifyURL     string
	RefundURL     string
}
