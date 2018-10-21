package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    // switch文に条件を書かない場合、「switch true」と同義となる
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
