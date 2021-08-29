package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"path"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type TwitterService struct{}

func NewTwitterService() *TwitterService {
	return &TwitterService{}
}

func (t *TwitterService) DownloadImages(tweetLink string, downloadPath string) ([]string, error) {
	url, err := url2.Parse(tweetLink)
	if err != nil {
		return nil, err
	}
	tweetId := path.Base(url.Path)

	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet(tweetId)
	if err != nil {
		return nil, err
	}
	photos := tweet.Photos

	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		err := os.Mkdir(downloadPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	var filenames []string
	for i, img := range photos {
		fileName := fmt.Sprintf("%s_%s_%d.%s", tweet.Username, tweet.ID, i, photos[0][len(photos[0])-3:])
		err := downloadFile(img, downloadPath+fileName)
		if err != nil {
			return nil, err
		}
		log.Printf("Downloaded image: %s\n", downloadPath+fileName)
		filenames = append(filenames, fileName)
	}
	return filenames, nil
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
