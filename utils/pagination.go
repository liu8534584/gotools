package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/liu8534584/gotools/consts"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * consts.PageSize
	}
	return result
}
