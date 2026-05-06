package service

import (
    "agencycli/model"
    "agencycli/repository"
    "time"
)

type AgencyService struct {
    repo *repository.AgencyRepository
}

func NewAgencyService(repo *repository.AgencyRepository) *AgencyService {
    return &AgencyService{repo: repo}
}

func (s *AgencyService) GetAllAgencies() ([]model.Agency, error) {
    return s.repo.LoadAll()
}

func (s *AgencyService) GetAgenciesByRegion(region string) ([]model.Agency, error) {
    agencies, err := s.repo.LoadAll()
    if err != nil {
        return nil, err
    }

    if region == "" {
        return agencies, nil
    }

    var filtered []model.Agency
    for _, a := range agencies {
        if a.Region == region {
            filtered = append(filtered, a)
        }
    }
    return filtered, nil
}

func (s *AgencyService) GetAgencyByID(id uint) (*model.Agency, error) {
    agencies, err := s.repo.LoadAll()
    if err != nil {
        return nil, err
    }
    return s.repo.FindByID(agencies, id), nil
}

func (s *AgencyService) CreateAgency(name, address, phone, region string, employeeCount *uint32) (*model.Agency, error) {
    agencies, err := s.repo.LoadAll()
    if err != nil {
        return nil, err
    }

    newAgency := model.Agency{
        ID:             s.repo.GetMaxID(agencies) + 1,
        Name:           name,
        Address:        address,
        Phone:          phone,
        MembershipDate: time.Now(),
        EmployeeCount:  employeeCount,
        Region:         region,
    }

    agencies = append(agencies, newAgency)

    if err := s.repo.SaveAll(agencies); err != nil {
        return nil, err
    }

    return &newAgency, nil
}

func (s *AgencyService) UpdateAgency(id uint, name, address, phone, region string, employeeCount *uint32) (*model.Agency, error) {
    agencies, err := s.repo.LoadAll()
    if err != nil {
        return nil, err
    }

    index := s.repo.FindByIDIndex(agencies, id)
    if index == -1 {
        return nil, nil
    }

    if name != "" {
        agencies[index].Name = name
    }
    if address != "" {
        agencies[index].Address = address
    }
    if phone != "" {
        agencies[index].Phone = phone
    }
    if region != "" {
        agencies[index].Region = region
    }
    if employeeCount != nil {
        agencies[index].EmployeeCount = employeeCount
    }

    if err := s.repo.SaveAll(agencies); err != nil {
        return nil, err
    }

    return &agencies[index], nil
}

func (s *AgencyService) GetStatus() (int, uint32, map[string]int, error) {
    agencies, err := s.repo.LoadAll()
    if err != nil {
        return 0, 0, nil, err
    }

    totalAgencies := len(agencies)
    var totalEmployees uint32 = 0
    regionCount := make(map[string]int)

    for _, agency := range agencies {
        regionCount[agency.Region]++
        if agency.EmployeeCount != nil {
            totalEmployees += *agency.EmployeeCount
        }
    }

    return totalAgencies, totalEmployees, regionCount, nil
}