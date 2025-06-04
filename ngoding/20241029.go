package main
import "fmt"

func main() {
    var pelat string
    
    fmt.Scanln(&pelat)
    
    isTrue := len(pelat) > 0 && pelat[0] == 'D'
    
    fmt.Println(isTrue)
}