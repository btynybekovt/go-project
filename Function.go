//#1
package main

import "fmt"

func average(xs []float64)float64{
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float(len(xs))
}
func main(){
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}

//#2

func makeEvenGenerator() func() uint(){
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}
func main(){
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven)
	fmt.Println(nextEven)
	fmt.Println(nextEven)
}
//#3
func factorial(x uint) uint{
	if x == 0{
		return 1
	}
	return x * factorial(x-1)
}
//#4
func first(){
	fmt.Println("1st")
}
func second(){
	fmt.Println("2nd")
}
func main(){
	defer first()
	second()
}
//#5
func main(){
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	panic("Panic")
}