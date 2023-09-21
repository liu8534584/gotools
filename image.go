package gotools

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"mime/multipart"
	"os"
	"strings"
)

type Image struct {
	SavePath  string
	PrefixUrl string
	AllowExts []string
	MaxSize   int
}

func NewImage(path, prefixUrl string, allowExts []string, maxSize int) *Image {
	return &Image{
		SavePath:  path,
		PrefixUrl: prefixUrl,
		AllowExts: allowExts,
		MaxSize:   maxSize,
	}
}

func (m *Image) GetImagePath() string {
	return m.SavePath
}

func (m *Image) GetImageFullName(imageFileName string) string {
	return m.PrefixUrl + "/" + m.GetImagePath() + imageFileName
}

func (m *Image) GetImageName(name string) string {
	ext := NewFile().GetExt(name)
	fileName := strings.Trim(name, "")
	fileName = goutils.Md5(fileName)
	return fileName + ext
}

func (m *Image) GetImageFullPath() string {
	return m.SavePath + m.GetImagePath()
}

func (m *Image) CheckImageExt(fileName string) bool {
	ext := NewFile().GetExt(fileName)
	for _, allowExt := range m.AllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func (m *Image) CheckImageSize(f multipart.File) bool {
	size, err := NewFile().GetSize(f)
	if err != nil {
		return false
	}
	return size <= m.MaxSize
}

func (m *Image) CheckImage(fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err:%v", err)
	}
	err = NewFile().CheckIsNotMakeDir(dir + "/" + fileName)
	if err != nil {
		return fmt.Errorf("isNotMakeDir err:%v", err)
	}
	perm := NewFile().CheckPermission(fileName)
	if perm == true {
		return fmt.Errorf("checkImage CheckPermission src : %v", fileName)
	}
	return nil
}
