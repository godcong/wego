package app

// SystemConfig ...
type SystemConfig struct {
	Debug bool
	Cache string
}

type LogConfig struct {
	Path string
}

type SandboxConfig struct {
	AppID  string
	Secret string
}

type BaseConfig struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

type HttpConfig struct {
	TimeOut   int
	KeepAlive int
}

type OfficialAccountConfig struct {
	Sandbox SandboxConfig
	def     BaseConfig
}

// Config ...
type Config struct {
}
