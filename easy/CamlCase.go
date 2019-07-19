package main
import "fmt"


func camelcase(s string) int32 {
    var sum int32
    for _, v := range s {
        if v > 64 && v < 91 {
            sum++
        }
    }
    return sum +1

}
func main() {
    r := camelcase("aaaaAaaaaaAaaaaaAaaaa")
    fmt.Println(r)

}
