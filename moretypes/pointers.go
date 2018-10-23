package main

import "fmt"

// Goはポインタを扱う
// 変数Tのポインタは*Tで、ゼロ値はnil
// var p *int
// &オペレータはそのオペランドへのポインタを引き出す
// i := 42
// p = &i
// *オペレータはポインタの指す先の変数を示す
// fmt.Println(*p)  // ポインタpを通してiから値を読み出す
// *p = 21          // ポインタpを通してiへ値を代入する
// C言語と異なりポインタ演算はない

func main() {
    i, j := 42, 2701

    p := &i     // iのポイントをpに
    fmt.Println(*p)     // ポインタpを通してiの値を読み出す
    *p = 21             // ポインタpを通してiへ値を代入する
    fmt.Println(i)

    p = &j
    *p = *p / 37
    fmt.Println(j)
}
