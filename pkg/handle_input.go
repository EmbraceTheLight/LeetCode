package pkg

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

func stringToBytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
func makeElementSlice[T constraints.Ordered](str string) []T {
	var ret []T
	tmp := strings.Trim(str, "[]")
	if len(tmp) == 0 {
		return nil
	}
	parts := strings.Split(tmp, ",")

	for _, p := range parts {
		p = strings.TrimSpace(p)
		var val any
		var err error

		switch any(*new(T)).(type) {
		case int:
			var v int
			v, err = strconv.Atoi(p)
			val = v
		case float64:
			var v float64
			v, err = strconv.ParseFloat(p, 64)
			val = v
		case string:
			p = strings.Trim(p, `"`)
			val = p
		case byte:
			p = strings.Trim(p, `'`)
			b := stringToBytes(p)
			val = b[0]
		default:
			fmt.Printf("unsupported type\n")
			return ret
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing %q: %v\n", p, err)
			continue
		}

		ret = append(ret, val.(T))
	}
	return ret
}

func CreateSlice[T constraints.Ordered]() []T {
	var ret []T
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter slice (e.g. [1, 2, 3]): ")
	if !scanner.Scan() {
		return ret
	}
	input := strings.TrimSpace(scanner.Text())
	return makeElementSlice[T](input)
}

func CreateSlice2D[T constraints.Ordered]() [][]T {
	fmt.Println(`Input 2D Array Slice. (e.g. [[1,2,3],[4,5,6]])`)
	scanner := bufio.NewScanner(os.Stdin)
	input := ""
	var ret = make([][]T, 0)
	if scanner.Scan() {
		input = scanner.Text()
	}
	input = strings.TrimSpace(input)
	input = input[1 : len(input)-1] // 去掉首尾的方括号

	var leftBound, rightBound int
	leftBound = strings.Index(input, "[")
	for leftBound != -1 {
		rightBound = strings.Index(input, "]")
		if rightBound == -1 {
			panic("input is not a valid 2D array slice")

		}
		row := makeElementSlice[T](input[leftBound : rightBound+1])
		ret = append(ret, row)

		input = input[rightBound+1:]
		leftBound = strings.Index(input, "[")
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

func CreateString() string {
	var str string
	scan := bufio.NewScanner(os.Stdin)
	if !scan.Scan() {
		panic("input error")
	}
	str = scan.Text()
	return str
}
