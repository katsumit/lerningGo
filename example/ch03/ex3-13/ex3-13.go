package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	// ここは挙動が違う
	// 期待値[2 3 4 4] 4
	// 実行値[2 3] 2
	fmt.Println(x, num)
}
