package wego

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var debug = false

func initLog(system System) {
	debug = system.Debug
	if Debug() {
		os.Rename(system.Log.File, time.Now().Format("060102150405")+"_"+system.Log.File)

		i := strings.LastIndexAny(system.Log.File, "/")
		r := []rune(system.Log.File)
		os.MkdirAll(string(r[:i]), os.ModePerm)

		f, e := os.OpenFile(system.Log.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if e != nil {
			log.Println("cannot open file: " + system.Log.File)
		}
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	}

}

func DebugOn() {
	debug = true
}

func DebugOff() {
	debug = false
}

func Debug() bool {
	return debug
}

func Println(v ...interface{}) {
	if Debug() {
		log.Println(v)

	}
}

func Print(v ...interface{}) {
	if Debug() {
		log.Print(v)
	}
}
