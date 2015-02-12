package entities

import "github.com/ane/ebi/service"
import "reflect"

// Entity represents an entity. It needs a Validate method so that interactors can
// validate incoming requests and transformations. It also needs a Data method that
// can produce a response from this object for a specific type.
type Entity interface {
	Validate(service.Request) error
	Data(reflect.Type) (service.Response, error)
}
