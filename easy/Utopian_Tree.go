package main
import "fmt"

func utopianTree(n int32) int32 {

    var i int32
    var h int32 = 0

    for i = 0; i <= n; i++ {
        if i % 2 == 0 {
            h++
        } else {
            h *= 2
        }
    }
    return h

}

func main() {
    h := utopianTree(5)
    fmt.Println(h)

}
