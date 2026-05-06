package handler

import (
    "agencycli/model"
    "agencycli/service"
    "agencycli/utils"
    "fmt"
)

type AgencyHandler struct {
    service *service.AgencyService
    input   *utils.InputHelper
}

func NewAgencyHandler(service *service.AgencyService) *AgencyHandler {
    return &AgencyHandler{
        service: service,
        input:   utils.NewInputHelper(),
    }
}

func (h *AgencyHandler) List(region string) {
    agencies, err := h.service.GetAgenciesByRegion(region)
    if err != nil {
        fmt.Println("Error loading agencies:", err)
        return
    }

    if len(agencies) == 0 {
        fmt.Println("No agencies found.")
        return
    }

    if region != "" {
        fmt.Printf("Filtered agencies for region: %s\n", region)
    } else {
        fmt.Println("All agencies:")
    }

    for _, agency := range agencies {
        fmt.Printf("  [%d] %s | %s | %s\n",
            agency.ID, agency.Name, agency.Region, agency.Phone)
    }
}

func (h *AgencyHandler) Get() {
    id, err := h.input.ReadUint("Enter ID: ")
    if err != nil {
        fmt.Println("Invalid ID format:", err)
        return
    }

    agency, err := h.service.GetAgencyByID(id)
    if err != nil {
        fmt.Println("Error loading agency:", err)
        return
    }

    if agency == nil {
        fmt.Printf("No agency found with ID %d.\n", id)
        return
    }

    h.printAgencyDetails(agency)
}

func (h *AgencyHandler) Create(flagRegion string) {
    name, err := h.input.ReadRequired("Enter Agency Name: ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    var region string
    if flagRegion != "" {
        region = flagRegion
        fmt.Printf("Using region from flag: %s\n", region)
    } else {
        region, err = h.input.ReadRequired("Enter Region: ")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
    }

    phone, err := h.input.ReadRequired("Enter Phone: ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    address, err := h.input.ReadRequired("Enter Address: ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    empCount, err := h.input.ReadUint32("Enter Employee Count (optional): ")
    if err != nil {
        fmt.Println("Invalid employee count format, skipping...")
        empCount = nil
    }

    agency, err := h.service.CreateAgency(name, address, phone, region, empCount)
    if err != nil {
        fmt.Println("Error creating agency:", err)
        return
    }

    fmt.Printf("Agency '%s' created successfully with ID %d.\n", agency.Name, agency.ID)
}

func (h *AgencyHandler) Edit(flagRegion string) {
    id, err := h.input.ReadUint("Enter ID to edit: ")
    if err != nil {
        fmt.Println("Invalid ID format:", err)
        return
    }

    agency, err := h.service.GetAgencyByID(id)
    if err != nil {
        fmt.Println("Error loading agency:", err)
        return
    }

    if agency == nil {
        fmt.Printf("No agency found with ID %d.\n", id)
        return
    }

    fmt.Println("\nCurrent agency details:")
    fmt.Printf("   Name: %s | Phone: %s | Address: %s\n",
        agency.Name, agency.Phone, agency.Address)

    name, _ := h.input.ReadWithDefault("Enter Agency Name", agency.Name)
    phone, _ := h.input.ReadWithDefault("Enter Agency Phone", agency.Phone)
    address, _ := h.input.ReadWithDefault("Enter Agency Address", agency.Address)

    var region string
    if flagRegion != "" {
        region = flagRegion
        fmt.Printf("Using region from flag: %s\n", region)
    } else {
        region, _ = h.input.ReadWithDefault("Enter Agency Region:", agency.Region)
    }

    empCount, _ := h.input.ReadUint32("Enter Agency Employee Count:")

    updated, err := h.service.UpdateAgency(id, name, address, phone, region, empCount)
    if err != nil {
        fmt.Println("Error updating agency:", err)
        return
    }

    fmt.Printf("Agency '%s' updated successfully.\n", updated.Name)
}

func (h *AgencyHandler) Status() {
    totalAgencies, totalEmployees, regionCount, err := h.service.GetStatus()
    if err != nil {
        fmt.Println("Error loading status:", err)
        return
    }

    if totalAgencies == 0 {
        fmt.Println("System Status: No agencies currently registered.")
        return
    }

    fmt.Println("╔═══════════════════════════════════════╗")
    fmt.Println("║          SYSTEM STATUS REPORT         ║")
    fmt.Println("╠═══════════════════════════════════════╣")
    fmt.Printf("║ Total Agencies: %-20d  ║\n", totalAgencies)
    fmt.Printf("║ Total Employees: %-19d  ║\n", totalEmployees)
    fmt.Println("╠═══════════════════════════════════════╣")
    fmt.Println("║ Agencies per Region:                  ║")

    for region, count := range regionCount {
        displayRegion := region
        if displayRegion == "" {
            displayRegion = "Unknown"
        }
        fmt.Printf("║   • %-10s: %-17d     ║\n", displayRegion, count)
    }
    fmt.Println("╚═══════════════════════════════════════╝")
}

func (h *AgencyHandler) printAgencyDetails(agency *model.Agency) {
    fmt.Println("\n┌─────────────────────────────────────┐")
    fmt.Printf("│ ID:           %d\n", agency.ID)
    fmt.Printf("│ Name:         %s\n", agency.Name)
    fmt.Printf("│ Region:       %s\n", agency.Region)
    fmt.Printf("│ Phone:        %s\n", agency.Phone)
    fmt.Printf("│ Address:      %s\n", agency.Address)
    if agency.EmployeeCount != nil {
        fmt.Printf("│ Employees:    %d\n", *agency.EmployeeCount)
    } else {
        fmt.Printf("│ Employees:    Not specified\n")
    }
    fmt.Printf("│ Member Since: %s\n", agency.MembershipDate.Format("2006-01-02"))
    fmt.Println("└─────────────────────────────────────┘")
}