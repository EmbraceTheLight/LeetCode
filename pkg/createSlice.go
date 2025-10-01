package pkg

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strings"
)

func CreateSlice[T constraints.Ordered]() []T {
	fmt.Println("Input row length:")
	var length int
	fmt.Scan(&length)
	var in T

	var ret = make([]T, 0)
	for i := 0; i < length; i++ {
		fmt.Scan(&in)
		ret = append(ret, in)
	}
	//reader := bufio.NewReader(os.Stdin)
	//for
	return ret
}

func CreateIntSlice2[T constraints.Ordered]() [][]T {
	fmt.Println("Input rows:")
	var rows int
	fmt.Scan(&rows)
	var ret = make([][]T, 0)
	for i := 0; i < rows; i++ {
		var row = CreateSlice[T]()
		ret = append(ret, row)
	}
	return ret
}

func CreateByteSlice2() [][]byte {
	fmt.Println("Input rows:")
	var rows int
	fmt.Scan(&rows)
	fmt.Println("Input bytes per row(space separated):")
	var ret = make([][]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < rows; i++ {
		var str string
		str, _ = reader.ReadString('\n')
		tmp := strings.Split(str, " ")
		bytes := make([]byte, 0)
		for _, bt := range tmp {
			bytes = append(bytes, []byte(bt)[0])
		}

		ret = append(ret, bytes)
	}
	return ret
}
