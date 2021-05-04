package main

import (
	"fmt"
	"log"

	"github.com/gitinreallife/greetings"
	"github.com/gitinreallife/morestrings"
	// go mod edit --replace=github.com/gitinreallife/greetings=../greetings
)

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

// go mod init gofirst(package name)
func main() {

	fmt.Println("##################################################")
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

	fmt.Println("-----Roles bit------")
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
