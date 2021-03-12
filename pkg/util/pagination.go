
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/GaoHaHa-IronMan/go-gin-example/pkg"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * pkg.PageSize
	}

	return result
}