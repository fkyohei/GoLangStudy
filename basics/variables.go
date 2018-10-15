package main

import "fmt"

// 変数はvarで宣言する
var c, python, java bool

func main() {
    // 関数内でも同じ
    var i int
    fmt.Println(i, c, python, java)
}
