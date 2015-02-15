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
	resp, err := api.Gophers.Find(requests.FindGopher{ID: 0})
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		r, ok := resp.(responses.FindGopher)
		if !ok {
			fmt.Printf("Not responses.GetGopher, but %T.\n", r)
		} else {

			fmt.Printf("%#v\n", r)
		}
	}
}
