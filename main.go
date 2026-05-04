package main

import (
	"flag"
	"fmt"
)
type Agency struct{
	ID uint
	Name string
	Address string
	Phone string
	MembershipDate string
	EmployeeCount uint32
	Region Region
}
type Region struct{
	ID uint
	Name string
}

func main()  {
	command := flag.String("command", "", "Command to execute: list, get, create, edit, status")
	region := flag.String("region", "", "Filter or apply to a specific region (optional)")

	flag.Parse()

	if *command == ""{
		fmt.Println("Error: Please provide a command.")
		fmt.Println("Example: go run main.go -command=list -region=Tehran")
		fmt.Println("Available commands: list, get, create, edit, status")
		
		return
	}

	if *region != "" {
		fmt.Printf("=> Executing command '%s' for region '%s'...\n", *command, *region)
	} else {
		fmt.Printf("=> Executing command '%s' for ALL regions...\n", *command)
	}
}