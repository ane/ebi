package responses

type FindGopher struct {
	ID   int
	Name string
	Age  int
}

type CreateGopher struct{}
