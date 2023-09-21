package upload

import (
	"bytes"
	"context"
	"fmt"
	"github.com/liu8534584/gotools/utils/logging"
	"github.com/liu8534584/gotools/utils/myredis"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

type Qiniu struct {
	cfg       storage.Config
	AccessKey string
	SecretKey string
	Bucket    string
}

func NewQiniu(accessKey, secretKey, bucket string) *Qiniu {
	return &Qiniu{AccessKey: accessKey, SecretKey: secretKey, Bucket: bucket}
}

func (q *Qiniu) getUploadToken() string {

	token, err := myredis.Client.Get("key_qiniu_upload_token").Result()
	if err == nil {
		return token
	}
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	myredis.Client.Set("key_qiniu_upload_token", upToken, 7200)
	return upToken
}

func (q *Qiniu) getCfg() {
	q.cfg = storage.Config{}
	// 空间对应的机房
	q.cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	q.cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	q.cfg.UseCdnDomains = false
}

// 字节流上传
func (q *Qiniu) UploadBytes(k string, data []byte) (string, error) {
	q.getCfg()

	formUploader := storage.NewFormUploader(&q.cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, q.getUploadToken(), k, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {

		return "", err
	}

	return k, nil
}

// 上传文件
func (q *Qiniu) UploadFile(k string, filepath string) (string, error) {
	q.getCfg()
	formUploader := storage.NewFormUploader(&q.cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, q.getUploadToken(), k, filepath, &putExtra)
	if err != nil {
		logInfo := fmt.Sprintf("文件上传失败,filepath:%s,err:%v", filepath, err)
		logging.Error(logInfo)
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash)
	return ret.Key, nil
}
