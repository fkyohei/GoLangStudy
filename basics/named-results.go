package main

import "fmt"

// 関数の戻り値の型定義部分に名前をつけることが可能。
// 関数の最初で定義した変数として扱われる。
// 名前をつけることで、return文にはなにも書かなくて良くなる（naked return と呼ばれる）。
// ただし、長い関数で使用すると読みづらくなると公式でも注意がある。
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
