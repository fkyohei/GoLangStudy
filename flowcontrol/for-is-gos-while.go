package main

import "fmt"

func main() {
    sum := 1
    // goのwhile文はforを使用する
    // セミコロンを省略することでwhile文と同じものを表すことができる
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
