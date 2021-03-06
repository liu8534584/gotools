package e

const (
	SUCCESS = 200

	BaseError = 400

	ParamErrCode = 401

	AuthErrCode = 500

	ErrorUploadSaveImageFail    = 30001
	ErrorUploadCheckImageFail   = 30002
	ErrorUploadCheckImageFormat = 30003

	ErrorJoinRedis = 10001
)

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	BaseError:                   "系统错误",
	ParamErrCode:                "参数错误",
	AuthErrCode:                 "签名验证失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadSaveImageFail:    "图片保存失败",
	ErrorJoinRedis:              "加入队列失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[BaseError]
}
