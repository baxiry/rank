package main
import "fmt"


func bonAppetit(bill []int32, k int32, b int32) {
    var sum int32
    sum = 0
    for _,v := range(bill) {
        sum += v
    }
    rslt := (sum - bill[k])/2
    if b == rslt {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(b-rslt)
    }
}

func main() {
    arr := []int32{3, 10, 2, 9}
    bonAppetit(arr, 1, 12)
}
