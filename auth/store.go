package auth

type Store interface {
	GetClient(id string) (AuthClient, bool)
}
