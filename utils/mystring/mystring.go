package mystring

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
	"reflect"
	"regexp"
	"strings"
	"test.liuda.com/gotest/utils/myerr"
)

//首字母大写
func UcFirst(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func IsGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func preNUm(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i int = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}
func IsUtf8(data []byte) bool {
	for i := 0; i < len(data); {
		if data[i]&0x80 == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if data[i]&0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func U2S(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		_ = binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

//去掉多余信息
func StringRemove(html string, removeStringList string) string {

	if removeStringList == "" {
		html = strings.ReplaceAll(html, "<p>", "    ")
		html = strings.ReplaceAll(html, "</p>", "\r\n")
		html = TrimHtml(html)
		return html
	}

	jsonData, err := simplejson.NewJson([]byte(removeStringList))
	if err != nil {
		return html
	}

	jsonArr, err := jsonData.Array()
	if err != nil {
		return html
	}

	var replaceMap map[string]interface{}
	for _, v := range jsonArr {
		replaceMap = cast.ToStringMap(v)
		switch replaceMap["type"].(string) {
		case "string":
			html = strings.ReplaceAll(html, replaceMap["old"].(string), replaceMap["new"].(string))
		case "regexp":
			re, _ := regexp.Compile(replaceMap["old"].(string))
			html = re.ReplaceAllString(html, replaceMap["new"].(string))
		}
	}

	html = strings.ReplaceAll(html, "<p>", "    ")
	html = strings.ReplaceAll(html, "</p>", "\r\n")
	html = TrimHtml(html)

	return html
}

//获取正则表达式匹配内容
func GetRegexpContents(html []byte, rule string) (string, error) {
	reg := regexp.MustCompile(rule)
	res := reg.FindAllString(string(html), -1)
	return res[0], nil
}

func GetRegexpContentsList(html []byte, rule string) ([]string, error) {
	reg := regexp.MustCompile(rule)
	res := reg.FindAllString(string(html), -1)
	return res, nil
}

func GetCssSelectorContentsList(html []byte, rule string) ([]string, error) {
	var urlList []string
	var attr string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	var logInfo string
	if err != nil {
		logInfo = fmt.Sprintf("转换成goquery对象失败,err:%v", err)
		return urlList, myerr.NewError(4001, logInfo)
	}

	if strings.Index(rule, ":") != -1 {
		attrInfo := strings.Split(rule, ":")
		rule = attrInfo[0]
		attr = attrInfo[1]
	}

	doc.Find(rule).Each(func(i int, selection *goquery.Selection) {
		if attr != "" {
			val, b := selection.Attr("href")
			if b {
				urlList = append(urlList, val)
			}
		} else {
			val := selection.Text()

			urlList = append(urlList, val)
		}
	})

	return urlList, nil

	//if err != nil {
	//	logInfo = fmt.Sprintf("rule:%s,查询不到内容,err:%v", rule, err)
	//}
	//return info, nil
}

//获取css选择器匹配内容
func GetCssSelectorContents(html []byte, rule string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	var logInfo, info string
	if err != nil {
		logInfo = fmt.Sprintf("转换成goquery对象失败,err:%v", err)
		return "", myerr.NewError(400, logInfo)
	}

	if strings.Contains(rule, ":href") {
		rule = strings.Replace(rule, ":href", "", 1)
		info, _ = doc.Find(rule).First().Attr("href")
	} else if strings.Contains(rule, ":src") {
		rule = strings.Replace(rule, ":src", "", 1)
		info, _ = doc.Find(rule).First().Attr("href")
	} else {
		info, err = doc.Find(rule).First().Html()
	}

	if err != nil {
		logInfo = fmt.Sprintf("rule:%s,查询不到内容,err:%v", rule, err)
	}
	return info, nil
}
