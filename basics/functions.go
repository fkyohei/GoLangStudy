package main

import "fmt"

// 関数は0個以上の引数を取ることが可能
// 型名は変数名の「後ろ」に書く
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
