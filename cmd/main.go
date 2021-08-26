package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"path"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

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

func main() {
	rawUrl := "https://twitter.com/antsstyle/status/1430636115724230667?s=20"
	url, _ := url2.Parse(rawUrl)
	tweetId := path.Base(url.Path)

	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet(tweetId)
	if err != nil {
		panic(err)
	}

	photos := tweet.Photos

	for i, img := range photos {
		fmt.Println(img)
		err := downloadFile(img, fmt.Sprint(i)+".jpg")
		if err != nil {
			panic(err)
		}
	}
}
