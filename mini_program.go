package wego

import (
	"github.com/godcong/wego/app/mini"
)

//OfficialAccount 小程序*/
func MiniProgram() *mini.Program {
	return App().MiniProgram("mini_program.default")
}
