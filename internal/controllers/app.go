package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
)

// 方法实现集合

// 生成64位随机string, 以时间戳作为random种子;
// 可控制BHVS出现的权重控制难度级别；
func initGrid(weight []uint) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var results []string

	chooser, _ := wr.NewChooser(
		wr.Choice{Item: "B", Weight: weight[0]},
		wr.Choice{Item: "H", Weight: weight[1]},
		wr.Choice{Item: "V", Weight: weight[2]},
		wr.Choice{Item: "S", Weight: weight[3]},
	)
	for i := 0; i < 64; i++ {
		results = append(results, chooser.Pick().(string))
	}

	return strings.Join(results, "")
}

// 根据64位string m5作为关卡唯一值
func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 输入64位string, 输出为8*8表盘信息
func outputPattern(s string) string {
	var g string
	for i := 1; i <= 8; i++ {
		g += s[8*(i-1) : 8*(i)]
		g += "\n"
	}
	return g
}
