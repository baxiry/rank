package main
import (
    "fmt"
)

func fibonacciModified(t1 int32, t2 int32, n int32) uint64 {
    a := uint64(t1)
    b := uint64(t2)
    c := uint64(0)

    for i:=0; i < int(n); i++ {
        fmt.Print(a,",")
        c = a+(b*b)
        a = b
        b = c
    }
    println()
    return c

}

func main() {
    res := fibonacciModified(0, 1, 10)
    fmt.Println(res)
}
