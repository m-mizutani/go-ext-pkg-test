package chain

type Repository interface {
	AppendItem(key, value string) error
	FetchItems(key string) ([]string, error)
}
