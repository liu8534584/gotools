package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"test.liuda.com/gotest/utils/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s/%s%s/", setting.AppPath, setting.AppSetting.LogSavePath, time.Now().Format("20060102"))
}

// 获取完整路径
func getLogFileFullPath(fileName string) string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s", fileName, setting.AppSetting.LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

//打开文件 如果不存在则创建
func OpenLogFile(f string) *os.File {
	_, err := os.Stat(f)
	switch {
	case os.IsNotExist(err):
		mkDir(filepath.Dir(f))
	case os.IsPermission(err):
		log.Fatalln("Permission %v", err)
	}
	handle, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

// 创建文件夹
func mkDir(filePath string) {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
