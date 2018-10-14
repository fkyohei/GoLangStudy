package main

import "fmt"

// 引数の型がすべて同じ場合、最後の型だけ残してあとは省略が可能
func add(x, y int) int {
    return x + y
}
// 下記と同義
/*
func add(x int, y int) int {
    return x + y
}
*/

func main() {
    fmt.Println(add(42, 13))
}
