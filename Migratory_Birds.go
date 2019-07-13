package main
import "fmt"

func typebird(arr []int32) int32 {
    types := make(map[int32]int32)

    for _, v := range arr {
        types[v]++
    }

    var larg int32
    for _, v := range types {
        if larg < v{
            larg = v
        }
    }
    fmt.Println(larg)
    fmt.Println("types is ", types)

    var slc []int32

    for k, v := range types {

        if v == larg {
            slc = append(slc, k)
        }
    }

    result := slc[0]
    for _,v := range slc {
        if result > v {
            result = v
        }
    }
    return result
}
func main() {
    arr := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
    result := typebird(arr)
    fmt.Println(arr)
    fmt.Println(result)
}
