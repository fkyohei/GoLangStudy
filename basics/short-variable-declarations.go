package main

import "fmt"

// 関数の外での暗黙的な型宣言はできない（エラーになる）
// l := "this is error."

func main() {
    // 通常の変数定義と初期化
    var i, j int = 1, 2
    // 暗黙的な型宣言（「:=」を使用すると「var」も型もいらなくなる）
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}
