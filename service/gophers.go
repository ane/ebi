package service

import (
	"github.com/ane/ebi/service/requests"
	"github.com/ane/ebi/service/responses"
)

// Gophers is a boundary that can do things with gophers.
type Gophers interface {
	Create(requests.CreateGopher) (responses.CreateGopher, error)
	Find(requests.FindGopher) (responses.FindGopher, error)
}
