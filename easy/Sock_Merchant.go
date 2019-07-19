package main
import "fmt"

func sockMerchant(n int, ar []int32) map[int32]int {
    mp := make(map[int32]int)
    total :=0
    for _, v := range(ar) {
        mp[v]++
    }
    for _, v := range(mp) {
         if v % 2 != 0 {
             v--
         }

         total += v
    }

    total /= 2
    fmt.Println("total is :", total)
    return mp
}

func main() {
    arr := []int32{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
    fmt.Println(arr)
    res := sockMerchant(len(arr),arr)
    fmt.Println(res)
}
