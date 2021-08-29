package main

import (
	"go-twitter-downloader/pkg/app"
)

const configPath = "configs/config"

// @title go-twitter-downloader
// @version 0.1
// @description REST API for downloading images from twitter.
func main() {
	app.Run(configPath)
}
