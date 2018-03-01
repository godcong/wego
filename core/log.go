package core

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var debug = false

func initLog(system System) {
	debug = system.Debug
	if IsDebug() {
		i := strings.LastIndexAny(system.Log.File, "/")
		y := strings.LastIndexAny(system.Log.File, ".")
		r := []rune(system.Log.File)
		e := os.Rename(system.Log.File, string(r[:y])+"_"+time.Now().Format("060102150405")+string(r[y:]))
		os.MkdirAll(string(r[:i]), os.ModePerm)
		f, e := os.OpenFile(system.Log.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if e != nil {
			log.Println("cannot open file: " + system.Log.File)
		}
		log.SetFlags(log.LstdFlags)
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	}

}

func DebugOn() {
	debug = true
}

func DebugOff() {
	debug = false
}

func IsDebug() bool {
	return debug
}

func Println(v ...interface{}) {
	if IsDebug() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(f, "|", l, "|", v)
	}
}

func Debug(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Println("Debug:", f, "|", l, "|", v)

}

func Error(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Println("ERROR:", f, "|", l, "|", v)
}

func Info(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Println("INFO:", f, "|", l, "|", v)
}

func Print(v ...interface{}) {
	if IsDebug() {
		_, f, l, _ := runtime.Caller(1)
		log.Print(f, "|", l, "|", v)
	}
}
