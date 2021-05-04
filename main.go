package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gitinreallife/greetings"
	"github.com/gitinreallife/morestrings"
	//  go mod edit --replace=github.com/gitinreallife/govars=../govars
)

// go mod init gofirst(package name)
func main() {

	fmt.Println("##################################################")
	PrintArrays()
	// PrintVars()
	fmt.Println("##################################################")
	log.SetPrefix("Exception greetings: ")
	log.SetFlags(0)
	name := ""
	fmt.Println(("What's your name? "))
	fmt.Scanln(&name)
	strQuote, err := greetings.GetQuote(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#####Calling Greeting Module")
	fmt.Println(strQuote)
	fmt.Println("#####Calling Reverse Module")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

}

const myConst int = 40
const (
	_  = iota
	KB = 1 << (10 * iota)
	// << is called bit shifting from integer to bit
	GB
	TB
)

const (
	KB2 = iota
)

func PrintArrays() {
	fmt.Println("###############1 D Array Playground###############")

	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Jisoo"
	students[2] = "Rosé"
	fmt.Printf("Students: %v\n", students)

	fmt.Println("###############Multidimentional array Playground###############")
	var idMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	//or declare it like this
	idMatrix[0] = [3]int{1, 0, 0}
	idMatrix[1] = [3]int{0, 1, 0}
	idMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(idMatrix)

	fmt.Println("############### Pointer array Playground###############")
	//start pointer array
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("############### Slice array Playground###############")
	//slice
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Slice Length: " + strconv.Itoa(len(sliceA)))
	fmt.Println("Slice Capacity: " + strconv.Itoa(cap(sliceA)))

	sliceB := sliceA[:]   //slice all elements
	sliceC := sliceB[3:]  //slice from 4th element to end
	sliceD := sliceB[:6]  //slice untill 6th elements
	sliceE := sliceB[3:6] //slice from 4-6
	sliceA[5] = 42
	fmt.Println(sliceA)
	fmt.Println(sliceB)
	fmt.Println(sliceC)
	fmt.Println(sliceD)
	fmt.Println(sliceE)

	makeA := make([]int, 3, 6) //create slice len 3 cap 6
	makeA = append(makeA, 1)
	makeA = append(makeA, []int{2, 3, 4, 5}...)
	fmt.Println(makeA)
	fmt.Printf("Length Make: %v\n", len(makeA))
	fmt.Printf("Capacity Make: %v\n", cap(makeA))

	//removing slice
	makeB := makeA[3:8]
	fmt.Println(makeB)

	makeC := append(makeB[:1], makeB[2:]...)
	fmt.Println(makeC)
}

func PrintVars() {

	fmt.Println("###############Variables Playground###############")
	fmt.Println("#####bool")
	n := 1 == 1
	fmt.Printf("%v, %T\n", n, n)

	fmt.Println("#####int")
	a := 10.2 //1010
	b := 3.1  //0011
	fmt.Printf("%v, %T\n", a, a)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(float32(a) / float32(b))
	// fmt.Println(a % b)
	// fmt.Println("----")

	// fmt.Println(a & b)  //binary boolean 1010 & 0011 = 0010 = int 2
	// fmt.Println(a | b)  //binary boolean 1010 | 0011 = 1011 = int 11
	// fmt.Println(a ^ b)  //binary boolean 1010 not or 0011 = 1001
	// fmt.Println(a &^ b) //binary boolean 1010 not and 0011 = 0100

	fmt.Println("#####")
	fmt.Println("###############End of Variables Playground#####")
	fmt.Println("###############Const Playground###############")
	const myConst int = iota
	fmt.Printf("%v, %T\n", KB, KB)
	fmt.Printf("%v, %T\n", GB, GB)
	fmt.Printf("%v, %T\n", TB, TB)
	fmt.Printf("%v, %T\n", KB2, KB2)
	fmt.Printf("%v\n", KB == KB2)

	filesize := 30000000.
	fmt.Printf("%.2fGB \n", filesize/GB)

	fmt.Println("-----Roles bitwise & bitshifting------")
	const (
		isAdmin = 1 << iota
		isHQ
		canSeeFinance

		canSeeAccounting
		canSeeMarketing
		canSeeIT
	)
	var roles byte = isAdmin | canSeeFinance | canSeeIT
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Println("###############End ofConst Playground##########")
	fmt.Println("###############################################")
}
