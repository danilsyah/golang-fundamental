package main

import "fmt"

// parameter variabel argument / variadic , variadic hanya bisa di parameter terakhir
func sumAll(numbers ...int)int{
	total := 0
	
	for _, value := range numbers{
		total += value
	}
	return total
}

func main() {
	total := sumAll(2,5,10,22)
	fmt.Println(total)

	slice := []int{10,33,21,43}
	total = sumAll(slice...)
	fmt.Println(total)

}