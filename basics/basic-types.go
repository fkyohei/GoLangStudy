package main

import (
    "fmt"
    "math/cmplx"
)

/**
 * Goの基本の型は以下
 * bool
 * string
 * int int8 int16 int32 int64
 * uint uint8 uint16 uint32 uint64 uintptr
 * byte  // uint8 の別名
 * rune  // int32の別名。Unicode のコードポイントを表す
 * float32 float64
 * complex64 complex128
 *
 * int, uint, uintptr型は、32bitのシステムでは32bit、64bitのシステムでは64bit。
 * 特別な理由がない限り、整数の変数は「int」の使用が推奨
 */

 // 変数定義もインポートステートメントと同様まとめて宣言が可能
 var (
    ToBe    bool        = false
    MaxInt  uint64      = 1<<64 - 1
    z       complex128  = cmplx.Sqrt(-5 + 12i)
 )

 func main() {
     fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
     fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
     fmt.Printf("Type: %T Value: %v\n", z, z)
 }
