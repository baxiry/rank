package main
import "fmt"

func pickingNumbers(arr []int32) int32 {
    max := 1
    cond := 0

    for k1, v1 := range arr {
        cond = 0
        for k2, v2 := range arr {
            if k1 != k2 && (v1 == (v2-1)) || (v1 == v2) || (v1 == (v2-1)) {
                cond++
            }
        }
        if max < cond {
            max = cond
        }
    }
    return int32(max)
}


func main() {
    n := []int32{4, 6, 5, 3, 3, 1}
    res := pickingNumbers(n)
    fmt.Println(res)
}
