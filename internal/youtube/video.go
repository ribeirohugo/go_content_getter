package youtube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func (y Youtube) GetVideoInfo(url string) (Video, error) {
	cmd := exec.Command("yt-dlp", "-j", url)

	data, err := cmd.Output()
	if err != nil {
		return Video{}, fmt.Errorf("yt-dlp failed: %v", err)
	}

	var video Video
	if err := json.Unmarshal(data, &video); err != nil {
		return Video{}, fmt.Errorf("data unmarshall failed: %v", err)
	}

	return video, nil
}

func (y Youtube) DownloadVideo(url, videoFormat, audioFormat string) ([]byte, error) {
	format := fmt.Sprintf("%s+%s", videoFormat, audioFormat)
	if videoFormat == "" {
		format = audioFormat
	}
	if audioFormat == "" {
		format = videoFormat
	}

	log.Printf("yt-dlp -f %s -o - %s", format, url)

	cmd := exec.Command("yt-dlp",
		"-f", format,
		"-o", "-",
		url,
	)

	var out bytes.Buffer    // will hold the video bytes
	var stderr bytes.Buffer // will hold yt-dlp logs

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return []byte{}, err
	}

	return out.Bytes(), nil
}
