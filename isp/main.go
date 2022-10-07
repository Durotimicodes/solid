package main

import "fmt"

func main() {

	doc := &Document{
		docName:  "Solid Principle Text book",
		nosPages: 203,
	}

	mM := MultiFunctionalMachine{}

	mM.Print(doc)
	s := mM.Fax(doc)
	f := mM.Scan(doc)
	fmt.Println(s, f)

}
