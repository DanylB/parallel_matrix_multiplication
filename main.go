package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rowArrA int = 6
const columnArrA int = 7

const rowArrB int = 7
const columnArrB int = 8

func main() {
	rand.Seed(time.Now().UnixNano())

	var a = generateArrayA()
	var b = generateArrayB()

	printArrayA(a)
	printArrayB(b)

	var resultArray = addArrays(a, b)

	printRessultArray(resultArray)

}

func generateArrayA() [rowArrA][columnArrA]int {
	var a = [rowArrA][columnArrA]int{}
	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrA; j++ {
			a[i][j] = rand.Intn(1000) + 1
		}
	}
	return a
}

func generateArrayB() [rowArrB][columnArrB]int {
	var b = [rowArrB][columnArrB]int{{}, {}}
	for i := 0; i < rowArrB; i++ {
		for j := 0; j < columnArrB; j++ {
			b[i][j] = rand.Intn(1000) + 1
		}
	}
	return b
}

func printArrayA(a [rowArrA][columnArrA]int) {
	print("\n====================== Array A ===================\n")

	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrA; j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		print("\n")
	}
}
func printArrayB(b [rowArrB][columnArrB]int) {
	print("\n========================= Array B ========================\n")
	for i := 0; i < rowArrB; i++ {
		for j := 0; j < columnArrB; j++ {
			fmt.Printf("%d\t", b[i][j])
		}
		print("\n")
	}
}

func printRessultArray(resultArray [rowArrA][columnArrB]int) {
	print("\n========================= RESULT Array ========================\n")
	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrB; j++ {
			fmt.Printf("%d\t", resultArray[i][j])
		}
		print("\n")
	}
}

func addArrays(a [rowArrA][columnArrA]int, b [rowArrB][columnArrB]int) [rowArrA][columnArrB]int {

	var res [rowArrA][columnArrB]int

	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < columnArrA; k++ {
				res[i][j] = a[i][k] * b[k][j]
			}
		}
	}
	return res
}
