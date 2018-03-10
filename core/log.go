package core

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var debug = false
var logs = Log{
	Level: "debug",
	File:  "config.toml",
}

func initLog(system System) {
	debug = system.Debug
	logs = system.Log
	if IsDebug() {
		i := strings.LastIndexAny(logs.File, "/")
		y := strings.LastIndexAny(logs.File, ".")
		r := []rune(logs.File)
		e := os.Rename(logs.File, string(r[:y])+"_"+time.Now().Format("060102150405")+string(r[y:]))
		os.MkdirAll(string(r[:i]), os.ModePerm)
		f, e := os.OpenFile(logs.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if e != nil {
			log.Println("cannot open_platform file: " + logs.File)
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
	_, f, l, _ := runtime.Caller(1)
	log.Println(fmt.Sprintf("%s|%d", f, l), v)
}

func Print(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Print(fmt.Sprintf("%s|%d", f, l), v)
}

func Debug(v ...interface{}) {
	if DEBUG <= logs.LevelInt() || IsDebug() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[Debug]%s|%d", f, l), v)
	}
}

func Error(v ...interface{}) {
	if ERROR <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[ERROR]%s|%d", f, l), v)
	}
}

func Info(v ...interface{}) {
	if INFO <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[INFO]%s|%d", f, l), v)
	}

}

func Warn(v ...interface{}) {
	if WARN <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[WARN]%s|%d", f, l), v)
	}

}

func Fatal(v ...interface{}) {
	if FATAL <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[FATAL]%s|%d", f, l), v)
	}

}
