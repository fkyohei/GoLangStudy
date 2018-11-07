package main

import "fmt"

type Vertex struct {
    X, Y int
}

// structリテラル
// フィールドの値を列挙することで、新しい構造体の初期値を割り当てる
// Name: 構文を使って、フィールドの一部だけを列挙することもできる。（この場合、指定順序は関係ない）
var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1}  // この場合Yは0
    v3 = Vertex{}      // この場合XとYは0
    p  = &Vertex{1, 2} // 新しく割り当てられた構造体のポインタ
)

func main() {
    fmt.Println(v1, v2, v3, p)
}
