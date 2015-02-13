package entities

import (
	"reflect"

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
func (g Gopher) As(as interface{}) (service.Response, error) {
	t := reflect.TypeOf(as)
	switch {
	case t == reflect.TypeOf(responses.FindGopher{}):
		return g.asFindGopher()
	default:
		return nil, fmt.Errorf("Unrecognized response model %T", as)
	}
}

func validate(req requests.CreateGopher) error {
	return nil
}

// Validate validates the incoming requests for this entity.
func (g Gopher) Validate(req *requests.CreateGopher) error {
	if req.Age < 0 {
		return errors.New("My age can't be negative!")
	}

	if req.Name == "" {
		return errors.New("I need a non-empty name.")
	}
	return nil
}
