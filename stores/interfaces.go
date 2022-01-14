package stores

type Customer interface {
	Get(name string) bool
	Create(name string, id int)
}
