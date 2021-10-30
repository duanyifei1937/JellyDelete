package main

import (
	"jellydelete/global"
	"jellydelete/internal/model"
)

func main() {
	db := global.DBEngine
	j := model.Jelly{
		UUID: "aaa",
		Exec: 0,
		Comb: "xxx",
	}

	db.Create(&j) // 通过数据的指针来创建
}
