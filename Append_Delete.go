package main
import "fmt"

func ad(s string, t string, k int32) string {
    ls := len(s)
    lt := len(t)

    if s == t {
        if ls+lt < int(k) {
            fmt.Println(1)
            return "Yes"
        }
    }

    lm := 0
    for i:= 0; i < ls && i < lt;i++ {
        if s[i] != t[i] {
            lm = i
            fmt.Println("lm is :",lm)
            break
        }
        lm++
    }
    nomatch := (ls-lm) + (lt-lm)
    fmt.Println("no match is :", nomatch)
    if nomatch > int(k) {
        fmt.Println(2)
        return "No"
    }
    if ls < lt && ls+lt > int(k) {
        fmt.Println("NNNN")
        return "No"
    }
    fmt.Println(3)
    return "Yes"

}

func main() {
    s := "zzzzz"
    t := "zzzzzzz"

    str := ad(s, t, 9)
    fmt.Println(str)
}
