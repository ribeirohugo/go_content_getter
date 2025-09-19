package video

import "errors"

const errorMessage = "ffmpeg and/or yt-dlp dependencies is/are not installed"

// Invalid is a Video getter service to be used when required dependencies are not installed.
type Invalid struct {
	ffmpeg bool
	ytDlp  bool
}

// GetVideoInfo returns video information for a given video URL.
func (i Invalid) GetVideoInfo(_ string) (Video, error) {
	return Video{}, errors.New(errorMessage)
}

// DownloadYoutubeVideo allows to get Youtube video stream.
func (i Invalid) DownloadYoutubeVideo(_, _, _ string) ([]byte, error) {
	return []byte{}, errors.New(errorMessage)
}

// DownloadVideo allows to get video stream.
func (i Invalid) DownloadVideo(_, _, _ string) ([]byte, error) {
	return []byte{}, errors.New(errorMessage)
}

// DownloadAudio allows to get audio stream.
func (i Invalid) DownloadAudio(_, _ string) ([]byte, error) {
	return []byte{}, errors.New(errorMessage)
}

// GetTitle allows to get title string for a given video URL.
func (i Invalid) GetTitle(_ string) (string, error) {
	return "", errors.New(errorMessage)
}
