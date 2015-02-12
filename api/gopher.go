package api

import "github.com/ane/ebi/service/boundaries"

type GopherAPI struct {
	Gophers boundaries.Gophers
}

func NewAPI(gophers boundaries.Gophers) *GopherAPI {
	return &GopherAPI{gophers}
}
