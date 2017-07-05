//#1
package main

import "fmt"

func one(xPtr *int) {
    *xPtr = 1
}
func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr) 
}
//Home Work #1

func swap(x *int, y *int) {
*x, *y = *y, *x
}
func main() {
x, y := 2, 3
fmt.Println("x =", x, ", y =", y)
swap(&x, &y)
fmt.Println("x =", x, ", y =", y)
}