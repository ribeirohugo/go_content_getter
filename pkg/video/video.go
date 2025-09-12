package video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func (y Getter) GetVideoInfo(url string) (Video, error) {
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

func (y Getter) DownloadYoutubeVideo(url, videoFormat, audioFormat string) ([]byte, error) {
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

func (y Getter) DownloadVideo(url, videoQuality, audioQuality string) ([]byte, error) {
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

func (y Getter) DownloadAudio(url, audioFormat string) ([]byte, error) {
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

func (y Getter) GetTitle(url string) (string, error) {
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

	sanitizedFilename := sanitizeFilename(out.String())

	return sanitizedFilename, nil
}

func sanitizeFilename(name string) string {
	// Define invalid characters for most filesystems: \ / : * ? " < > |
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	safe := re.ReplaceAllString(name, "_")

	// Optional: trim spaces at start/end
	safe = strings.TrimSpace(safe)

	safe, _, _ = transform.String(charmap.ISO8859_1.NewDecoder(), safe)

	return safe
}
