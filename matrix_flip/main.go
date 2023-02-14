package main

import "fmt"

func matrixFilp(foo *[][]int) [][]int {
	_foo := *foo
	// 矩阵大小 = 单行长度
	l := len(_foo)
	ll := 0
	res := make([][]int, l)
	for ll <= l-1 {
		_res := make([]int, l)
		for i := l - 1; i >= 0; i-- {
			child := _foo[i]
			_res[l-1-i] = child[ll]
			// _res[len(_foo)-1-i] = child[ll]
			continue
		}
		res[ll] = _res
		ll = ll + 1
	}
	return res
}

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	fmt.Println(matrixFilp(&matrix))
}
