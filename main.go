package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {

	max := 5

	for i := 0; i < max; i++ {
		var youtubeVideo string
		var videoURL string

		fmt.Println("Copie a URL do vídeo a ser baixado: ")
		fmt.Scan(&youtubeVideo)

		videoID := strings.Split(youtubeVideo, "https://www.youtube.com/watch?v=")

		videoURL = strings.Join(videoID, "")

		client := youtube.Client{}

		video, err := client.GetVideo(videoURL)
		if err != nil {
			panic(err)
		}

		// only gets videos with audio

		formats := video.Formats.WithAudioChannels()

		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			panic(err)
		}

		file, err := os.Create(video.Title + ".mp4")
		if err != nil {
			panic(err)
		}

		defer file.Close()

		file2, err := io.Copy(file, stream)
		if err != nil {
			panic(err)
		}

		fmt.Println(file2)
	}

}
