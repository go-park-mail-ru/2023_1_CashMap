package usecase

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

type Sticker interface {
	GetStickerByID(stickerId uint) (*entities.Sticker, error)
	GetStickerpackInfo(packId uint) (*entities.StickerPack, error)
	GetStickerPack(packId uint) (*entities.StickerPack, error)
	GetStickerPacksInfoByEmail(email string, limit, offset int) ([]*entities.StickerPack, error)
	GetStickerPacksInfoByAuthor(author string, limit, offset int) ([]*entities.StickerPack, error)
	GetDepechePacks(limit, offset int) ([]*entities.StickerPack, error)
	GetNewStickerPacksInfo(limit, offset int) ([]*entities.StickerPack, error)

	AddStickerPack(email string, dto *dto.AddStickerPack) error

	UploadStickerPack(authorEmail string, pack *dto.UploadStickerPack) (*entities.StickerPack, error)
}
