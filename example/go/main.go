package main

import (
	"fmt"

	"github.com/ane/ebi/api"
	"github.com/ane/ebi/core/interactors"
	"github.com/ane/ebi/service/requests"
)

func main() {
	api := api.NewAPI(interactors.NewGophers())

	cr, err := api.Gophers.Create(requests.CreateGopher{Name: "Munch", Age: 2})
	if err != nil {
		fmt.Println("Couldn't create:", err.Error())
		return
	}
	fmt.Printf("Created gopher at ID %d.\n", cr.ID)
	resp, err := api.Gophers.Find(requests.FindGopher{ID: cr.ID})
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Printf("Found: %+v\n", resp)
}
