package main

import "fmt"

// varで宣言時に初期化することもできる（型の宣言の後ろに「=」と値を書く）
var i, j int = 1, 2

func main() {
    // 初期化する場合は型を省略しても良い
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
