package service

type Twitter interface {
	DownloadImages(tweetLink string, downloadPath string) ([]string, error)
}

type Services struct {
	Twitter Twitter
}

func NewServices() *Services {
	return &Services{
		Twitter: NewTwitterService(),
	}
}
