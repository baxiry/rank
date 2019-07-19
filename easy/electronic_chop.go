package main
import "fmt"

func getMoneySpent(keyboards, drives []int32, b int32) int32 {
    big := int32(-1)
    result := big
    for _, kb := range keyboards {
        for _, usb := range drives {
            if kb + usb <= b {
                big = kb+usb
            }
            if result < big {
                result = big
            }
        }
    }
    return result
}

func main() {
    k := []int32{4, 3, 7, 1}
    u := []int32{4, 6, 5}
    res := getMoneySpent(k, u, 12)
    fmt.Println(res)

}
