package main
import "fmt"

func minimum(slc []int32) int32 {
    min := 2147483647
    for k1, v1 := range(slc) {
        for k2, v2 := range(slc[k1+1:]) {
            if v1 == v2 && (k1+k2+1)-k1 < min {
                min = (k1+k2+1)-k1
                break
            }
        }
    }
    if min == 2147483647 {
        min = -1
    }
    return int32(min)
}

func main() {
    slc1 := []int32{1, 2, 11, 4, 11, 5, 6, 3 , 8, 10, 7, 5, 9, 3, 1, 10, 8, 2, 4, 6}
    fmt.Println(slc1)
    //slc1 := []int32{1, 3, 2, 4, 5, 6, 3 , 8, 5, 6, 1}
    fmt.Println(minimum(slc1))
    //fmt.Println(3-5,"\n", 5-3)
}
