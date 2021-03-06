package qrcode

import (
	"github.com/boombuler/barcode/qr"
	"test.liuda.com/gotest/utils/common"
	"test.liuda.com/gotest/utils/setting"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const EXT_IMAGE = ".jpg"

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    EXT_IMAGE,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodePath
}

func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetQrCodePath()
}

func GetQrCodeFileName(value string) string {
	return common.Md5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}
