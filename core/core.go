package core

//Core core set
type Core struct {
	//config
}

var isSandbox = false

//SandBoxOn open sandbox
func SandBoxOn() {
	isSandbox = true
}

//SandBoxOff close sandbox
func SandBoxOff() {
	isSandbox = false
}

//IsSandBox check sandbox
func IsSandBox() bool {
	return isSandbox
}
