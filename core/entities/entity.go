package entities

import "github.com/ane/ebi/service"

// Validator validates a request transformation.
type Validator interface {
	Validate(service.Request) error
}

// Translator translates an object into a response DTO.
type Translator interface {
	As(service.Response) (service.Response, error)
}

// Entity represents an entity. It needs a Validate method so that interactors can
// validate incoming requests and transformations. It also needs a Data method that
// can produce a response from this object for a specific type.
type Entity interface {
	Translator
	Validator
}
