package requests

type GetGopher struct {
	ID int
}

type CreateGopher struct {
	Name string
	Age  int
}
