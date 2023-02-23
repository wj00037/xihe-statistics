package domain

type DownloadRecord struct {
	UserName     Account `json:"username"`
	DownloadPath string  `json:"download_path"`
	CreateAt     int64   `json:"create_at"`
}

type CloneCountsRecord struct {
	Counts   int64 `json:"counts"`
	CreateAt int64 `json:"create_at"`
}
