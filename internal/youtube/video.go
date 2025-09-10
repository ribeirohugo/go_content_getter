package youtube

import (
	"encoding/json"
	"fmt"
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
	cmd := exec.Command("yt-dlp",
		"-f", videoFormat+"+"+audioFormat,
		"-o -",
		url,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, fmt.Errorf("yt-dlp failed: %v", err)
	}

	return output, nil
}
