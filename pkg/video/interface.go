package video

// GetterService defines the video service interface.
type GetterService interface {
	// GetVideoInfo returns video metadata for the provided URL.
	GetVideoInfo(url string) (Video, error)
	// DownloadYoutubeVideo downloads and muxes a YouTube video combining the selected video and audio formats.
	DownloadYoutubeVideo(url, videoFormat, audioFormat string) ([]byte, error)
	// DownloadVideo downloads a video selecting the best matching qualities for video and audio.
	DownloadVideo(url, videoQuality, audioQuality string) ([]byte, error)
	// DownloadAudio downloads only the audio stream in the requested format.
	DownloadAudio(url, audioFormat string) ([]byte, error)
	// GetTitle returns a sanitized title for the video at the provided URL.
	GetTitle(url string) (string, error)
}
