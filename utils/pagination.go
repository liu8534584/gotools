package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"test.liuda.com/gotest/consts"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * consts.PageSize
	}
	return result
}