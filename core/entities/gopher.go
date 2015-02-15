package entities

import (
	"errors"
	"fmt"

	"github.com/ane/ebi/service"
	"github.com/ane/ebi/service/requests"
	"github.com/ane/ebi/service/responses"
)

// Gopher represents a tiny rodent.
type Gopher struct {
	ID   int
	Name string
	Age  int
}

func (g Gopher) asFindGopher() (responses.FindGopher, error) {
	return responses.FindGopher{ID: g.ID, Name: g.Name, Age: g.Age}, nil
}

// As implements the Translator interface, converting this entity to some DTO.
func (g Gopher) Translate(req service.Response) (service.Response, error) {
	switch req.(type) {
	case requests.FindGopher:
		return g.asFindGopher()
	default:
		return nil, fmt.Errorf("Unrecognized response model %T", req)
	}
}

func (g Gopher) Validate(req service.Request) error {
	switch r := req.(type) {
	case requests.CreateGopher:
		if r.Age < 0 {
			return errors.New("My age can't be negative!")
		}
		if r.Name == "" {
			return errors.New("I need a non-empty name.")
		}
	default:
		return errors.New("I don't know how to validate that!")
	}
	return nil
}
