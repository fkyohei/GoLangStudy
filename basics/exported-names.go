package main

import (
    "fmt"
    "math"
)

func main() {
    // path.piではエラーになる。
    // fmt.Println(math.pi);
    // パッケージ内の外部から参照できる変数は大文字始まりで決まっている。
    fmt.Println(math.Pi);
}
