package test

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}

/*
多行参数
*/
func Test1(t *testing.T) {
	N := 10
	/* 	for i, j := 0, N; i < j; i, j = i+1, j-1 {
		fmt.Printf("%v, %v\n", i, j)
	} */

	for i := 0; i < N; i++ {
		for j := N; j >= 0 && j > i; j-- {
			fmt.Printf("%v,%v \n", i, j)
		}
	}
}

func TestStr(t *testing.T) {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

func TestStr2(t *testing.T) {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}

}

func Test2(t *testing.T) {
	// 1:
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2:
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	fmt.Println("==============")
	if i < 15 {
		goto START
	}
}

func Test3(t1 *testing.T) {
	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}

func Test1_1(t *testing.T) {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
		fmt.Printf("%d ", v)
	}

}

// func  Test4(t *testing.T) {
// 	for true {

// 	}
// }
