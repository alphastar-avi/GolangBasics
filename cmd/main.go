package main

import(
	f "fmt"
	"golangbasics/internals/utils"
)

func main(){
	// for array element sum
	arr := []int {1,2,3,4}
	f.Println(utils.ArrSumElem(arr))

	// map implementation
	var name string
	f.Print("enter the name: ")
	f.Scanln(&name)

	if age, check := utils.MapAge(name); check {
		f.Printf("%s is maybe %d years old\n", name , age)
	} else {
		f.Println("% age not found")
	}


}
