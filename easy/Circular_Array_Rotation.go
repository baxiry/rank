package main
import "fmt"


func circular(arr, queries []int32, k int32) []int32 {

    ln := len(arr)
    mv := int(k) % ln

    res := make([]int32, ln)

    for i:= 0; i < ln; i++ {

        if mv+i < ln {
            res[mv+i] = arr[i]
            fmt.Println(arr[i])
        } else {
            res[(i+mv)-ln] = arr[i]
            fmt.Println(arr[i])
        }

    }
    ln = len(queries)
    result := make([]int32, ln)

    for i:= 0; i < ln; i++ {
        result[i] = res[queries[i]]

    }

    fmt.Println("res is : ",  res)

    return result
}

func main() {
    arr := []int32{1, 2, 3}
    q := []int32{0, 1, 2}

    fmt.Println(arr)
    res := circular(arr, q, 2)
    fmt.Println(res)

}
