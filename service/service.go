package service

// Request is a request to any service.
type Request interface{}

// Response is a response to a request.
type Response interface{}

// Creator is a boundary that creates resources.
type Creator interface {
	Create(Request) (Response, error)
}

// Finder is a boundary that finds resources.
type Finder interface {
	Find(Request) (Response, error)
}
