package domain

type DownloadRecord struct {
	UserName     string `json:"username"`
	DownloadPath string `json:"download_path"`
	CreateAt     int64  `json:"create_at"`
}
