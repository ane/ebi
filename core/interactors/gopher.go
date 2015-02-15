package interactors

import (
	"errors"

	"github.com/ane/ebi/core/entities"
	"github.com/ane/ebi/service/requests"
	"github.com/ane/ebi/service/responses"
)

type Gophers struct {
	burrow map[int]entities.Gopher
}

func NewGophers() *Gophers {
	return &Gophers{
		burrow: make(map[int]entities.Gopher),
	}
}

func (g Gophers) getFreeKey() int {
	key := 0
	for range g.burrow {
		key++
	}
	return key + 1
}

// Find finds a gopher from storage.
func (g Gophers) Find(req requests.FindGopher) (responses.FindGopher, error) {
	gopher, exists := g.burrow[req.ID]
	if !exists {
		return responses.FindGopher{}, errors.New("Not found.")
	}

	return gopher.ToFindGopher()
}

func (g Gophers) FindAll(req requests.FindGopher) ([]responses.FindGopher, error) {
	var resps []responses.FindGopher
	for _, gopher := range g.burrow {
		fg, err := gopher.ToFindGopher()
		if err != nil {
			return []responses.FindGopher{}, err
		}
		resps = append(resps, fg)
	}
	return resps, nil
}

// Create creates a gopher.
func (g Gophers) Create(req requests.CreateGopher) (responses.CreateGopher, error) {
	var gopher entities.Gopher
	if err := gopher.Validate(req); err != nil {
		return responses.CreateGopher{}, err
	}

	gopher.ID = g.getFreeKey()
	gopher.Name = req.Name
	gopher.Age = req.Age
	g.burrow[gopher.ID] = gopher

	return responses.CreateGopher{ID: gopher.ID}, nil
}
