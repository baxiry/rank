package main
import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
    p := int(k)

    var res int32 = 100
    for i:= 0; i<len(c); i=i+p {
        if c[i] > 0 {
            res = res-1-2
        } else {
            res = res -1
        }
    }
    return res
}

func main() {
    k := []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}
    res := jumpingOnClouds(k, 3)
    fmt.Println(res)
}
