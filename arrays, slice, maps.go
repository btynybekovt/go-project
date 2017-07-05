//#1
package main

import "fmt"

func main(){
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for i := 0; i < 5; i++{
		total += x[i]
	}
	fmt.Println(total / 5)
}
//#2
func main(){
slice1: := []int{1,2,3}
slice2: := append(slice1, 4, 5)
fmt.Println(slice1, slice2)
}
//#3
func main(){
	slice1: := [] int{1,2,3}
	slice2: := [] make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2) 
}
//#4
func main(){
	elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements["Li"])
}
}
//Home Work
func main(){
	x := []int{
	48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
	}
	y := x[0]
	for i := 1; i <len(x); i++{
		if y > x[i]{
			y = x[i]
		}
		fmt.Println(y)
	}
}