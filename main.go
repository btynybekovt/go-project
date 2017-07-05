package main
//импорт данных из одного файла в другой
import "fmt"
import "go/src/math"

func main() {
    xs := []float64{1,2,3,4}
    avg := math.Average(xs)
    fmt.Println(avg)
}