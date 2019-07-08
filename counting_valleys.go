package main
import "fmt"

func countingValleys(n int32, s string) int32 {
    n=0// we not need this becouse we note use arry.
    var staps int
    var valleys int32
    for _, v := range s {
        if string(v) == "U" {
            staps++
            if staps == 0 {
                valleys++
            }
        } else {
            staps--
        }
    }
    return valleys
}


func main() {
    v := countingValleys(8, "DUDUDUDUDUDUDUDUDUDUDUDUDUDU")
    fmt.Println(v)


}
