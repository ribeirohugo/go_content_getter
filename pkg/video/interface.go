package video

// GetterService defines the video service interface.
type GetterService interface {
	GetVideoInfo(url string) (Video, error)
	DownloadYoutubeVideo(url, videoFormat, audioFormat string) ([]byte, error)
	DownloadVideo(url, videoQuality, audioQuality string) ([]byte, error)
	DownloadAudio(url, audioFormat string) ([]byte, error)
	GetTitle(url string) (string, error)
}
