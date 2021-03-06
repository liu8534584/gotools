package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	contents, err := ioutil.ReadAll(f)
	return len(contents), err
}

//获取文件后缀名
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//检测文件是否存在
func CheckExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsExist(err)
}

//检测是否有权限
func CheckPermission(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsPermission(err)
}

//检测是否创建文件夹
func CheckIsNotMakeDir(fileName string) error {
	if exist := CheckExist(fileName); exist == false {
		if err := MakeDir(fileName); err != nil {
			return err
		}
	}
	return nil
}

//创建级联目录
func MakeDir(fileName string) error {
	err := os.MkdirAll(fileName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//打开文件
func Open(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(fileName, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
