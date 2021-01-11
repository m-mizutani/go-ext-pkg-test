package chain

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

const (
	userKey = "user"
)

func (x *Usecase) SaveUser(user string) error {
	if err := x.repo.AppendItem(userKey, user); err != nil {
		return err
	}
	return nil
}

func (x *Usecase) GetUsers() ([]string, error) {
	return x.repo.FetchItems(userKey)
}
