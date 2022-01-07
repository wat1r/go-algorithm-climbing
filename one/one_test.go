package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
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

	// for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	// 	s = i+1, j+1, s+"a" {
	// 	fmt.Println("Value of i, j, s:", i, j, s)
	// }

	// str := "hello"
	// c := []byte(str)
	// c[0] = 'c'
	// fmt.Println(string(c))

	// substr := str[2:4] //取头不取尾
	// fmt.Println(substr)

	// for idx, ch := range str {
	// 	fmt.Println(idx, string(ch))
	// }

	// // gives only the bytes:
	// for i:=0; i < len(str); i++ {
	// … = str[i]
	// }
	// // gives the Unicode characters:
	// for ix, ch := range str {
	// …
	// }

	// fmt.Println(utf8.RuneCountInString(str))

	str := "hello 你好"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	// fmt.Println("rune:", len([]rune(str)))

	fmt.Println("rune:", len([]rune(str)))
}

func t2() {
	// var str string = "This is an example of a string"
	// fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	// fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	// var str string = "Hi, I'm Marc, Hi."

	// fmt.Printf("The position of \"Marc\" is: ")
	// fmt.Printf("%d\n", strings.Index(str, "Marc"))

	// fmt.Printf("The position of the first instance of \"Hi\" is: ")
	// fmt.Printf("%d\n", strings.Index(str, "Hi"))
	// fmt.Printf("The position of the last instance of \"Hi\" is: ")
	// fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	// fmt.Printf("The position of \"Burger\" is: ")
	// fmt.Printf("%d\n", strings.Index(str, "Burger"))

	// var str string = "Hello, how is it going, Hugo?"
	// var manyG = "gggggggggg"

	// fmt.Printf("Number of H's in %s is: ", str)
	// fmt.Printf("%d\n", strings.Count(str, "H"))

	// fmt.Printf("Number of double g's in %s is: ", manyG)
	// fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	// var origS string = "Hi there! "
	// var newS string

	// newS = strings.Repeat(origS, 3)
	// fmt.Printf("The new repeated string is: %s\n", newS)

	// strings

	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

}

func t3() {
	// s := [3]int{1, 2, 3}[:]
	// var x = []int{2, 3, 5, 7, 11}

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])

	fmt.Println("end")
}

func TestMain(t *testing.T) {
	// a()
	// b()
	// t1()
	// t2()
	t3()
}
