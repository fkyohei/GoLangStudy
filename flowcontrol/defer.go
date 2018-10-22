package main

import "fmt"

func main() {
    // deferステートメントはdeferへ渡した関数の実行を呼び出し元の関数の終わりまで遅延させるもの
    // deferへ渡した関数の引数はすぐに評価されるが、その関数自体の呼び出しは呼び出し元の関数がreturnするとき
    defer fmt.Println("world")

    fmt.Println("hello")
}
