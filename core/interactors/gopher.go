package interactors

import (
	"errors"

	"github.com/ane/ebi/core/entities"
	"github.com/ane/ebi/service"
	"github.com/ane/ebi/service/requests"
	"github.com/ane/ebi/service/responses"
)

// Gophers is an interactor that implements the boundaries.Gopher boundary.
type Gophers struct {
	Burrow map[int]entities.Gopher
}

func NewGophers() *Gophers {
	return &Gophers{
		Burrow: make(map[int]entities.Gopher),
	}
}

func (g Gophers) getFreeKey() int {
	key := 0
	for range g.Burrow {
		key++
	}
	return key
}

// Find finds a gopher from storage.
func (g Gophers) Find(req service.Request) (service.Response, error) {
	r, is := req.(requests.GetGopher)
	if !is {
		return nil, errors.New("Invalid request DTO given.")
	}

	gopher, exists := g.Burrow[r.ID]
	if !exists {
		return nil, errors.New("Not found.")
	}

	return gopher.Data(responses.GetGopher{})
}

// Create creates a gopher.
func (g Gophers) Create(req service.Request) (service.Response, error) {
	r, is := req.(requests.CreateGopher)
	if !is {
		return nil, errors.New("Invalid request DTO given.")
	}

	var gopher entities.Gopher
	if err := gopher.Validate(&r); err != nil {
		return nil, err
	}

	gopher.ID = g.getFreeKey()
	gopher.Name = r.Name
	gopher.Age = r.Age

	g.Burrow[gopher.ID] = gopher

	return responses.CreateGopher{}, nil
}
