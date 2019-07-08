package main
import "fmt"

func pageCount(num int32, page int32) int32 {
    if num % 2 != 0 {
        num--
  }

    front := page/2
    back := (num-page+1)/2
    if back < front {
        return back
    }
    fmt.Println("back is:", back, "front is :", front)
    return front
}

func main() {

    result:=pageCount(6, 5)
    fmt.Println(result)

}

