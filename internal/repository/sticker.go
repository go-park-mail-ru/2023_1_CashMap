package repository

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Sticker interface {
	GetStickerById(stickerID uint) (*entities.Sticker, error)
	GetStickerpackById(packID uint) (*entities.StickerPack, error)
	GetStickersByPackId(packID uint) ([]*entities.Sticker, error)
	GetStickerPacksInfoByEmail(email string, limit, offset int) ([]*entities.StickerPack, error)
	GetStickerPacksInfoByAuthor(author string, limit, offset int) ([]*entities.StickerPack, error)
	GetNewStickerPacksInfo(limit, offset int) ([]*entities.StickerPack, error)
	GetDepechePacks(limit, offset int) ([]*entities.StickerPack, error)
	AddStickerPack(email string, dto *dto.AddStickerPack) error

	GetAuthorLink(authorEmail string) (string, error)
	UploadStickerpack(link string, pack *dto.UploadStickerPack, creationTime string) (*entities.StickerPack, error)
}
