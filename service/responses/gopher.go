package responses

type GetGopher struct {
	ID   int
	Name string
	Age  int
}

type CreateGopher struct{}
