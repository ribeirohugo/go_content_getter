package video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"github.com/ribeirohugo/go_content_getter/internal/file"
)

// GetVideoInfo returns video information for a given video URL.
func (g Getter) GetVideoInfo(url string) (Video, error) {
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

// DownloadYoutubeVideo allows to get Youtube video stream.
func (g Getter) DownloadYoutubeVideo(url, videoFormat, audioFormat string) ([]byte, error) {
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

	return downloadBytes(cmd)
}

// DownloadVideo allows to get video stream.
func (g Getter) DownloadVideo(url, videoQuality, audioQuality string) ([]byte, error) {
	format := fmt.Sprintf(
		"bestvideo[ext=mp4][height<=%s]+bestaudio[abr<=%s]/best[ext=mp4][height<=%s]",
		videoQuality,
		audioQuality,
		videoQuality,
	)

	log.Printf("yt-dlp -f %s -o - %s", format, url)

	cmd := exec.Command("yt-dlp",
		"-f", format,
		"-o", "-",
		url,
	)

	return downloadBytes(cmd)
}

// DownloadAudio allows to get audio stream.
func (g Getter) DownloadAudio(url, audioFormat string) ([]byte, error) {
	format := fmt.Sprintf(
		"bestaudio[abr<=%s]/best[abr<=%s]",
		audioFormat,
		audioFormat,
	)

	cmd := exec.Command("yt-dlp",
		"-f", format,
		"-x", "--audio-format", "mp3",
		"-o", "-",
		url,
	)

	return downloadBytes(cmd)
}

func downloadBytes(cmd *exec.Cmd) ([]byte, error) {
	var out bytes.Buffer    // will hold the video bytes
	var stderr bytes.Buffer // will hold yt-dlp logs

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return []byte{}, err
	}

	log.Println(stderr.String())

	return out.Bytes(), nil
}

// GetTitle allows to get title string for a given video URL.
func (g Getter) GetTitle(url string) (string, error) {
	cmd := exec.Command("yt-dlp", "--get-title", url)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Println("yt-dlp error:", err)
		log.Println(stderr.String())
		return "", fmt.Errorf("yt-dlp error: %v", err)
	}

	reader := transform.NewReader(bytes.NewReader(out.Bytes()), charmap.ISO8859_1.NewDecoder())
	utf8Output, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	sanitized := file.SanitizeFilename(string(utf8Output))

	return sanitized, nil
}
