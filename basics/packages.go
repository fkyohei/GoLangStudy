package main

import (
    "fmt"   // Go言語の標準コードフォーマットに変換してくれる。基本的にはimportすべき?
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is ", rand.Intn(10))
}
