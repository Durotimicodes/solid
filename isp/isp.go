package main

import "fmt"

/*
The interface segregation principle states that you do not necessarily have to put all your methods
into an interface, in some cases its better to segregate this methods

try to break-up interface methods into parts people will definately need
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

type OldMachine struct{}

func (old *OldMachine) Print(doc *Document) {
	fmt.Println("Old machine printing document",doc)
}

/* TO IMPLEMENT THE ISP */
type Printer interface{
	Print(d Document)
}

type Scanner interface{
	Scan(d Document)
}

type MyPrinter struct{}

func (mP *MyPrinter) Print(doc Document){}

type Photocopier struct{}

func (p *Photocopier) Print(doc Document){}
func (p *Photocopier) Scan(doc Document){}