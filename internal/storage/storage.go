package storage

type URLShortener struct {
	urls map[string]string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}
