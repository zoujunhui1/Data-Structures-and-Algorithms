package main

import "fmt"

func fibonacci(data int) (res int) {
	if data <= 1 {
		res = 1
	}else {
		res = fibonacci(data - 1) + fibonacci(data - 2)
	}
	return res
}

func main(){
	for i:=0; i<=10; i++{
		fmt.Printf("---fibonacci(%d) is %d ---\n",i,fibonacci(i))
	}
}
