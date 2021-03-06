package upload

import (
	"fmt"
	"github.com/liu8534584/gotools/utils/common"
	"github.com/liu8534584/gotools/utils/file"
	"github.com/liu8534584/gotools/utils/logging"
	"github.com/liu8534584/gotools/utils/setting"
	"mime/multipart"
	"os"
	"strings"
)

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullName(imageFileName string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + imageFileName
}

func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.Trim(name, "")
	fileName = common.Md5(fileName)
	return fileName + ext
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Waring(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		logging.Waring("checkImage Getwd err :%v", err)
		return fmt.Errorf("os.Getwd err:%v", err)
	}
	err = file.CheckIsNotMakeDir(dir + "/" + fileName)
	if err != nil {
		logging.Waring("checkImage checkIsNotMakeDir err : %v", err)
		return fmt.Errorf("isNotMakeDir err:%v", err)
	}
	perm := file.CheckPermission(fileName)
	if perm == true {
		logging.Waring("checkImage CheckPermission src : %v", fileName)
		return fmt.Errorf("checkImage CheckPermission src : %v", fileName)
	}
	return nil
}
