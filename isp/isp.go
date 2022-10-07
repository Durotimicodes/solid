package main

import "fmt"

/*
The interface segregation principle states that you do not necessarily have to put all your methods
into an interface, in some cases its better to segregate this methods
*/

// for example
type Document struct {
	docName  string
	nosPages int
}

type Machine interface {
	Print(d *Document)
	Scan(d *Document)
	Fax(d *Document)
}
type MultiFunctionalMachine struct{}

func (m *MultiFunctionalMachine) Print(doc *Document) {
	fmt.Println(doc)
}
func (m *MultiFunctionalMachine) Scan(doc *Document) string {
	s := fmt.Sprintf("Doc is scanned successfully %v", doc)
	return s
}
func (m *MultiFunctionalMachine) Fax(doc *Document) string {
	s := fmt.Sprintf("Doc is faxed successfully %v", doc)
	return s
}


