package main
import "fmt"


func main() {
    for i:= 0; i < 101; i++ {
        if (0 == i % 5) || (0 == i % 3){
            fmt.Println(i, "is fuz and buz")
        } else if i % 5 == 0 {
            fmt.Println(i, "is fuzz")
        } else {
            fmt.Println(i, "is Buzz")
        }
    }
}
