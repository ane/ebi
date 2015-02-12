package interactors

import (
	"reflect"

	"github.com/ane/ebi/service"
)

// Translator translates an object into a response.
type Translator interface {
	Translate(reflect.Type) (service.Response, error)
}
