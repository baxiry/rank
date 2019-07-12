package main
import "fmt"


func designerPdfViewer(h []int32, word string) int32 {

    // find length of word
    l := len(word)

    var max int32 = 0
    mp := make(map[string]int32)
    latter := "abcdefghijklmnopqrstuvwxyz"

    // get hight of all character
    for k, v := range latter {
        mp[string(v)] = h[k]
        //fmt.Println(string(v))
    }

    // find max hight char
    for _, v := range word {
        fmt.Println(mp[string(v)])
        if max < mp[string(v)] {
            max = mp[string(v)]
        }
    }
    return max*int32(l)

}



func main() {
    slc := []int32{1,2,3,4,5,6,7,8,9,10,11, 12,13,14,15,16,17,18,19,20,21,22,23,24,25,26}
    h := designerPdfViewer(slc, "adam")
    fmt.Println(h)

}
