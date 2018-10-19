package main

import "fmt"

func main()  {
    sum := 0
    // for文の書き方は基本C等と同じ
    // ただし、丸括弧は不要。中括弧は必要。
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
