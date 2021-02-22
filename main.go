package main

import (
	"fmt"
	"math/rand"
	"time"
)

const columnArrA int = 6
const rowArrA int = 7

const columnArrB int = 7
const rowArrB int = 8

func main() {
	rand.Seed(time.Now().UnixNano())

	var a = generateArrayA()
	var b = generateArrayB()

	// var a = [2][3]int{{1, 1}, {2, 2, 2}}
	// var b = [3][4]int{{1, 1, 1}, {2, 2, 2, 2}}

	printArrayA(a)
	printArrayB(b)

	var resultArray = addArrays(a, b)

	printRessultArray(resultArray)

}

func generateArrayA() [columnArrA][rowArrA]int {
	var a = [columnArrA][rowArrA]int{}
	for i := 0; i < columnArrA; i++ {
		for j := 0; j < rowArrA; j++ {
			a[i][j] = rand.Intn(1000) + 1
		}
	}
	return a
}

func generateArrayB() [columnArrB][rowArrB]int {
	var b = [columnArrB][rowArrB]int{{}, {}}
	for i := 0; i < columnArrB; i++ {
		for j := 0; j < rowArrB; j++ {
			b[i][j] = rand.Intn(1000) + 1
		}
	}
	return b
}

func printArrayA(a [columnArrA][rowArrA]int) {
	print("\n====================== Array A ===================\n")

	for i := 0; i < columnArrA; i++ {
		for j := 0; j < rowArrA; j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		print("\n")
	}
}
func printArrayB(b [columnArrB][rowArrB]int) {
	print("\n========================= Array B ========================\n")
	for i := 0; i < columnArrB; i++ {
		for j := 0; j < rowArrB; j++ {
			fmt.Printf("%d\t", b[i][j])
		}
		print("\n")
	}
}

func printRessultArray(resultArray [columnArrA][rowArrB]int) {
	print("\n========================= RESULT Array ========================\n")
	for i := 0; i < columnArrA; i++ {
		for j := 0; j < rowArrB; j++ {
			fmt.Printf("%d\t", resultArray[i][j])
		}
		print("\n")
	}
}

func addArrays(a [columnArrA][rowArrA]int, b [columnArrB][rowArrB]int) [columnArrA][rowArrB]int {
	var i int
	var j int
	var k int

	var bt [8][7]int
	var res [6][8]int

	for i = 0; i < 8; i++ {
		for j = 0; j < 7; j++ {
			bt[i][j] = b[j][i]

		}
	}

	for i = 0; i < 6; i++ {
		for j = 0; j < 8; j++ {
			res[i][j] = 0
			for k = 0; k < 7; k++ {
				res[i][j] = a[i][k] * bt[j][k]
			}
		}
	}

	return res

}
