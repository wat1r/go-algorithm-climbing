package test

import (
	"fmt"
	"testing"
)

func TestOne1(t *testing.T) {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	// fmt.Printf("end")
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

}

func min(s ...int) int {

	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min

}

func TestDefer(t *testing.T) {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func TestMain(t *testing.T) {
	// a()
	// b()
	t1()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func t1() {
	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// s := "a"
	// for s != "aaaaaa" {
	// 	fmt.Println("%s", s)
	// 	s = s + "a"
	// }

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

}
