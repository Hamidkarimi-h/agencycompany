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