package dto

type FeedDTO struct {
	BatchSize    uint   `json:"batch_size"`
	LastPostDate string `json:"last_post_date"`
}
