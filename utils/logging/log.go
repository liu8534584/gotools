package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"test.liuda.com/gotest/utils/mystring"
	"test.liuda.com/gotest/utils/setting"
)

type Level int

var (
	F *os.File

	//默认前缀
	DefaultPrefix = ""

	DefaultCallerDepth = 2

	logger *log.Logger

	//日志前缀
	logPrefix = ""

	//日志级别
	LevelFlags = []string{"DEBUG", "INFO", "WARING", "ERROR", "FATAL"}
)

type MyLogger struct {
}

const (
	DEBUG Level = iota
	INFO
	WARING
	ERROR
	FATAL
)

func Debug(v ...interface{}) {
	setLevelPrefix(DEBUG)
	writeLog(v)
}

func Info(v ...interface{}) {
	setLevelPrefix(INFO)
	writeLog(v)
}

func Waring(v ...interface{}) {
	setLevelPrefix(WARING)
	writeLog(v)
}

func (m *MyLogger) Print(v ...interface{}) {
	filePath := getLogFileFullPath("sql")
	F = OpenLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	writeLog(v)
}

func Error(v ...interface{}) {
	setLevelPrefix(ERROR)
	writeLog(v)
}

func Fatal(v ...interface{}) {
	setLevelPrefix(FATAL)
	writeLog(v)
}

func writeLog(v ...interface{}) {
	if setting.GetEnv() == "" {
		fmt.Println(v)
		logger.Println(v)
	} else {
		fmt.Println(v)
		logger.Println(v)
	}
}

func D(fileName string, v ...interface{}) {
	setPrefix(fileName, DEBUG)
	writeLog(v)

}

func I(fileName string, v ...interface{}) {
	setPrefix(fileName, INFO)
	writeLog(v)
}

func E(fileName string, v ...interface{}) {
	setPrefix(fileName, ERROR)
	writeLog(v)
}

func W(fileName string, v ...interface{}) {
	setPrefix(fileName, WARING)
	writeLog(v)
}

// 设置写入文件前缀
func setLevelPrefix(level Level) {
	//拿到文件名和行数
	if _, f, line, ok := runtime.Caller(DefaultCallerDepth); ok {
		//获取文件名的相对路径 去掉.go后缀 按照路径拼接文件名
		relativeFile := strings.Replace(f, setting.AppPath, "", -1)
		f = strings.Replace(relativeFile, ".go", "", -1)

		//首字母大写 路径按照_拼接 去掉最前面的_
		sp := strings.Split(f, "/")
		var ss string

		//定义一个值 遇到程序根目录之后的名字才能写进去
		// 例如/var/www/html/gotest/cache/token.go  /var/www/html/gotest 这一段是没有意义的 只要后面的cache/token
		var nn bool
		nn = false
		for k, _ := range sp {
			if nn == true {
				ss += "_" + mystring.UcFirst(sp[k])
			}
			if sp[k] == "gotest" {
				nn = true
			}
		}

		if nn == false {
			for k, _ := range sp {
				ss += "_" + mystring.UcFirst(sp[k])
			}
		}

		runeStr := []rune(ss)
		ss = string(runeStr[1:len(runeStr)])

		//设置文件名
		filePath := getLogFileFullPath(ss)
		F = OpenLogFile(filePath)

		logger = log.New(F, DefaultPrefix, log.LstdFlags)

		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], relativeFile, line)
	} else {
		filePath := getLogFileFullPath(LevelFlags[level])
		F = OpenLogFile(filePath)
		logger = log.New(F, DefaultPrefix, log.LstdFlags)
		logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

func setPrefix(fileName string, level Level) {
	//设置文件名
	filePath := getLogFileFullPath(fileName)
	F = OpenLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)

	//获取实际文件名和行数
	_, f, line, ok := runtime.Caller(DefaultCallerDepth)
	relativeFile := strings.Replace(f, setting.AppPath, "", -1)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], relativeFile, line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
