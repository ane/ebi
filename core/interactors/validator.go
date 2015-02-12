package interactors

import "github.com/ane/ebi/service"

// Validator validates a request for a type.
type Validator interface {
	Validate(service.Request) error
}
