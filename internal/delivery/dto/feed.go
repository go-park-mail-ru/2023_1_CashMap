package dto

//go:generate easyjson --all feed.go

type FeedDTO struct {
	BatchSize    uint   `form:"batch_size"`
	LastPostDate string `form:"last_post_date"`
}
