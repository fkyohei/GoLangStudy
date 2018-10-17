package main

import "fmt"

// 定数はconstを使用して定義
// constは文字・文字列・boolean・数値に使用できる
const Pi = 3.14

func main() {
    // constのときは「:=」は使用できない
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
