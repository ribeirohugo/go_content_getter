package video

import (
	"log"
	"os/exec"
)

const (
	ffmpegDependency = "ffmpeg"
	ytDlpDependency  = "yt-dlp"
)

// Getter is the video struct that holds video download methods and dependencies.
type Getter struct{}

// NewGetter is a Getter constructor.
func NewGetter() GetterService {
	dependencies := []string{ffmpegDependency, ytDlpDependency}

	invalid := Invalid{}

	for _, dep := range dependencies {
		if checkDependency(dep) {
			log.Printf("%s is installed ✅\n", dep)
			if dep == ffmpegDependency {
				invalid.ffmpeg = true
			}
			if dep == ytDlpDependency {
				invalid.ytDlp = true
			}
		} else {
			log.Printf("%s is NOT installed ❌\n", dep)
		}
	}

	if !invalid.ffmpeg || !invalid.ytDlp {
		return invalid
	}

	return Getter{}
}

func checkDependency(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// Video struct is a Youtube video stream data struct.
type Video struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Uploader    string   `json:"uploader"`
	UploaderID  string   `json:"uploader_id"`
	UploaderURL string   `json:"uploader_url"`
	UploadDate  string   `json:"upload_date"`
	Duration    int      `json:"duration"`
	ViewCount   int      `json:"view_count"`
	Description string   `json:"description"`
	Categories  []string `json:"categories,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Formats     []Format `json:"formats"`
}

// Format is a Video format data struct.
type Format struct {
	FormatID string  `json:"format_id"`
	Ext      string  `json:"ext"`
	Protocol string  `json:"protocol,omitempty"`
	URL      string  `json:"url,omitempty"`
	Acodec   string  `json:"acodec,omitempty"`
	Vcodec   string  `json:"vcodec,omitempty"`
	Height   int     `json:"height,omitempty"`
	Width    int     `json:"width,omitempty"`
	Filesize int64   `json:"filesize,omitempty"`
	Abr      float64 `json:"abr,omitempty"`
	Fps      float64 `json:"fps,omitempty"`
	Tbr      float64 `json:"tbr,omitempty"`
}
