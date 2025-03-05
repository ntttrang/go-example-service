package internal

func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// func Adderx() int {
// 	sum := 0
// 	f := func(x int) int {
// 		sum += x
// 		return sum
// 	}
// 	fcall := f(3)
// 	fmt.Println(fcall)
// 	return 0
// }
