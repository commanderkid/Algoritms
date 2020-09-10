package main

import "fmt"

func main() {
    fmt.Println(GCD(6, 8))
}

func GCD(a uint32, b uint32) uint32 {

    var devider uint32 = 1
    var answer uint32 = 1 
    var biggest, smallest uint32
    if a < b {
        biggest = b
        smallest = a
    } else if a > b {
        biggest = a
        smallest = b
    } else {
        return a
    }

    for devider <= smallest {
        if smallest % devider == 0 && biggest % devider == 0 {
            answer = devider
        }
        devider += 1
    }
    return answer 

}
