package core

var sandbox = false

//SandBoxOn open sandbox
func SandBoxOn() {
	sandbox = true
}

//SandBoxOff close sandbox
func SandBoxOff() {
	sandbox = false
}

//UseSandBox check sandbox
func UseSandBox() bool {
	return sandbox
}
