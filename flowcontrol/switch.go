package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    // GoのSwitch文はcaseの最後に自動でbreakが挿入されるため書く必要はなく、選択されたcase文しか実行されない
    // また、caseは定数である必要はなく、関係する値は整数である必要がない
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}
