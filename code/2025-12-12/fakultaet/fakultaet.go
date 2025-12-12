package main

func Factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result = result * i
	}
	return result
}

func Factorial_v2(n int) int {
	if n == 0 {
		return 1 //da 0!=1
	}
	return n * Factorial_v2(n-1)
	//Rekursion
	//da 5!=5*4*3*2*1
	//5! = 5*4!
	//4! = 4*3!...
}
