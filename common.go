package gotools

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 获取文件内容 / 获取url地址内容
func FileGetContents(urls string) ([]byte, error) {
	if strings.HasPrefix(urls, "http") {
		return GetResponse(urls)
	}
	if !FileExists(urls) {
		return nil, errors.New("文件不存在")
	}
	bytes, err := ioutil.ReadFile(urls)
	if err != nil {
		log.Printf("read file error : '%v'", err)
		return nil, err
	}
	return bytes, nil
}

/*
*
写入文件
*/
func FilePutContent(filename string, data string) bool {
	var f *os.File
	var err error
	dirName := filepath.Dir(filename)
	if !NewFile().CheckExist(dirName) {
		_ = NewFile().MakeDir(dirName)
	}
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(2, "FilePutContent OpenFile fail : '%v'", err)
		return false
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return false
	}
	return true
}

/*
*
json转map
*/
func JsonDecode(jsonStr string) map[string]interface{} {
	data := []byte(jsonStr)
	var dat map[string]interface{}
	_ = json.Unmarshal(data, &dat)
	return dat
}

/*
*
json格式化
*/
func JsonEncode(m map[string]interface{}) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return b
}

/*
*
构造成功返回数据
*/
func JsonSuccessReturn(c *fiber.Ctx, d interface{}) {
	var resultDataMap map[string]interface{} /*创建集合 */
	resultDataMap = make(map[string]interface{})
	resultDataMap["code"] = 0
	resultDataMap["data"] = d
	resultDataMap["message"] = "ok"
	startTime := c.Locals("beginTime")
	if startTime != "" {
		resultDataMap["execTime"] = (cast.ToFloat64(time.Now().UnixMilli()) - cast.ToFloat64(startTime)) / 1000
	}
	resultDataMap["time"] = time.Now().Unix()
	c.JSON(resultDataMap)
}

// 成功返回空
func JsonSuccessReturnNull(c *fiber.Ctx) {
	var resultDataMap map[string]interface{} /*创建集合 */
	resultDataMap = make(map[string]interface{})
	resultDataMap["code"] = 0
	resultDataMap["message"] = "ok"
	resultDataMap["time"] = time.Now().Unix()
	c.JSON(resultDataMap)
}

/*
*
构造失败返回数据
*/
func JsonFailReturn(c *fiber.Ctx, code int, message string) {
	var resultDataMap map[string]interface{} /*创建集合 */
	resultDataMap = make(map[string]interface{})
	resultDataMap["code"] = code
	resultDataMap["data"] = make(map[string]interface{})
	resultDataMap["message"] = message
	resultDataMap["time"] = time.Now().Unix()
	c.JSON(resultDataMap)
}

/*
*
md5加密
*/
func Md5(str string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(str))                      //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}

/*
*
判断文件是否存在
*/
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func UrlEncode(encodeString string) string {
	ss := url.QueryEscape(encodeString)
	fmt.Println(ss)
	return encodeString
}

func UrlDecode(encodeString string) string {
	return ""
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func GzipAndBase64Encode(encodeStr []byte) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(encodeStr); err != nil {
		return "", err
	}
	if err := gz.Flush(); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func GetLocalIP() (IpAddr string) {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		IpAddr = "localhost"
		return
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return
			}
		}
	}
	IpAddr = "localhost"
	return
}

func GetLocalHostname() string {
	if hostname, err := os.Hostname(); err == nil {
		return hostname
	}
	return ""
}
