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

	// var a = generateArrayA()
	// var b = generateArrayB()

	var a = [6][7]int{
		{1, 2, 3, 4, 4, 3, 2},
		{4, 5, 6, 6, 6, 5, 4},
		{4, 5, 6, 7, 7, 6, 5},
		{1, 2, 3, 4, 4, 3, 2},
		{4, 5, 6, 6, 6, 5, 4},
		{4, 5, 6, 7, 7, 6, 5},
	}

	var b = [7][8]int{
		{1, 2, 3, 4, 4, 3, 2, 1},
		{4, 5, 6, 6, 6, 5, 4, 3},
		{4, 5, 6, 7, 7, 6, 5, 4},
		{1, 2, 3, 4, 4, 3, 2, 1},
		{4, 5, 6, 6, 6, 5, 4, 3},
		{4, 5, 6, 7, 7, 6, 5, 4},
		{4, 5, 6, 7, 7, 6, 5, 4},
	}

	printArrayA(a)
	printArrayB(b)

	// var resultArray = addArrays(a, b)

	// printRessultArray(resultArray)

	addArrays(a, b)

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

// func addArrays(a [columnArrA][rowArrA]int, b [columnArrB][rowArrB]int) [columnArrA][rowArrB]int {
// 	var i int
// 	var j int
// var k int
//
// 	var bt [8][7]int
// 	var res [6][8]int
//
// 	for i = 0; i < 8; i++ {
// 		for j = 0; j < 7; j++ {
// 			bt[i][j] = b[j][i]
//
// 		}
// 	}
//
// 	for i = 0; i < 6; i++ {
// 		for j = 0; j < 8; j++ {
// 			res[i][j] = 0
// 			for k = 0; k < 7; k++ {
// 				res[i][j] = a[i][k] * bt[j][k]
// 			}
// 		}
// 	}
//
// 	return res
//
// }

func addArrays(a [rowArrA][columnArrA]int, b [rowArrB][columnArrB]int) {

	addLine1(a, b)
	addLine2(a, b)

}

func addLine1(a [rowArrA][columnArrA]int, b [rowArrB][columnArrB]int) {

	var i int
	var j int
	var k int

	var res [1][8]int

	for i = 0; i < 1; i++ {
		for j = 0; j < 8; j++ {
			res[i][j] = 0
			for k = 0; k < 7; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
			print(res[i][j], " ")
		}

	}
	print("\n")
	return
}

func addLine2(a [rowArrA][columnArrA]int, b [rowArrB][columnArrB]int) {

	var i int
	var j int
	var k int

	var res [2][8]int

	for i = 1; i < 2; i++ {
		for j = 0; j < 8; j++ {
			res[i][j] = 0
			for k = 0; k < 7; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
			print(res[i][j], " ")
		}

	}
	print("\n")
	return
}
