package main
import (
    "fmt"
    "math/rand"
)

func quicksort(a []int32) []int32 {
    if len(a) < 2 {
        return a
    }
    left, right := 0, len(a)-1
    pivot := rand.Int() % len(a)
    a[pivot], a[right] = a[right], a[pivot]
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
    a[left], a[right] = a[right], a[left]
    quicksort(a[:left])
    quicksort(a[left+1:])
    return a
}

func pair(k int32, arr []int32) int32 {
    ar := quicksort(arr)
    i := 0
    j :=1
    sum := 0
    ln := len(arr)

    for j < ln {
        diff := ar[j] - ar[i]

        if diff == k {
            sum++
            j++
        } else if diff > k {
            i++
        } else if diff < k {
            j++
        }
    }

    return int32(sum)
}

func main() {
    arr := []int32{1, 5, 3, 4, 2}
    res := pair(2, arr)
    fmt.Println(res)

}

