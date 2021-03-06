package utils_filter

import (
	"github.com/astaxie/beego/validation"
)


func checkMobile(mobile string)  {
	f := make(map[string]interface{})
	f["mobile"] = mobile
	valid := validation.Validation{}
	valid.Phone(f,"mobile")
}
