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

// ToFindGopher produces a Gopher out of this entity.
func (g *Gopher) ToFindGopher() (responses.FindGopher, error) {
	return responses.FindGopher{ID: g.ID, Name: g.Name, Age: g.Age}, nil
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
	}
	return fmt.Errorf("I don't know how to validate %T", req)
}
