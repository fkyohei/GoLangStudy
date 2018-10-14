package main

import "fmt"

// 関数は複数の戻り値を返すことが可能
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    // 受け取り側もカンマ区切りで変数を準備してあげるだけ。
    // 「:=」を使用すると変数宣言と代入を同時にできる。
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
