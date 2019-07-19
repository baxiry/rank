package main
import "fmt"
import "strconv"

func cavityMap(grid []string) {// []string {
    le := len(grid[0])-1
    for _, str := range grid {
        for i:= 0; i <= le; i++ {

            if i > 0 && i < le {
                s0, _ := strconv.Atoi(string(str[i-1]))
                s1, _ := strconv.Atoi(string(str[i]))
                s2, _ := strconv.Atoi(string(str[i+1]))
                if s1 > s0 && s1 > s2 {
                    n = []
                }
            }
            //fmt.Println(k, string(val2))
        }
        fmt.Println()
    }

}

func main() {
    grid := []string{"1112", "1912", "1892", "1234"}
    fmt.Println(grid)
    cavityMap(grid)

}
