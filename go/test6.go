package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int 

	c = a + b
	fmt.Printf("the value of the -c in the first line is %d\n", c)
	c = a - b 
	fmt.Printf("the value of the -c in the second line is %d\n", c)
	c = a * b
	fmt.Printf("the value of the -c in the third line is %d\n", c)
	c = a / b
	fmt.Printf("the value of the -c in the fourth line is %d\n", c)
	c = a % b
	fmt.Printf("the value of the -c in the fifth line is %d\n", c)
	a++
	fmt.Printf("the value of the -a in the sixth line is %d\n", a)
	a=21
	a--
	fmt.Printf("the value of the -a in the seventh line is %d\n", a)
}
