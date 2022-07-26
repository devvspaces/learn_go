package main

import (
	"fmt"
	"math"
)

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	allowed_flags := []string{
// 		"-f",
// 	}
// 	if len(os.Args) > 1 {
// 		flag := os.Args[1]
// 		found := false
// 		for _, value := range allowed_flags {
// 			if flag == value {
// 				found = true
// 				file_name := os.Args[2]
// 				fmt.Println(file_name)
// 				break
// 			}
// 		}
// 		if found == false {
// 			panic("Invalid Flag used")
// 		}
// 	}
// }

func half(num uint) (value bool) {
	num = num / 2
	if num%2 == 0 {
		value = true
	} else {
		value = false
	}
	return
}

func max(nums ...int) int {
	max := nums[0]
	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	return max
}

func makeOddGenerator() func() int {
	num := -1
	return func() int {
		// defer func() { num += 2 }()
		num += 2
		return num
	}
}

func fib(n uint) uint {
	if (n == 0) || (n == 1) {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n uint) uint {
	float_n := float64(n)
	numerator := math.Pow(1+math.Sqrt(5), float_n)
	denominator := math.Pow(2, float_n) * math.Sqrt(5)
	return uint(math.Abs(numerator / denominator))
}

func main() {
	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(98, 93, 77, 82, 83))
	// fmt.Println(max(1, 2, 3, 4, 5, 6, 7, 9, 23))
	// nextOdd := makeOddGenerator()
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	defer fmt.Println("Hello Third Call")
	defer fmt.Println("Hello")
	fmt.Println(fib2(14))
}
