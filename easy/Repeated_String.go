package main
import "fmt"

func repeatedString(str string, n int64) int64 {
    ln := len(str)
    num := 0
    for i := 0; i < ln; i++ {
        if string(str[i]) == "a" {
            num++
        }
    }
    countStr := int(n)/ln
    num = countStr*num
    a := int(n)%ln
    for i:= 0; i< a; i++ {
        if string(str[i]) == "a" {
            num++
        }
    }
    return int64(num)
}


func main() {
    str := "adaad"
    n := repeatedString(str, 100)
    fmt.Println(n)
}



