package main

//noinspection ALL
import (
	"fmt"
	"utools"
)

func main() {
	/* 整数、常量、赋值操作 */
	var d int = 1
	var e, f, g int = 2, 3, 4
	var h = 5
	i, j := 6, 7

	const LENGTH int = 10
	const WIDTH int = 5
	const a, b, c = 1, false, "str"

	area := LENGTH * WIDTH
	fmt.Printf("面积为:%d\n", area)
	println("fuxing")
	println(a, b, c, d, e, f, g, h, i, j)
	fmt.Println("Hello, World!")
	fmt.Println(utools.Max(1, 2))
	fmt.Printf("average is %v\n", utools.Average(10, 99))

	//var str,rawstr = "abc123", `
	//	hello
	//	 	world
	//	 	\n
	//`
	//var (
	//	a int
	//	b int32
	//	c []byte
	//)
	//c = []byte(str)
	//c[0] = 't'
	//fmt.Printf("string c is %v\n", string(c))
	//fmt.Printf("raw string is %v\n", rawstr[2:len(rawstr)-1])
	//fmt.Printf("a:%v b:%v", a, b)
}
