package requests

type FindGopher struct {
	ID int
}

type CreateGopher struct {
	Name string
	Age  int
}
