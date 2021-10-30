package controllers

import (
	"github.com/gin-gonic/gin"

	"jellydelete/global"
	"jellydelete/pkg/convert"
)

// 吐出初始pattern:
// 按权重随机，先吐出来，在保存；
func GetInitPattern(c *gin.Context) {
	l := convert.StrTo(c.Query("level")).MustInt()

	// weight: level难度； g:64string； m: uuid; p: 8*8string信息；
	weight := global.AppSetting.Level[l]
	g := initGrid(weight)
	m := md5V(g)
	p := outputPattern(g)

	// content := m + "\n" + p
	c.String(200, m+"\n"+p)
}
