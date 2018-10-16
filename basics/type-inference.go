package main

import "fmt"

func main() {
    // 型を指定せずに変数定義すると、右側の値から型推論が行われる
    v := 3  // Change me
    fmt.Printf("v is of Type %T\n", v)
}
