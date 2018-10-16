package main

import (
    "fmt"
    "math"
)

func main() {
    // 変数v, 型Tがあった場合、T(v)は変数vをT型に変換する
    // var i int = 42
    // var f float64 = float64(i)
    // var u uint = uint(f)
    // 下記のようにシンプルに書くことも可能
    // i := 42
    // f := float64(i)
    // u ;= uint(f)
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}
