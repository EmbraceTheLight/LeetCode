package pkg

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
)

func CreateSlice[T constraints.Ordered]() []T {
	var ret []T
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter slice (e.g. [1, 2, 3]): ")
	if !scanner.Scan() {
		return ret
	}
	input := scanner.Text()
	tmp := strings.Trim(input, "[]")
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
