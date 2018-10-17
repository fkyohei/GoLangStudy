package main

import "fmt"

// ※ ちょっと説明が理解できない・・・。
// 数値の定数は高精度な値
// 型のない定数はその状況によって必要な型をとる
// https://go-tour-jp.appspot.com/basics/16

const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int {
    return x*10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    // fmt.Println(needInt(Big))  これはOverflowでエラーとなる
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
