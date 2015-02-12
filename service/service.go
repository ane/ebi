package service

type Request interface{}

type Response interface{}

// Creator is a boundary that creates resources.
type Creator interface {
	Create(Request) (Response, error)
}

// Finder is a boundary that finds resources.
type Finder interface {
	Find(Request) (Response, error)
}

