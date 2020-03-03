package main

import "fmt"

func printPrime(str string) {

next:
	for outer := 2;outer < 10; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				fmt.Println("456")
				continue next
			}
		}
		fmt.Printf("%s:%d\n", str, outer)
	}
	fmt.Println("Completed", str)
}

func main() {
	printPrime("A")
}
