package main

import (
    "fmt"
)

func insertionSort(arr []int32) int32 {
    ln := len(arr)
    //ln1 := len(arr)

    mv := 0
    for i := 0; i < ln; i++ {
        for  j:=i+1; j < ln; j++{
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
                i++
                 mv++
                 fmt.Println(mv)
            }
            fmt.Println(arr)
        }

    }
    return int32(mv)
}

func main() {
    arr := []int32{2, 4, 3, 1, 1, 1, 1, 2}
    res := insertionSort(arr)
    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Println()
    fmt.Println(arr)
    fmt.Println(res)
}
