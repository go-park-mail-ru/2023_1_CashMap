package dto

type UploadStickerPack struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	Stickers    []string `json:"stickers"`
}

type AddStickerPack struct {
	StickerPackId uint `json:"stickerpack_id"`
}
