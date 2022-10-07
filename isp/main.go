package main

func main() {

	doc := &Document{
		docName:  "Solid Principle Text book",
		nosPages: 203,
	}

	mM := MultiFunctionalMachine{}
	mM.Print(doc)

}
