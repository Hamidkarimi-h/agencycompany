package main

import (
    "agencycli/handler"
    "agencycli/repository"
    "agencycli/service"
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

func printUsage() {
    fmt.Println("Usage: go run main.go <command> [options]")
    fmt.Println("Commands: list, get, create, edit, status, exit")
    fmt.Println("Options:")
    fmt.Println("  -region=<name>  Filter or apply to a specific region (list, create, edit)")
    fmt.Println("Examples:")
    fmt.Println("  go run main.go list")
    fmt.Println("  go run main.go list -region=Tehran")
    fmt.Println("  go run main.go create -region=Tehran")
}

func main() {
    fmt.Println("╔══════════════════════════════════════════╗")
    fmt.Println("║   Welcome to Agency Management CLI!      ║")
    fmt.Println("╚══════════════════════════════════════════╝")

    if len(os.Args) < 2 {
        printUsage()
        return
    }

    command := os.Args[1]

    if command == "exit" {
        fmt.Println("Goodbye!")
        os.Exit(0)
    }


    var region string
    commandsWithRegion := map[string]bool{
        "list":   true,
        "create": true,
        "edit":   true,
    }

    if commandsWithRegion[command] {
        myFlags := flag.NewFlagSet("myFlags", flag.ExitOnError)
        regionPtr := myFlags.String("region", "", "Filter or apply to a specific region")
        myFlags.Parse(os.Args[2:])
        region = *regionPtr

        if region != "" {
            fmt.Printf("=> Executing command '%s' for region '%s'...\n", command, region)
        } else {
            fmt.Printf("=> Executing command '%s' for ALL regions...\n", command)
        }
    } else {
        fmt.Printf("=> Executing command '%s'...\n", command)
    }

    // Initialize dependencies
    repo := repository.NewAgencyRepository()
    service := service.NewAgencyService(repo)
    handler := handler.NewAgencyHandler(service)


    if region != "" {
        runCommand(command, region, handler)
        fmt.Println("\n Command executed. Exiting.")
        return
    }

    scanner := bufio.NewScanner(os.Stdin)
    shouldContinue := true

    for shouldContinue {
        fmt.Print("\nEnter Command (list/get/create/edit/status/exit): ")
        if !scanner.Scan() {
            break
        }
        command = strings.TrimSpace(scanner.Text())

        if command == "" {
            continue
        }


        if commandsWithRegion[command] && region == "" {
            fmt.Print("Enter Region (or press Enter for all): ")
            if scanner.Scan() {
                region = strings.TrimSpace(scanner.Text())
            }
        }

        shouldContinue = runCommand(command, region, handler)

        region = ""
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Input error:", err)
    }
    fmt.Println("Goodbye!")
}

func runCommand(command string, region string, handler *handler.AgencyHandler) bool {
    switch command {
    case "list":
        handler.List(region)
    case "create":
        handler.Create(region)
    case "edit":
        handler.Edit(region)
    case "get":
        handler.Get()
    case "status":
        handler.Status()
    case "exit":
        return false
    default:
        fmt.Printf("Unknown command: %s\n", command)
        printUsage()
    }
    return true
}