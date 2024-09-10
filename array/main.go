package main
import "fmt"
// slice decl
var sliceArr []int

func arrayAddEle () {
	for  i :=0; i<5; i++ {
		sliceArr =append(sliceArr, i+1)
	}
	fmt.Println(sliceArr)
}

func arrayTraversal () {
	for  i :=0; i<5; i++ {
		fmt.Println(sliceArr[i])
		}
}

	func getEle (ele int) int  {
			return sliceArr[ele]
		}

func main () {
	arrayAddEle()
	arrayTraversal()
	fmt.Println("Element at index 2 :", getEle(1))
}

