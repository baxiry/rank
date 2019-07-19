package main
import "fmt"

func divisibleSumPairs(n, k int32, ar []int32) int32 {

    var result int32
    for k1, v1 := range ar {
        for i:= k1+1; i < int(n); i++ {
            if (v1+ar[i]) % k == 0 {
                result++
            }
        }
    }
    return result


}

func main() {
    ar := []int32{1, 3, 2, 6, 1, 2}
    result := divisibleSumPairs(6, 3, ar)
    fmt.Println(result)

}
