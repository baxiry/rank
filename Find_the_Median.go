package main
import "fmt"

// 0 1 2 4 6 5 3
func median(ar []int32) int32 {
    li := len(ar)
    lj := li-1
    for i := 1; i < li; i++ {
        for j := lj; j > i; j-- {
            if ar[j] < ar[j-1] {
                ar[j-1], ar[j] = ar[j], ar[j-1]            }
                fmt.Println(ar)
        }
    }
    median := len(ar)/2


    return ar[median]
}

func main() {
    ar := []int32{0,9,8,7,4,3,2,1,5}
    m := median(ar)
    fmt.Println("median Is :", m)

}
