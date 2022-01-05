package test

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d , j is %d\n", i, j)
		}
	}
}

func Test1_2(t *testing.T) {
	i := 0
HERE:
	fmt.Printf("%d\n", i)
	i++
	if i == 5 {
		return
	}
	goto HERE

}

func Test1_3(t *testing.T) {
	// 	a := 1
	// 	goto TARGET // compile error
	// 	b := 9
	// TARGET:
	// 	b += a
	// 	fmt.Printf("a is %v *** b is %v", a, b)

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}

func Test6(t *testing.T) {
	fmt.Printf("end")
}

// func  Test4(t *testing.T) {

// }

// fun
