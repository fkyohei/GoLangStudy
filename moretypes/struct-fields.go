package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    // 構造体のフィールドにはドットを用いてアクセスする
    v.X = 4
    fmt.Println(v.X)
}
