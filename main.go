package main

import (
	"fmt"
	"log"

	"github.com/gitinreallife/greetings"
	"github.com/gitinreallife/morestrings"
	// go mod edit --replace=github.com/gitinreallife/greetings=../greetings
)

// go mod init gofirst(package name)
func main() {
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
