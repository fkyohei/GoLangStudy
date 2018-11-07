package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    // 構造体のフィールドは、構造体のポインタを通してアクセスすることも可能
    // C言語とかだと(*p).Xのような表記（だったはず）だが、Goではp.Xと書ける
    p := &v
    p.X = 1e9
    fmt.Println(v)
}
