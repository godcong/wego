package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

/*log types */
const (
	OFF = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

var logList = map[string]int{
	"OFF":   OFF,
	"FATAL": FATAL,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"ALL":   ALL,
}

/*Log Log*/
type Log struct {
	//level = 'debug'
	Level string
	//file = 'logs/wechat.log'
	File string
}

var debug = false
var logs = Log{
	Level: "debug",
	File:  "config.toml",
}

func initLog(l Log, d bool) {
	debug = d
	logs = l
	if IsDebug() {
		i := strings.LastIndexAny(logs.File, "/")
		y := strings.LastIndexAny(logs.File, ".")
		r := []rune(logs.File)
		e := os.Rename(logs.File, string(r[:y])+"_"+time.Now().Format("060102150405")+string(r[y:]))
		os.MkdirAll(string(r[:i]), os.ModePerm)
		f, e := os.OpenFile(logs.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		if e != nil {
			log.Println("cannot open file: " + logs.File)
		}
		log.SetFlags(log.LstdFlags)
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	}

}

/*InitLog log init */
func InitLog(l Log, d bool) {
	initLog(l, d)
}

/*DebugOn turn on Debug */
func DebugOn() {
	debug = true
}

/*DebugOff turn off Debug */
func DebugOff() {
	debug = false
}

/*IsDebug check IsDebug */
func IsDebug() bool {
	return debug
}

/*Println output Println log */
func Println(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Println(fmt.Sprintf("%s|%d", f, l), v)
}

/*Print output Print log */
func Print(v ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	log.Print(fmt.Sprintf("%s|%d", f, l), v)
}

/*Debug output Debug log */
func Debug(v ...interface{}) {
	if DEBUG <= logs.LevelInt() || IsDebug() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[Debug]%s|%d", f, l), v)
	}
}

/*Error output Error log */
func Error(v ...interface{}) {
	if ERROR <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[ERROR]%s|%d", f, l), v)
	}
}

/*Info output Info log */
func Info(v ...interface{}) {
	if INFO <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[INFO]%s|%d", f, l), v)
	}

}

/*Warn output Warn log */
func Warn(v ...interface{}) {
	if WARN <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[WARN]%s|%d", f, l), v)
	}

}

/*Fatal output fatal log */
func Fatal(v ...interface{}) {
	if FATAL <= logs.LevelInt() {
		_, f, l, _ := runtime.Caller(1)
		log.Println(fmt.Sprintf("[FATAL]%s|%d", f, l), v)
	}

}

/*LevelInt 获取level */
func (l *Log) LevelInt() (i int) {
	if v, b := logList[strings.ToUpper(l.Level)]; b {
		i = v
	}
	return i
}
