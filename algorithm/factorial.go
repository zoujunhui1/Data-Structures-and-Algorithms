package main

import "fmt"

/*
阶乘 3!=3*2*1
*/
func factorial(n int) (res int) {
	if(n <= 0){
		res = 1
	}else {
		res = n*factorial(n-1)
	}
	return
}

func main() {
	res := factorial(4)
	fmt.Printf("----factorial is %d---",res)
}
