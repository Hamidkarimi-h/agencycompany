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

    // Initialize dependencies
    repo := repository.NewAgencyRepository()
    service := service.NewAgencyService(repo)
    handler := handler.NewAgencyHandler(service)

    if len(os.Args) > 1 {
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
            flag.CommandLine.Init("myFlags", flag.ContinueOnError)
            regionPtr := flag.CommandLine.String("region", "", "Filter or apply to a specific region")
            flag.CommandLine.Parse(os.Args[2:])
            region = *regionPtr

            if region != "" {
                fmt.Printf("=> Executing command '%s' for region '%s'...\n", command, region)
            } else {
                fmt.Printf("=> Executing command '%s' for ALL regions...\n", command)
            }
        } else {
            fmt.Printf("=> Executing command '%s'...\n", command)
        }

        runCommand(command, region, handler)
        fmt.Println("\n✅ Done.")
        return
    }

    fmt.Println("Type 'help' for usage info.\n")

    scanner := bufio.NewScanner(os.Stdin)
    region := ""

    for {
        fmt.Print("Enter Command (list/get/create/edit/status/exit): ")
        if !scanner.Scan() {
            break
        }
        command := strings.TrimSpace(scanner.Text())

        if command == "" {
            continue
        }

        if command == "help" {
            printUsage()
            continue
        }

        if command == "exit" {
            fmt.Println("Goodbye!")
            break
        }

        commandsWithRegion := map[string]bool{
            "list":   true,
            "create": true,
            "edit":   true,
        }

        if commandsWithRegion[command] {
            fmt.Print("Enter Region (or press Enter for all): ")
            if scanner.Scan() {
                region = strings.TrimSpace(scanner.Text())
            }
        }

        runCommand(command, region, handler)
        region = ""
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Input error:", err)
    }
}

func runCommand(command string, region string, handler *handler.AgencyHandler) {
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
    default:
        fmt.Printf("Unknown command: %s\n", command)
    }
}