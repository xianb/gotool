package utils

import (
	"fmt"
	"testing"
)

func TestContain(t *testing.T) {
	testMap()
	testArray()
	testSlice()
}

func testArray() {
	a := 1
	b := [3]int{1, 2, 3}

	fmt.Println(Contain(a, b))

	c := "a"
	d := [4]string{"b", "c", "d", "a"}
	fmt.Println(Contain(c, d))

	e := 1.1
	f := [4]float64{1.2, 1.3, 1.1, 1.4}
	fmt.Println(Contain(e, f))

	g := 1
	h := [4]interface{}{2, 4, 6, 1}
	fmt.Println(Contain(g, h))

	i := [4]int64{}
	fmt.Println(Contain(a, i))
}

func testSlice() {
	a := 1
	b := []int{1, 2, 3}

	fmt.Println(Contain(a, b))

	c := "a"
	d := []string{"b", "c", "d", "a"}
	fmt.Println(Contain(c, d))

	e := 1.1
	f := []float64{1.2, 1.3, 1.1, 1.4}
	fmt.Println(Contain(e, f))

	g := 1
	h := []interface{}{2, 4, 6, 1}
	fmt.Println(Contain(g, h))

	var i []int64
	fmt.Println(Contain(a, i))
}

func testMap() {
	var a = map[int]string{1: "1", 2: "2"}
	fmt.Println(Contain(3, a))

	var b = map[string]int{"1": 1, "2": 2}
	fmt.Println(Contain("1", b))

	var c = map[string][]int{"1": {1, 2}, "2": {2, 3}}
	fmt.Println(Contain("6", c))
}
