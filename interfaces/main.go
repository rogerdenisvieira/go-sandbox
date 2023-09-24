package main

import "fmt"

type File interface {
	Open()
	Read()
}

type CSV struct {
	Name      string
	Delimiter string
}

type XML struct {
	Name    string
	Version string
}

func (c CSV) Read() {
	fmt.Println("Reading a CSV file:", c.Name)
}

func (c CSV) Open() {
	fmt.Println("Opening CSV file:", c.Name)
}

func main() {

	var file File = CSV{Name: "finances.csv", Delimiter: ";"}

	file.Open()
	file.Read()

}
