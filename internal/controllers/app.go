package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"jellydelete/global"
	"math/rand"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/schwarmco/go-cartesian-product"
)

// 方法实现集合

// 生成64位随机string, 以时间戳作为random种子;
// 可控制BHVS出现的权重控制难度级别；
// 可控制产生的位数
func initGrid(weight []uint, n int) string {
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

// 64位string转换为定义8*8二维数组
func stringToPlate(s string) *[8][8]string {
	var m [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			m[i][j] = string(s[i*8+j])
		}
	}
	return &m
}

//8*8二维数组转换为64string
func plateToString(p *[8][8]string) string {
	var s []string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			s = append(s, p[i][j])
		}
	}
	return strings.Join(s, "")
}

// 两个点定位，输出所有点坐标, row, col
func getElementList(r1, c1, r2, c2 int) chan []interface{} {
	// row中间数字slice
	rl := []interface{}{}
	if r1 > r2 {
		r1, r2 = r2, r1
	}
	for i := r1; i <= r2; i++ {
		rl = append(rl, i)
	}

	// col中间数字slice
	cl := []interface{}{}
	if c1 > c2 {
		c1, c2 = c2, c1
	}
	for i := c1; i <= c2; i++ {
		cl = append(cl, i)
	}

	// 笛卡尔积
	result := cartesian.Iter(rl, cl)
	return result
}

// 读取plate中特定元素
func getElement(p *[8][8]string, e []int) string {
	return p[e[0]][e[1]]
}

// B元素移除动作
func moveBElement(p *[8][8]string, e []int) *[8][8]string {
	p[e[0]][e[1]] = " "
	return p
}

// H元素移除动作
func moveHElement(p *[8][8]string, e []int) *[8][8]string {
	for i := 0; i < 8; i++ {
		p[e[0]][i] = " "
	}
	return p
}

// V元素移除动作
func moveVElement(p *[8][8]string, e []int) *[8][8]string {
	for i := 0; i < 8; i++ {
		p[i][e[1]] = " "
	}
	return p
}

// S元素移除动作
func moveSElement(p *[8][8]string, e []int) {
	n, m := e[0], e[1]

	for i := n - 1; i <= n+1; i++ {
		for j := m - 1; j <= m+1; j++ {
			if i < 0 || j < 0 || i >= 8 || j >= 8 || (i == n && j == m) {
				continue
			} else {
				p[i][j] = " "
			}
		}
	}
}

// nil元素下沉
func elementSink(p *[8][8]string) *[8][8]string {
	var np [8][8]string
	// for i := 0; i < 8; i++ {
	// 	var j1 int = 7
	// 	for j := 7; j >= 0; j-- {
	// 		if p[i][j] != " " {
	// 			np[i][j1] = p[i][j]
	// 			j1--
	// 		}
	// 	}
	// }

	// 移动过程从左下角开始，先列，后行
	for j := 0; j < 8; j++ {
		i1 := 7
		for i := 7; i >= 0; i-- {
			// log.Println("xxx: ", p[i][j])
			if p[i][j] != " " {
				np[i1][j] = p[i][j]
				i1--
			}
		}
	}

	// 空格替换
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if np[i][j] == "" {
				np[i][j] = " "
			}
		}
	}

	return &np
}

// nil元素补位
func elementFillIn(p *[8][8]string) *[8][8]string {

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if p[i][j] == " " {
				p[i][j] = genRandomString(global.AppSetting.Level[999])
			}
		}
	}
	return p
}

// 生成随机字符
func genRandomString(weight []uint) string {
	rand.Seed(time.Now().UTC().UnixNano())

	chooser, _ := wr.NewChooser(
		wr.Choice{Item: "B", Weight: weight[0]},
		wr.Choice{Item: "H", Weight: weight[1]},
		wr.Choice{Item: "V", Weight: weight[2]},
		wr.Choice{Item: "S", Weight: weight[3]},
	)

	return chooser.Pick().(string)

}
