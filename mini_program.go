package wego

import (
	"github.com/godcong/wego/app/mini"
)

// MiniProgram ...
func MiniProgram() *mini.Program {
	return App().MiniProgram("mini_program.default")
}
