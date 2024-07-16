package log

import (
	"github.com/zhangyiming748/lumberjack"
	"log"
	"os"
	"strings"
	"telegraph/constant"
)

func SetLog() {
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{constant.GetRoot(), "telegraph.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}
	//consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	//log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetOutput(fileLogger)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
