package simplex

import (
	"context"
	"errors"
)

// Interface
type TerminologyService interface {
	TerminologySimplified(ctx context.Context, id int) (SimpleTerms, error)
}

// Constructor
func NewTerminologyService() *terminologyService {
	return &terminologyService{}
}

// Struct that implements the struct
type terminologyService struct{}

// TerminologySimplified: This methods iterate over the database and match the ID. If it finds the ID it will fill the struct and returns it.
// This service endpoint is called from the curat to elaborate the ID.
func (t terminologyService) TerminologySimplified(ctx context.Context, id int) (SimpleTerms, error) {
	var response SimpleTerms
	var found bool
	found = false
	for _, simple := range SimpleData.Glossary {
		if simple.ID == id {
			response = SimpleTerms{
				Name:        simple.Name,
				FullForm:    simple.FullForm,
				Description: simple.Description,
			}
			found = true
		}
	}
	if !found {
		return response, errors.New("requested id not found")
	}
	return response, nil
}
