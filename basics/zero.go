package main

import "fmt"

func main() {
    // 変数に初期値を与えずに宣言すると、ゼロ値が与えられる
    // ゼロ値は形によって以下のようになる
    /**
     * 数値型（int, float, etc）：0
     * bool型：false
     * string型：""
     */
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q", i, f, b, s)
}
