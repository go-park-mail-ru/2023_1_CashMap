package service

import (
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/internal/usecase"
)

const MAX_STICKERPACK_SIZE = 20

type Sticker struct {
	repo repository.Sticker
}

func NewStickerService(repo repository.Sticker) usecase.Sticker {
	return &Sticker{
		repo: repo,
	}
}

func (s *Sticker) GetStickerPacksInfoByAuthor(author string, limit, offset int) ([]*entities.StickerPack, error) {
	return s.repo.GetStickerPacksInfoByAuthor(author, limit, offset)
}

func (s *Sticker) GetDepechePacks(limit, offset int) ([]*entities.StickerPack, error) {
	packs, err := s.repo.GetDepechePacks(limit, offset)
	if err != nil {
		return nil, err
	}
	for _, pack := range packs {
		stickerPtrs, err := s.repo.GetStickersByPackId(pack.ID)
		if err != nil {
			return nil, err
		}
		stickers := make([]entities.Sticker, len(stickerPtrs))
		for i, sticker := range stickerPtrs {
			stickers[i] = *sticker
		}
		pack.Stickers = stickers
	}
	return packs, nil
}

//func (s *Sticker) GetStickerByID(stickerId uint) (*entities.Sticker, error) {
//	return s.repo.GetStickerById(stickerId)
//}
//
//func (s *Sticker) GetStickerpackInfo(packId uint) (*entities.StickerPack, error) {
//	return s.repo.GetStickerpackById(packId)
//}
//
//func (s *Sticker) GetStickerPack(packId uint) (*entities.StickerPack, error) {
//	pack, err := s.repo.GetStickerpackById(packId)
//	if err != nil {
//		return nil, err
//	}
//
//	stickerPtrs, err := s.repo.GetStickersByPackId(packId)
//	if err != nil {
//		return nil, err
//	}
//
//	stickers := make([]entities.Sticker, len(stickerPtrs))
//	for i, sticker := range stickerPtrs {
//		stickers[i] = *sticker
//	}
//
//	pack.Stickers = stickers
//	return pack, nil
//}
//
//func (s *Sticker) GetStickerPacksInfoByEmail(email string, limit, offset int) ([]*entities.StickerPack, error) {
//	return s.repo.GetStickerPacksInfoByEmail(email, limit, offset)
//}
//
//func (s *Sticker) GetNewStickerPacksInfo(limit, offset int) ([]*entities.StickerPack, error) {
//	return s.repo.GetNewStickerPacksInfo(limit, offset)
//}
//
//func (s *Sticker) AddStickerPack(email string, dto *dto.AddStickerPack) error {
//	return s.repo.AddStickerPack(email, dto)
//}
//
//func (s *Sticker) UploadStickerPack(authorEmail string, pack *dto.UploadStickerPack) (*entities.StickerPack, error) {
//	if len(pack.Stickers) > MAX_STICKERPACK_SIZE {
//		return nil, apperror.NewServerError(apperror.TooManyStickers, nil)
//	}
//
//	link, err := s.repo.GetAuthorLink(authorEmail)
//	if err != nil {
//		return nil, err
//	}
//
//	result, err := s.repo.UploadStickerpack(link, pack, utils.CurrentTimeString())
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
