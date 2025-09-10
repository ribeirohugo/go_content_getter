package youtube

type Youtube struct{}

func NewYoutube() Youtube {
	return Youtube{}
}

type Video struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Uploader     string   `json:"uploader"`
	UploaderID   string   `json:"uploader_id"`
	UploaderURL  string   `json:"uploader_url"`
	UploadDate   string   `json:"upload_date"`
	Duration     int      `json:"duration"`
	ViewCount    int      `json:"view_count"`
	LikeCount    int      `json:"like_count,omitempty"`
	CommentCount int      `json:"comment_count,omitempty"`
	Description  string   `json:"description"`
	Categories   []string `json:"categories,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Formats      []Format `json:"formats"`
}

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
