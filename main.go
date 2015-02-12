package main

import (
	"fmt"

	"github.com/ane/ebi/api"
	"github.com/ane/ebi/core/interactors"
	"github.com/ane/ebi/service/requests"
	"github.com/ane/ebi/service/responses"
)

func main() {
	api := api.NewAPI(interactors.NewGophers())

	api.Gophers.Create(requests.CreateGopher{Name: "Munch", Age: 2})
	resp, err := api.Gophers.Find(requests.GetGopher{ID: 0})
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	r, ok := resp.(responses.GetGopher)
	if !ok {
		fmt.Printf("Not responses.GetGopher, but %T.\n", r)
	}

	fmt.Printf("%#v\n", r)
}
