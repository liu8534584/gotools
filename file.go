package gotools

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}

// GetSize 获取文件大小
func (m *File) GetSize(f multipart.File) (int, error) {
	contents, err := ioutil.ReadAll(f)
	return len(contents), err
}

// GetExt 获取文件后缀名
func (m *File) GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检测文件是否存在
func (m *File) CheckExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsExist(err)
}

// CheckPermission 检测是否有权限
func (m *File) CheckPermission(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsPermission(err)
}

// CheckIsNotMakeDir 检测是否创建文件夹
func (m *File) CheckIsNotMakeDir(fileName string) error {
	if exist := m.CheckExist(fileName); exist == false {
		if err := m.MakeDir(fileName); err != nil {
			return err
		}
	}
	return nil
}

// MakeDir 创建级联目录
func (m *File) MakeDir(fileName string) error {
	err := os.MkdirAll(fileName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Open 打开文件
func (m *File) Open(fileName string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(fileName, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
