package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)
type Agency struct{
	ID uint
	Name string
	Address string
	Phone string
	MembershipDate time.Time
	EmployeeCount uint32
	Region string
}


func main()  {
	
	if len(os.Args) < 2{
		fmt.Println("Error: Please provide a command.")
		fmt.Println("Example: go run main.go list -region=Tehran")
		fmt.Println("Available commands: list, get, create, edit, status")
		return
	}
	command := os.Args[1]
	myFlags := flag.NewFlagSet("myFlags", flag.ExitOnError)
	region := myFlags.String("region", "", "Filter or apply to a specific region")

	myFlags.Parse(os.Args[2:])

	
	if *region != "" {
		fmt.Printf("=> Executing command '%s' for region '%s'...\n", command, *region)
	} else {
		fmt.Printf("=> Executing command '%s' for ALL regions...\n", command)
	}
}