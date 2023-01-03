package repository

type DownloadRecord interface {
	Get() (int64, error)
}
