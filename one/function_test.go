package test

import (
	"fmt"
	// "stack"
	"testing"
)

// func (st *Stack) Pop() int {
// 	v := 0
// 	for ix := len(st) - 1; ix >= 0; ix-- {
// 		if v = st[ix]; v != 0 {
// 			st[ix] = 0
// 			return v
// 		}
// 	}
// 	return 0
// }

// func  Test4(t *testing.T) {

// }

var num int = 10
var numx2, numx3 int

func TestEntrance(t *testing.T) {
	// fmt.Printf("%d\n", MultiPly3Nums(2, 5, 6))
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func SumProductDiffN(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return
}
