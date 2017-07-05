//#1
package main

import "fmt"

func main(){
	fmt.Println('1, 2, 3, 4, 5, 6, 7, 8, 9, 10')
}
//#2

func main(){
	for i = 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
//#3 
func main(){
	for i := 1; i <=10; i++{
		if i % 2 == 0{
			fmt.Println(i, "Even")
		}
		else {
			fmt.Println(i, "Odd")
		}	
	}
}
//#4
func main(){
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unkown Number")
	}
}
//Home Work #1
func main(){
	for i := 1; i <=10; i++{
		if i % 3 == 0{
			fmt.Println(i, "True")
		}
		else {
			fmt.Println("Error")
		}
	}
}
//#2
func main(){
	for i :=1; i <=100; i++{
		if i % 15 == 0{
			fmt.Println("FizzBuzz")
		}
		else if i % 3{
			fmt.Println("Fizz")
		}
		else if i % 5{
			fmt.Println("Buzz")
		}
		else {
			fmt.Println(i)
		}
	}
}