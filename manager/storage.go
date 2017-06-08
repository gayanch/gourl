package manager

type Storage interface {
	ReadUrl(shorturl string) (string, error)
	SaveUrl(longurl string) (string, error)
}
