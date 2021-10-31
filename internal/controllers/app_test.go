package controllers

import (
	"fmt"
	"testing"
)

func TestGetElementList(t *testing.T) {
	el := getElementList(0, 1, 1, 2)
	for e := range el {
		t.Log(e[0], e[1])
		t.Logf("%T", e[0])
		// t.Log(e[1])
		t.Log("--")
	}
	fmt.Printf("%T", el)
}

func TestStringToPlate(t *testing.T) {
	s := "BBBVBBBBVVBBBBBBBHBBBVBBHBVHBBBVBBBBBBBBBBBBBBBBVBBBBBBBVBHBBBHV"
	p := stringToPlate(s)
	t.Log(p)
}

func TestPlateToString(t *testing.T) {
	p := [8][8]string{
		{"A", "B", "C"},
	}
	fmt.Println("--")
	fmt.Println(plateToString(&p))
	fmt.Println("--")
}

func TestElemenSink(t *testing.T) {
	s := "AB DEFGHIJKL NOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKL"
	p := stringToPlate(s)
	p1 := elementSink(p)

	fmt.Println(*p)
	fmt.Println(*p1)
	fmt.Printf("==%s==\n", p[0][0])
	fmt.Printf("==%s==", p1[0][0])
	// fmt.Printf("-%s-", string(p1[3][0]))

	// if string(p1[3][0]) == "" {
	// 	fmt.Println("ok")
	// }

	// fmt.Printf("%s", string(p1[3][1]))

}
