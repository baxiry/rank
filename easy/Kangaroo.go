package main
import "fmt"

func kangaroo(x1, v1, x2, v2 int32) string {
    fmt.Println(x1, v1, x2, v2)

    for i:=0; i <= 10000; i++ {
        if x1 == x2 {
            return "YES"
        }
        x1 += v1
        x2 += v2
    }
    return "NO"

}

func main() {
    r := kangaroo(0, 2, 5, 3)
    fmt.Println(r)

}
