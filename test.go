package main

import (
	"fmt"
	"unsafe"
)

func ToUpper(char byte) byte {
	if char < 'f' && char > 'a' {
		return char - 32
	}
	return char
}

type Stct struct {
	b bool
	a int64
}
type Stct2 struct {
	c int32
	b bool
	a int32
}

func main() {
	fmt.Println(string(ToUpper('g')))
	fmt.Println("Size of Struct{}{}", unsafe.Sizeof(struct{}{}))
	fmt.Println("Size of Struct{bool;int64}", unsafe.Sizeof(Stct{}))
	fmt.Println("Size of Struct{int64;bool}", unsafe.Sizeof(Stct2{}))
	fmt.Println("Size of string:", unsafe.Sizeof("jkdsnksjbgsjbg sjkbg jshkdbg g"))
	fmt.Println("Size of Slice:", unsafe.Sizeof(make([]int, 500)))
	fmt.Println("Size of Map:", unsafe.Sizeof(make(map[int]int)))
	fmt.Println("Size of Func:", unsafe.Sizeof(test))

}
func test(x *int) {

}
