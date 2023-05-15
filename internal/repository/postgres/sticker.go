package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"depeche/pkg/apperror"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Sticker struct {
	db *sqlx.DB
}

func NewStickerRepository(db *sqlx.DB) repository.Sticker {
	return &Sticker{
		db: db,
	}
}

func (s *Sticker) GetStickerById(stickerID uint) (*entities.Sticker, error) {
	sticker := &entities.Sticker{}
	err := s.db.QueryRowx(GetStickerByID, stickerID).Scan(sticker)
	if err == sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.StickerNotFound, nil)
	}
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return sticker, nil
}

func (s *Sticker) GetStickerpackById(packID uint) (*entities.StickerPack, error) {
	pack := &entities.StickerPack{}
	err := s.db.QueryRowx(GetStickerpackByID, packID).Scan(pack)
	if err == sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.StickerpackNotFound, nil)
	}
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return pack, nil
}

func (s *Sticker) GetStickersByPackId(packID uint) ([]*entities.Sticker, error) {
	var stickers []*entities.Sticker
	rows, err := s.db.Queryx(GetStickersByPack, packID)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	for rows.Next() {
		sticker := &entities.Sticker{}
		err = rows.Scan(sticker)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		stickers = append(stickers, sticker)
	}

	return stickers, nil
}

func (s *Sticker) GetStickerPacksInfoByEmail(email string, limit, offset int) ([]*entities.StickerPack, error) {
	var packs []*entities.StickerPack
	rows, err := s.db.Queryx(GetStickerPacksByEmail, email, limit, offset)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	for rows.Next() {
		pack := &entities.StickerPack{}
		err = rows.Scan(pack)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func (s *Sticker) GetStickerPacksInfoByAuthor(author string, limit, offset int) ([]*entities.StickerPack, error) {
	var packs []*entities.StickerPack
	rows, err := s.db.Queryx(GetStickerPacksByAuthor, author, limit, offset)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	for rows.Next() {
		pack := &entities.StickerPack{}
		err = rows.Scan(pack)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func (s *Sticker) GetNewStickerPacksInfo(limit, offset int) ([]*entities.StickerPack, error) {
	packs := make([]*entities.StickerPack, 0, limit)
	rows, err := s.db.Queryx(GetNewStickerPacks, limit, offset)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	for rows.Next() {
		pack := &entities.StickerPack{}
		err = rows.Scan(pack)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func (s *Sticker) GetDepechePacks(limit, offset int) ([]*entities.StickerPack, error) {
	rows, err := s.db.Queryx(GetDepechePacks, limit, offset)
	if err != nil && err != sql.ErrNoRows {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}
	packs := make([]*entities.StickerPack, 0, limit)
	for rows.Next() {
		pack := &entities.StickerPack{}
		err = rows.Scan(pack)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

func (s *Sticker) AddStickerPack(email string, dto *dto.AddStickerPack) error {
	_, err := s.db.Exec(AddStickerPack, email, dto.StickerPackId)
	if err != nil {
		return apperror.NewServerError(apperror.InternalServerError, err)
	}
	return nil
}

func (s *Sticker) GetAuthorLink(authorEmail string) (string, error) {
	var link string
	err := s.db.QueryRowx(GetStickerpackAuthorLink, authorEmail).Scan(&link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", apperror.NewServerError(apperror.UserNotFound, nil)
		}
		return "", apperror.NewServerError(apperror.InternalServerError, err)
	}
	return link, nil
}

func (s *Sticker) UploadStickerpack(link string, pack *dto.UploadStickerPack, creationTime string) (*entities.StickerPack, error) {
	var id uint
	err := s.db.QueryRowx(UploadStickerpack,
		pack.Title, pack.Description,
		pack.Cover, link, creationTime).Scan(&id)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	var params []interface{}
	params = append(params, id)
	query := `insert into sticker (url, stickerpack_id) values `
	for i, sticker := range pack.Stickers {
		query += fmt.Sprintf("($%d, $1), ", i+2)
		params = append(params, sticker)
	}
	query, _ = strings.CutSuffix(query, ", ")
	query += ` returning id`

	rows, err := s.db.Queryx(query, params...)

	stickerIDs := make([]uint, 0, len(pack.Stickers))
	for rows.Next() {
		var id uint
		err := rows.Scan(&id)
		if err != nil {
			return nil, apperror.NewServerError(apperror.InternalServerError, err)
		}
		stickerIDs = append(stickerIDs, id)
	}

	stickers := make([]entities.Sticker, 0, len(pack.Stickers))
	for i := 0; i < len(pack.Stickers); i++ {
		stickers = append(stickers, entities.Sticker{
			ID:            stickerIDs[i],
			Url:           pack.Stickers[i],
			StickerpackID: id,
		})
	}

	result := &entities.StickerPack{
		ID:           id,
		Title:        pack.Title,
		Description:  pack.Description,
		Cover:        pack.Cover,
		Author:       link,
		CreationDate: creationTime,
		Stickers:     stickers,
	}
	return result, nil
}
