package entities

import "github.com/ane/ebi/service"

// Entity represents an entity. It needs a Validate method so that interactors can
// validate incoming requests and transformations.
type Entity interface {
	Validate(service.Request) error
}
