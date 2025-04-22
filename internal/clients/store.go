package clients

type Store interface {
	GetClient(id string) (AuthClient, bool)
}
