package main
import "fmt"


func catAndMouse(x, y, z int32) string {

    var catx int32
    var caty int32

    if z > x {
        catx = z - x
    } else {
        catx = x- z
    }

    if z > y {
        caty = z-y
    } else {
        caty = y - z
    }

    if catx < caty {
        return "Cat A"
    } else if caty < catx {
        return "Cat B"
    } else {
        return "Mouse C"
    }
}


func main() {

    u := catAndMouse(12, 3, 2)
    fmt.Println(u)
}
