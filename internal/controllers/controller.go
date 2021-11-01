package controllers

import (
	"log"

	"github.com/gin-gonic/gin"

	"jellydelete/global"
	"jellydelete/internal/model"
	"jellydelete/pkg/convert"
)

func GetInitPattern(c *gin.Context) {
	l, err := convert.StrTo(c.Query("level")).MustInt()
	if err != nil {
		global.Logger.Errorf("level_arg not int, bad request", err)
		c.String(400, "INVALID PARAMS")
	}

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

	result := global.DBEngine.Create(&j)
	if result.Error != nil {
		global.Logger.Fatalf("write to db faild: ", result.Error)
	}

	c.String(200, m+"\n"+p)
}

func MovePattern(c *gin.Context) {
	uuid := convert.StrTo(c.Query("sessionId")).String()
	r0, e1 := convert.StrTo(c.Query("row0")).MustInt()
	c0, e2 := convert.StrTo(c.Query("col0")).MustInt()
	r1, e3 := convert.StrTo(c.Query("row1")).MustInt()
	c1, e4 := convert.StrTo(c.Query("col1")).MustInt()

	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || !(r0 >= 0 && r0 <= 7) || !(c0 >= 0 && c0 <= 7) || !(r1 >= 0 && r1 <= 7) || !(c1 >= 0 && c1 <= 7) {
		global.Logger.Errorf("get row or col arg faild.")
		c.String(400, "INVALID PARAMS")
	}
	selectElements := getElementList(r0, c0, r1, c1)

	j := model.Jelly{
		UUID: uuid,
	}
	// SELECT comb FROM jelly WHERE uuid="89b89b17a2f502d52cf3378c6cf902d3";
	// 根据uri arg uuid, 查询comb
	// global.DBEngine.Last(&j).Where("uuid = ?", uuid)
	result := global.DBEngine.Where("uuid = ?", uuid).Last(&j)
	if result.RowsAffected == 0 {
		global.Logger.Errorf("sessionID not found: %s", uuid)
		c.String(400, "arg_sessionId not found.")
	}

	// 查询到图案转换为8*8 plate
	plate := stringToPlate(j.Comb)

	// 对选中框内的元素判断、移除
	for e := range selectElements {
		er := e[0].(int)
		ec := e[1].(int)
		switch plate[er][ec] {
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

	// space element 下沉、补位
	nplate := elementFillIn(elementSink(plate))
	nstring64 := plateToString(nplate)

	// 补位后的结果写入db
	j = model.Jelly{
		UUID: uuid,
		Exec: j.Exec + 1,
		Comb: nstring64,
	}

	// global.DBEngine.Where("uuid = ?", uuid).Last(&j)
	result = global.DBEngine.Create(&j)
	if result.Error != nil {
		log.Fatalln("string64 to db faild: ", result.Error)
		c.String(500, "write to db faild")
	}

	outp := outputPattern(nstring64)
	c.String(200, outp)

}
