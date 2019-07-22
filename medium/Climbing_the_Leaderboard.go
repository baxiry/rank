package main
import "fmt"

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
    ls := len(scores)
    la:= len(alice)
    t := 0
    list := make([]int32, la)
    //scores = append(scores, int32(0))
    for i := 1; i < la; i++ {
        for j:= 1; j < ls; j++ {
            // if alice[i] > scores[j] {break}
            if alice[i-1] > scores[j] {
                    t++
                if scores[j-1] > scores[j] {
                }
            }
        fmt.Println(t)
        }
        t = 0
    }
    return list
}
func main() {
    sc := []int32{100,90,90,80,75,60}
    al := []int32{50,65, 77, 90, 102}
    res := climbingLeaderboard(sc, al)
    fmt.Println("result is :", res)

}
