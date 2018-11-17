package mini

import "github.com/godcong/wego/core"

/*OpenData OpenData */
type OpenData struct {
	*Program
}

func newOpenData(program *Program) interface{} {
	return &OpenData{
		Program: program,
	}
}

// NewOpenData ...
func NewOpenData(config *core.Config) *OpenData {
	return newOpenData(NewMiniProgram(config)).(*OpenData)
}
