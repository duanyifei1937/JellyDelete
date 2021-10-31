package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"jellydelete/global"
	"jellydelete/internal/model"
	"jellydelete/pkg/convert"
)

func GetInitPattern(c *gin.Context) {
	l := convert.StrTo(c.Query("level")).MustInt()

	// weight: level难度； g:64string； m: uuid; p: 8*8string信息；
	weight := global.AppSetting.Level[l]
	g := initGrid(weight, 64)
	m := md5V(g)
	p := outputPattern(g)

	// db := global.DBEngine
	j := model.Jelly{
		UUID: m,
		Exec: 0,
		Comb: g,
	}

	global.DBEngine.Create(&j)

	c.String(200, m+"\n"+p)
}

func MovePattern(c *gin.Context) {
	uuid := convert.StrTo(c.Query("sessionId")).String()
	r0, c0, r1, c1 := convert.StrTo(c.Query("row0")).MustInt(), convert.StrTo(c.Query("col0")).MustInt(), convert.StrTo(c.Query("row1")).MustInt(), convert.StrTo(c.Query("col1")).MustInt()

	selectElements := getElementList(r0, c0, r1, c1)

	j := model.Jelly{
		UUID: uuid,
	}
	// SELECT comb FROM jelly WHERE uuid="89b89b17a2f502d52cf3378c6cf902d3";
	// 根据uri arg uuid, 查询comb
	global.DBEngine.Where("uuid = ?", uuid).Last(&j)
	// global.DBEngine.Last(&j).Where("uuid = ?", uuid)

	// 查询到图案转换为8*8 plate
	plate := stringToPlate(j.Comb)

	for e := range selectElements {
		er := e[0].(int)
		ec := e[1].(int)
		switch plate[er][ec] {
		// switch plate[1][2] {
		case "B":
			moveBElement(plate, []int{er, ec})
		case "H":
			moveHElement(plate, []int{er, ec})
		case "V":
			moveVElement(plate, []int{er, ec})
		case "S":
			moveSElement(plate, []int{er, ec})
		}
	}

	// outp := outputPattern(plateToString(plate))
	// c.String(200, outp)

	// 下沉 + 补位
	plate = elementSink(plate)

	buwei := elementFillIn(plate)
	fmt.Println("--buwei: ", *buwei)

	// ss := plateToString(buwei)
	// fmt.Println("xxx===xxx: ", ss)

	// nplate := elementFillIn(elementSink(plate))
	// outp := outputPattern(plateToString(nplate))
	// c.String(200, outp)

}
