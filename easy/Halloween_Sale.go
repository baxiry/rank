package main
import "fmt"
func prices(price, diff, min, stoc int32 ) int32 {
    var result int32
    for stoc >= price {
        result++
        stoc  = stoc - price
        price -= diff
        if price < min {
            price = min
        }
    }
    return result
}

func main() {
    r := prices(20, 3, 6, 85)
    fmt.Println("+++++++++++", r)
}
