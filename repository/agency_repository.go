package repository

import (
    "agencycli/model"
    "bytes"
    "encoding/json"
    "os"
)

const datafile = "./data.json"

type AgencyRepository struct{}

func NewAgencyRepository() *AgencyRepository {
    return &AgencyRepository{}
}

func (r *AgencyRepository) LoadAll() ([]model.Agency, error) {
    var agencies []model.Agency
    data, err := os.ReadFile(datafile)
    if err != nil {
        if os.IsNotExist(err) {
            return []model.Agency{}, nil
        }
        return nil, err
    }
    if len(bytes.TrimSpace(data)) > 0 {
        if err := json.Unmarshal(data, &agencies); err != nil {
            return nil, err
        }
    }
    return agencies, nil
}

func (r *AgencyRepository) SaveAll(agencies []model.Agency) error {
    data, err := json.MarshalIndent(agencies, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(datafile, data, 0644)
}

func (r *AgencyRepository) FindByID(agencies []model.Agency, id uint) *model.Agency {
    for i := range agencies {
        if agencies[i].ID == id {
            return &agencies[i]
        }
    }
    return nil
}

func (r *AgencyRepository) FindByIDIndex(agencies []model.Agency, id uint) int {
    for i := range agencies {
        if agencies[i].ID == id {
            return i
        }
    }
    return -1
}

func (r *AgencyRepository) GetMaxID(agencies []model.Agency) uint {
    maxID := uint(0)
    for _, a := range agencies {
        if a.ID > maxID {
            maxID = a.ID
        }
    }
    return maxID
}