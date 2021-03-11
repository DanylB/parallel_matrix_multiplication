package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

const N int = 200000000
const coresNum int = 8

const rowArrA int = 6
const columnArrA int = 7
const rowArrB int = 7
const columnArrB int = 8

var a = [rowArrA][columnArrA]int{}
var arrA = &a
var b = [rowArrB][columnArrB]int{}
var arrB = &b

var res1 = [1][columnArrB]int{}
var resArr1 = &res1
var res2 = [2][columnArrB]int{}
var resArr2 = &res2
var res3 = [3][columnArrB]int{}
var resArr3 = &res3
var res4 = [4][columnArrB]int{}
var resArr4 = &res4
var res5 = [5][columnArrB]int{}
var resArr5 = &res5
var res6 = [6][columnArrB]int{}
var resArr6 = &res6
var res7 = [7][columnArrB]int{}
var resArr7 = &res7

var wg sync.WaitGroup

// BenchmarkCount  func main() {
func BenchmarkCount(bench *testing.B) {

	runtime.GOMAXPROCS(coresNum)

	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	go generateArrayA()
	go generateArrayB()
	wg.Wait()

	printArrayA()
	printArrayB()

	addArrays()

	printRessultArray()

}

func generateArrayA() {
	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrA; j++ {
			a[i][j] = rand.Intn(10) + 1
		}
	}
	*arrA = a
	wg.Done()
	return
}

func generateArrayB() {
	for i := 0; i < rowArrB; i++ {
		for j := 0; j < columnArrB; j++ {
			b[i][j] = rand.Intn(10) + 1
		}
	}
	*arrB = b
	wg.Done()
	return

}

func addArrays() {
	wg.Add(6)
	go addLine1()
	go addLine2()
	go addLine3()
	go addLine4()
	go addLine5()
	go addLine6()
	wg.Wait()
}

func addLine1() {
	var res [1][columnArrB]int
	for i := 0; i < 1; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr1 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func addLine2() {
	var res [2][columnArrB]int
	for i := 1; i < 2; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr2 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func addLine3() {
	var res [3][columnArrB]int
	for i := 2; i < 3; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr3 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func addLine4() {
	var res [4][columnArrB]int
	for i := 3; i < 4; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr4 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func addLine5() {
	var res [5][columnArrB]int
	for i := 4; i < 5; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr5 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func addLine6() {
	var res [6][columnArrB]int
	for i := 5; i < 6; i++ {
		for j := 0; j < columnArrB; j++ {
			res[i][j] = 0
			for k := 0; k < rowArrB; k++ {
				res[i][j] += arrA[i][k] * arrB[k][j]
			}
			*resArr6 = res
		}
	}
	myBench()
	wg.Done()
	return
}

func printArrayA() {
	print("\n====================== Array A ===================\n")

	for i := 0; i < rowArrA; i++ {
		for j := 0; j < columnArrA; j++ {
			fmt.Printf("%d\t", a[i][j])
		}
		print("\n")
	}
}
func printArrayB() {
	print("\n========================= Array B ========================\n")
	for i := 0; i < rowArrB; i++ {
		for j := 0; j < columnArrB; j++ {
			fmt.Printf("%d\t", b[i][j])
		}
		print("\n")
	}
}

func printRessultArray() {

	print("\n========================= RESULT Array ========================\n")

	for j := 0; j < columnArrB; j++ {
		fmt.Printf("%d\t", res1[0][j])
	}
	print("\n")

	for j := 0; j < columnArrB; j++ {
		fmt.Printf("%d\t", res2[1][j])
	}
	print("\n")

	for j := 0; j < columnArrB; j++ {
		fmt.Printf("%d\t", res3[2][j])
	}
	print("\n")

	for j := 0; j < columnArrB; j++ {
		fmt.Printf("%d\t", res4[3][j])
	}
	print("\n")

	for j := 0; j < columnArrB; j++ {
		fmt.Printf("%d\t", res5[4][j])
	}
	print("\n")

	for j := 0; j < 8; j++ {
		fmt.Printf("%d\t", res6[5][j])
	}
	print("\n\n")
}

func myBench() {
	var d = 0
	for e := 0; e < N; e++ {
		d += e
	}
}
