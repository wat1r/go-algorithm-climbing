package test

import (
	"fmt"
	"testing"
)

func Test222(t *testing.T) {
	fmt.Print("end")
}

func TestGoto1(t *testing.T) {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func TestFor1(t *testing.T) {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for i, c := range str {
		fmt.Printf("Character on position %d is: %c \n", i, c)
	}
}
