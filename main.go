package main

import (
	"bytes"
	"encoding/json"
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

const datafile = "./data.json"

func loadAgencies()([]Agency, error){
	var agencies []Agency

	data, err := os.ReadFile(datafile)
	if err != nil{
		
		if os.IsNotExist(err){
			return  []Agency{}, nil	
		}
		return  nil, err
	}

	if len(bytes.TrimSpace(data)) > 0 {
		if err := json.Unmarshal(data, &agencies); err != nil {
			return nil, err
		}
	}

	return agencies, nil

}
func saveAgencies(agencies []Agency) error {
	data, err := json.MarshalIndent(agencies, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(datafile, data, 0644)
}

func printUsage() {
    fmt.Println("Usage: go run main.go <command> [options]")
    fmt.Println("Commands: list, get, create, edit, status")
    fmt.Println("Example: go run main.go list -region=Tehran")
}

func main()  {
	
	if len(os.Args) < 2{
		printUsage()
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
	runCommand(command, *region)
}

func runCommand(command string, region string) {
    switch command {
    case "list":
        ListAgencies(region)
    case "create":
        fmt.Println("Create command not implemented yet.")
    case "edit":
        fmt.Println("Edit command not implemented yet.")
    case "get":
        fmt.Println("Get command not implemented yet.")
    case "status":
        fmt.Println("Status command not implemented yet.")
	case "exit":
		os.Exit(0)
    default:
        fmt.Printf("Unknown command: %s\n", command)
        printUsage()
    }
}
func ListAgencies(region string) {
	agencies, err := loadAgencies()
	if err != nil {
		fmt.Println("Can't load agencies:", err)
		return
	}

	if len(agencies) == 0 {
		fmt.Println("No agencies found.")
		return
	}

	if region != "" {
		fmt.Println("Filtered agencies for region:", region)
	} else {
		fmt.Println("All agencies:")
	}

	found := false
	for _, agency := range agencies {
	
		if region == "" || agency.Region == region {
			fmt.Printf("ID: %d, Name: %s, Region: %s\n", agency.ID, agency.Name, agency.Region)
			found = true
		}
	}

	if !found && region != "" {
		fmt.Printf("No agencies found for region '%s'.\n", region)
	}
}