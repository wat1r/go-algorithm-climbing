package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello go")
	// fmt.Printf("%s", "user")

	fmt.Println("hello go")
	// test1()
	if_test()
	// i := 10
	// increment(&i)
	// fmt.Println(i)
	// fmt.Println(add(1, 2))
	// sum, msg := add(10, 20)
	// fmt.Println(sum)
	// fmt.Println(msg)

	// alice := person{name: "Alice", age: 12, gender: "Female"}
	// fmt.Println(alice)
	// bob := person{"bob", 10, "Male"}
	// fmt.Println(bob)
	// fmt.Println(bob.age)
	// // bobnpt = &person{name: "bob", age: 19, gender: "Male"}
	// // fmt.Println(bobnpt.name)
	// bob.describe()

	fmt.Println("hah")
	// pp := &person{name: "Bob", age: 42, gender: "Male"}
	// pp.describe()
	// // => Bob is 42 years old
	// pp.setAge(45)
	// fmt.Println(pp.age)
	// //=> 45
	// pp.setName("Hari")
	// fmt.Println(pp.name)
	//=> Bob

	// var a animal
	// a = snake{Poisonous: true}
	// fmt.Println(a.description())
	// a = cat{Sound: "Meow!!"}
	// fmt.Println(a.description())
	// mapA := map[string]int{"apple": 5, "lettuce": 7}
	// mapB, _ := json.Marshal(mapA)
	// fmt.Println(string(mapB))
	// str := `{"page":1,"fruits":["apple","peach"]}`
	// res := response{}
	// json.Unmarshal([]byte(str), &res)
	// fmt.Println(res.PageNumber)
	// fmt.Println(res)
	// resp, err := http.Get("http://baidu.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(resp)
	// num := 5
	// num := -1
	// if inc, err := Increment(num); err != nil {
	// 	fmt.Printf("Failed Number : %v ,error message :%v", num, err)
	// } else {
	// 	fmt.Printf("Increment Number :%v", inc)
	// }

	// f()
	// fmt.Println("Returned normally from f.")
	// go c()
	// fmt.Println("I am main")
	// time.Sleep(time.Second * 2)

	// c := make(chan string)
	// go func() {
	// 	c <- "hello"
	// }()
	// msg := <-c
	// fmt.Println(msg)

	// c1 := make(chan string)
	// c2 := make(chan string)
	// go speed1(c1)
	// go speed2(c2)
	// fmt.Println("The first to arrive is:")
	// select {
	// case s1 := <-c1:
	// 	fmt.Println(s1)
	// case s2 := <-c2:
	// 	fmt.Println(s2)
	// }

	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
}

func test1() {
	// var a int
	// var a = 1
	message := "hello world\n"
	fmt.Print(message)
	// var b, c int = 2, 3
	// fmt.Print(b, c)
	// var a bool = true
	// var b int = 1
	// var c string = "hello"
	// var d float32 = 1.222
	// var x complex128 = cmplx.Sqrt(-5 + 12i)
	// fmt.Print(x)
	// var a [5]int
	// var multiD [2][3]int
	// var b[] int
	// numbers := make([]int, 5, 10)
	// fmt.Print(numbers)
	// numbers = append(numbers, 1, 2, 3, 4)
	// number2 := make([]int, 15)
	// copy(number2, numbers)
	// number2 := []int{1, 2, 3, 4}
	// fmt.Println(number2)
	// slice1 := number2[2:]
	// fmt.Println(slice1)
	// slice2 := number2[3:]
	// fmt.Println(slice2)
	// slice3 := number2[1:4]
	// fmt.Println(slice3)
	// m := make(map[string]int)
	// m["clearity"] = 2
	// m["simplicity"] = 3
	// a := 1.1
	// b := int(a)
	// fmt.Println(b)

	fmt.Println("end")
}

func if_test() {
	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negtive")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("none")
	// }
	// fmt.Println("end")
	// i := 0
	// sum := 0
	// for i < 10 {
	// 	sum += 1
	// 	i++
	// }
	// fmt.Println(sum)
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)
	// for {

	// }

	// var ap *int
	// a := 12
	// ap = &a
	// fmt.Println(ap)
	// fmt.Println(*ap)

}

func increment(i *int) {
	*i++
}

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a int, b int) (c int) {
// 	c = a + b
// 	return
// }

func add(a int, b int) (int, string) {
	c := a + b
	return c, "success"
}

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) describe() {
	fmt.Printf("%v is %v years old", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p person) setName(name string) {
	p.name = name
}

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

type response struct {
	PageNumber int      `json:"page"`
	Fruits     []string `json:"fruits"`
}

func Increment(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("math: cannot process negative number")
	}
	return (n + 1), nil
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Printing - in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func c() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am concurrent")
}

func speed1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "speed 2"
}
