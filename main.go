package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
	"strconv"
	
)
type Agency struct{
	ID uint
	Name string
	Address string
	Phone string
	MembershipDate time.Time
	EmployeeCount *uint32
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
        CreateAgency(region)
    case "edit":
        EditAgency(region)
    case "get":
        GetAgency()
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


func CreateAgency(flagRegion string) {

    agencies, err := loadAgencies()
    if err != nil {
        fmt.Println("Error loading agencies:", err)
        return
    }


    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Agency Name: ")
    name, nErr := reader.ReadString('\n')
    if nErr != nil {
        fmt.Println("Error reading name:", nErr)
        
		return
    }
    name = strings.TrimSpace(name)
    if name == "" {
        fmt.Println("Agency name cannot be empty.")
        
		return
    }

	var region string
    if flagRegion != "" {
        region = flagRegion
    } else {
        fmt.Print("Enter Region: ")
		region, rErr := reader.ReadString('\n')
		if rErr != nil {
			fmt.Println("Error reading region:", rErr)
			
			return
		}
		region = strings.TrimSpace(region)
		if region == "" {
			fmt.Println("Region cannot be empty.")
			
			return
		}
    }
    
	
	fmt.Print("Enter Phone: ")
    phone, pErr := reader.ReadString('\n')
    if pErr != nil {
        fmt.Println("Error reading phone:", pErr)
        
		return
    }
    phone = strings.TrimSpace(phone)
    if phone == "" {
        fmt.Println("Phone cannot be empty.")
        
		return
    }

	fmt.Print("Enter Address: ")
    address, aErr := reader.ReadString('\n')
    if aErr != nil {
        fmt.Println("Error reading address:", aErr)
        
		return
    }
    address = strings.TrimSpace(address)
    if address == "" {
        fmt.Println("Address cannot be empty.")
        
		return
    }


	fmt.Print("Enter Employee Count (optional, press Enter to skip): ")
    empCountStr, eErr := reader.ReadString('\n')
	if eErr != nil {
        fmt.Println("Error reading employee count:", eErr)
        
		return
    }
    empCountStr = strings.TrimSpace(empCountStr)
	var empCount *uint32
	if empCountStr != "" {
    count, _ := strconv.ParseUint(empCountStr, 10, 32)
    val := uint32(count)
    empCount = &val 
	}
    
    var maxID uint
    for _, a := range agencies {
        if a.ID > maxID {
            maxID = a.ID 
        }
    }
    newAgency := Agency{
        ID:             maxID + 1,
        Name:           name,
        Address:        address, 
        Phone:          phone,
        MembershipDate: time.Now(),
        EmployeeCount:  empCount,
        Region:         region,
    }


    agencies = append(agencies, newAgency)
    if err := saveAgencies(agencies); err != nil {
        fmt.Println("Error saving agency:", err)
        return
    }

   
    fmt.Printf("Agency '%s' created successfully with ID %d.\n", newAgency.Name, newAgency.ID)
}

func GetAgency(){
	agencies, err := loadAgencies()
	if err != nil {
		fmt.Println("Error loading agencies:", err)
		
		return
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter ID: ") 

	strID, iErr := reader.ReadString('\n')
	if iErr != nil {
		fmt.Println("Error reading ID:", iErr)
		
		return
	}
	strID = strings.TrimSpace(strID)
	if strID == "" {
		fmt.Println("ID cannot be empty.")
		
		return
	}
	u64, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		fmt.Println("Invalid number format:", err)
		
		return
	}

	ID := uint(u64)
	found := false

	for _, agency := range agencies {

		if ID == agency.ID {
			fmt.Printf("ID: %d\nName: %s\nRegion: %s\nPhone: %s\nAddress: %s\n", 
    			agency.ID, agency.Name, agency.Region, agency.Phone, agency.Address)

			if agency.EmployeeCount != nil {
				fmt.Printf("Employee Count: %d\n", *agency.EmployeeCount)
			} else {
				fmt.Println("Employee Count: Not specified")
			}
			found = true			
			break 
		}
	}

	if !found {
		fmt.Printf("No agency found with ID %d.\n", ID)
	}
}

func EditAgency(flagRegion string){
	agencies, err := loadAgencies()
	if err != nil {
		fmt.Println("Error loading agencies:", err)
		
		return
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter ID: ") 

	strID, iErr := reader.ReadString('\n')
	if iErr != nil {
		fmt.Println("Error reading ID:", iErr)
		
		return
	}
	strID = strings.TrimSpace(strID)
	if strID == "" {
		fmt.Println("ID cannot be empty.")
		
		return
	}
	u64, err := strconv.ParseUint(strID, 10, 0)
	if err != nil {
		fmt.Println("Invalid number format:", err)
		
		return
	}

	ID := uint(u64)
	found := -1

	for index, agency := range agencies {

		if ID == agency.ID {
			fmt.Printf("ID: %d\nName: %s\nRegion: %s\nPhone: %s\nAddress: %s\n", 
    			agency.ID, agency.Name, agency.Region, agency.Phone, agency.Address)

			if agency.EmployeeCount != nil {
				fmt.Printf("Employee Count: %d\n", *agency.EmployeeCount)
			} else {
				fmt.Println("Employee Count: Not specified")
			}
			found = index			
			break 
		}
	}
	if found == -1 {
    	fmt.Printf("No agency found with ID %d.\n", ID)
    	return
	}


    fmt.Printf("Enter Agency Name [%s]: ", agencies[found].Name)
    name, nErr := reader.ReadString('\n')
    if nErr != nil {
        fmt.Println("Error reading name:", nErr)
        
		return
    }else{
		name = strings.TrimSpace(name)
		if name != ""{
			agencies[found].Name = name
		}
	}
    
    fmt.Printf("Enter Agency Phone [%s]: ", agencies[found].Phone)
    phone, pErr := reader.ReadString('\n')
    if pErr != nil {
        fmt.Println("Error reading phone:", pErr)
        
		return
    }else{
		phone = strings.TrimSpace(phone)
		if phone != ""{
			agencies[found].Phone = phone
		}
	}

	fmt.Printf("Enter Agency Address [%s]: ", agencies[found].Address)
    address, aErr := reader.ReadString('\n')
    if aErr != nil {
        fmt.Println("Error reading address:", aErr)
        
		return
    }else{
		address = strings.TrimSpace(address)
		if address != ""{
			agencies[found].Address = address
		}
	}


	if agencies[found].EmployeeCount != nil {
    	fmt.Printf("Enter Agency Employee [%d]: ", *agencies[found].EmployeeCount)
	}else {
		fmt.Print("Enter Agency Employee [Not specified]: ")
	}
    empCountStr, eErr := reader.ReadString('\n')
	if eErr != nil {
        fmt.Println("Error reading employee count:", eErr)
        
		return
    }else{
		empCountStr = strings.TrimSpace(empCountStr)
		
		if empCountStr != "" {
		count, _ := strconv.ParseUint(empCountStr, 10, 32)
		val := uint32(count)
		agencies[found].EmployeeCount = &val 
		}

	}

	if flagRegion != "" {
		agencies[found].Region = flagRegion
	} else {
		fmt.Printf("Enter Agency Region [%s]: ", agencies[found].Region)
		newRegion, rErr := reader.ReadString('\n')
		if rErr != nil {
			fmt.Println("Error reading region:", rErr)
			
			return
		}
		newRegion = strings.TrimSpace(newRegion)
		if newRegion != "" {
			agencies[found].Region = newRegion
		}
	}

	if err := saveAgencies(agencies); err != nil {
    	fmt.Println("Error saving agencies:", err)
    	
		return
	}
	fmt.Printf("Agency '%s' updated successfully.\n", agencies[found].Name)

}
