package main

import f "fmt"

func add(a int, b int) int {
    f.Println(a + b)
    return a + b
}

func main() {
    f.Println("hello, meow")
    var name string = "avinash"
    f.Println(name)
    age := 18
    age2 := 20
    f.Println(age)

    add(age, age2)

	if age >= 18{
		f.Println("hey bud itme for responsibility")
	} else {
		f.Println("lil bud")
	}

	arr := []int{1,2,3,4,5,6}
	for x:=0 ; x<len(arr) ; x++ {
		f.Println(arr[x])
	}

	m := map[string]int{"avinash": 34}
	f.Println(m["avinash"])

}