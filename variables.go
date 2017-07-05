//#1
package main

import "fmt"

func main(){
	var x string = "Hello World"
	fmt.Println(x)
}
//#2

func main(){
	var x string
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}

//#3
func main(){
	x := "Max"
	fmt.Println("My dog's name "x)
}
//Home Work #1
func main(){
	fmt.Println("Введите градусы: ")
	var Farengheit float64
	fmt.Scanf("%f", %Farengheit)
	Celsicius := ((Farengheit - 32) * 5/9))
	fmt.Println(Celsicius)
}
//№2
func main(){
	fmt.Println("Введите длинну")
	var Foots float64
	fmt.Scanf("%f", %Foots)
	Metr := (Foots * 0.3048)
	fmt.Println(Metr)
}