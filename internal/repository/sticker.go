package repository

import (
	"depeche/internal/entities"
)

type Sticker interface {
	GetStickersByPackId(packID uint) ([]*entities.Sticker, error)
	GetStickerPacksInfoByAuthor(author string, limit, offset int) ([]*entities.StickerPack, error)
	GetDepechePacks(limit, offset int) ([]*entities.StickerPack, error)
	// GetStickerById(stickerID uint) (*entities.Sticker, error)
	// GetStickerpackById(packID uint) (*entities.StickerPack, error)
	// GetStickerPacksInfoByEmail(email string, limit, offset int) ([]*entities.StickerPack, error)
	// GetNewStickerPacksInfo(limit, offset int) ([]*entities.StickerPack, error)
	// AddStickerPack(email string, dto *dto.AddStickerPack) error

	// GetAuthorLink(authorEmail string) (string, error)
	// UploadStickerpack(link string, pack *dto.UploadStickerPack, creationTime string) (*entities.StickerPack, error)
}
