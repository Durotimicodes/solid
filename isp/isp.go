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
	Scan(d *Document) string
	Fax(d *Document) string
}
type MultiFunctionalMachine struct{}

func (m *MultiFunctionalMachine) Print(doc *Document) {
	fmt.Println(doc)
}
func (m *MultiFunctionalMachine) Scan(doc *Document) string {
	var s string
	if doc != nil {
		s += "Doc is scanned successfully"
	}

	return s
}
func (m *MultiFunctionalMachine) Fax(doc *Document) string {

	var s string
	if doc != nil {
		s += "Doc is faxed successfully"
	}

	return s
}

// type OldMachine struct{}

// func (old *OldMachine) Print(doc )
