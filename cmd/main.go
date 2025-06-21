package main

import(
	f "fmt"
	"net/http"
	// "golangbasics/internals/utils"
    ser "golangbasics/internals/server"
)

func main(){
	// // for array element sum
	// arr := []int {1,2,3,4}
	// f.Println("Array Sum: ", utils.ArrSumElem(arr))

	// // map implementation
	// var name string
	// f.Print("enter the name: ")
	// f.Scanln(&name)

	// if age, check := utils.MapAge(name); check {
	// 	f.Printf("%s is maybe %d years old\n", name , age)
	// } else {
	// 	f.Println("% age not found")
	// }

	// //Sort in acending order 
	// nums := []int {7,8,2,4,0,5,0,3,9,3}
	// utils.SortArrayAes(nums)
	// f.Println("Decending order: ")
	// for _, value := range nums {
	// 	f.Print(value, ",")
	// }
	// //Sort in decending order
	// utils.SortArrayDes(nums)
	// f.Println("Ascending order: ")
	// for _, value := range nums {
	// 	f.Print(value, ",")
	// }
	
	//serverrr

	ser.SetupRoute()
	f.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
