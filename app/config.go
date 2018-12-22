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

type OfficialAccountConfig struct {
	SandboxConfig
}

// Config ...
type Config struct {
}
