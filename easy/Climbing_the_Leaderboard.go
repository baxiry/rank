package main
import "fmt"

// 100 90 90 80 75 60
// 50 65 77 90 102

func climbingLeaderboard(scores, alice []int32)  []int32 {

    rank := 0
    var uniq int32
    for i := 0; i<len(scores); i++ {
        uniq = 0
        rank = 0
        for j := 0; j < len(alice); j++ {
            rank++
            if alice[j] > scores[i] && scores[i] != uniq {
                fmt.Print(rank, " ")
                uniq = scores[i]
            }
        }
        //    fmt.Println()
    }
    var result []int32
    return result
}

func main() {
    arr1 := []int32{100, 100, 50, 40, 40, 20, 10}
    arr2 := []int32{5, 25, 50, 120}
    slc := climbingLeaderboard(arr1, arr2)
    fmt.Println(slc)
    //fmt.Println(arr1)
    //fmt.Println(arr2)
}
