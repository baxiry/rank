package main
import "fmt"

func pass(str string) int32 {
    need := int32(4)
    mp := make(map[string]bool)
    for _,v := range str {
        // number : 48 57
        if v > 47 && v < 58 {
            mp["n"]=true
        }
        // lower  : 97 122
        if v > 96 && v < 123 {
            mp["l"]=true
        }
        // upper  : 65 90
        if v > 64 && v < 91 {
            mp["u"]=true
        }
        // speacel; 33 43
        if v > 32 && v < 46 {
            mp["s"]=true
        }
    }
    for _, v := range mp {
        if v == true { need-- }
    }
    ln := int32(len(str))
    if ln < 6 {
        ln = 6-ln
    } else {
        ln = 0
    }
    fmt.Println("need: ", need)
    fmt.Println("ln: ", ln)
    if ln > need {
        return ln
    }
    return need
}

func main() {
    s := "AUzs-nV"

    fmt.Println(s)
    res := pass(s)
    fmt.Println(res)
    fmt.Println(s[4])
}
