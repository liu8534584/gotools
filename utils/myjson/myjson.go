package myjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"test.liuda.com/gotest/utils/common"
)

func createStruct(jsonStr string, structName string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("转化错误:", err)
	}
	var buffer bytes.Buffer
	buffer.WriteString("type ")
	buffer.WriteString(structName)
	buffer.WriteString(" struct {\n")
	for k, v := range m {
		runes := []rune(k)
		buffer.WriteString(strings.ToUpper(string(runes[0])))
		buffer.WriteString(string(runes[1:]))
		buffer.WriteString("   ")
		buffer.WriteString(reflect.TypeOf(v).String())
		buffer.WriteString("     `json:\"")
		buffer.WriteString(k)
		buffer.WriteString("\"`")
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	fmt.Println(buffer.String())
}

func readJson(file string) ([]byte, error) {
	return common.FileGetContents(file)
}

func CreateJsonStruct(input string, output string) {
	r, err := readJson(input)
	if err != nil {
		return
	}
	createStruct(string(r), output)

}
