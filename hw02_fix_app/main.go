package main

import (
	"fmt" //nolint:gci
	"hw02_fix_app/printer"
	"hw02_fix_app/reader"
	"hw02_fix_app/types"
	//nolint:gci
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
