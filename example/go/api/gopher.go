package api

import "github.com/ane/ebi/service"

type GopherAPI struct {
	Gophers service.Gophers
}

func NewAPI(gophers service.Gophers) *GopherAPI {
	return &GopherAPI{gophers}
}
