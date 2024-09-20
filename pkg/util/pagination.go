package util

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
