package chain

type item struct {
	PKey  string `dynamo:"pk"`
	SKey  string `dynamo:"sk"`
	Value string `dynamo:"value"`
}

type Repository interface {
	SaveItem(key, value string) error
	GetItems(key string) ([]string, error)
}
