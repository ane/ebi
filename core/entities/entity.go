package entities

import "github.com/ane/ebi/service"
import "reflect"

// Translator translates an object into a response.
type Translator interface {
	As(reflect.Type) (service.Response, error)
}

// Validator validates a request for an entity.
type Validator interface {
	Validate(service.Request) error
}

// Entity represents an entity. It needs a Validate method so that interactors can
// validate incoming requests and transformations. It also needs a Data method that
// can produce a response from this object for a specific type.
type Entity interface {
	Translator
	Validator
}
