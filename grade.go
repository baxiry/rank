package main
import "fmt"


func Grades(grades []int32) []int32 {

    for k, v := range grades {
        if v > 37 {
            if (v + 1)%5 == 0 {
                grades[k] += 1
            }
            if (v + 2)%5 == 0 {
                grades[k] += 2
            }
        }
    }
    return grades
}


func main() {
    slc := []int32{38, 34, 5, 45, 56, 67, 78}
    fmt.Println(slc)
    result := Grades(slc)
    fmt.Println(result)
}
