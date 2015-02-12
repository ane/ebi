package boundaries

import "github.com/ane/ebi/service"

// Creator is a boundary that creates resources.
type Creator interface {
	Create(service.Request) (service.Response, error)
}

// Finder is a boundary that finds resources.
type Finder interface {
	Find(service.Request) (service.Response, error)
}

// Gophers is a boundary that can do things with gophers.
type Gophers interface {
	Creator
	Finder
}
