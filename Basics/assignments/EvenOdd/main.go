package main

// print even odd numbers

func main() {

	num1 := [10]int{2, 1, 3, 9, 5, 6, 7, 8, 4, 10}

	for i := range num1 {
		if num1[i]%2 == 0 {
			println(num1[i], " is Even number ")
		} else {
			println(num1[i], " is Odd number ")

		}

	}
}
