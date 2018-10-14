package main

import (
    "fmt"
    "math"
)
// 上記は下記と同義だが、上記の書き方が推奨されている
/*
import "fmt"
import "math"
*/

func main() {
    fmt.Printf("How you have %g problems.", math.Sqrt(7))
}
