package main
import "fmt"

func angryProfessor(k int32, arr []int32) string {

    var ok int32 = 0
	for _, v := range arr {
        if v <= 0 {
            ok++
        }
    }
    if ok >= k {
        return "NO"
    }
    return "YES"


}

func main() {
    arr := []int32{0, -1, 2, 1}

    res := angryProfessor(2, arr)
    fmt.Println(res)

}
