package main

import "fmt"

func main() {
    // deferへ渡した関数が複数ある場合は、その呼出はスタックされる
    // 呼び出し元の関数がreturnするとき、deferへ渡した関数はLIFOの順に実行される
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
