package response

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
)

//go:generate easyjson --all stickers.go

type GetStickerByIdResponse struct {
	Body GetStickerByIdBody `json:"body"`
}

type GetStickerByIdBody struct {
	Sticker *entities.Sticker `json:"sticker"`
}

type GetStickerPackInfoResponse struct {
	Body GetStickerPackInfoBody `json:"body"`
}

type GetStickerPackInfoBody struct {
	Stickerpack *entities.StickerPack `json:"stickerpack"`
}

type GetStickerPackResponse struct {
	Body GetStickerPackBody `json:"body"`
}

type GetStickerPackBody struct {
	Stickerpack *entities.StickerPack `json:"stickerpack"`
}

type GetUserStickerPacksResponse struct {
	Body GetUserStickerPacksBody `json:"body"`
}

type GetUserStickerPacksBody struct {
	Stickerpacks []*entities.StickerPack `json:"stickerpacks"`
}

type GetStickerPacksByAuthorResponse struct {
	Body GetStickerPacksByAuthorBody `json:"body"`
}

type GetStickerPacksByAuthorBody struct {
	Stickerpacks []*entities.StickerPack `json:"stickerpacks"`
}

type GetNewStickerPacksResponse struct {
	Body GetNewStickerPacksBody `json:"body"`
}

type GetNewStickerPacksBody struct {
	Stickerpacks []*entities.StickerPack `json:"stickerpacks"`
}

type AddStickerPackRequest struct {
	Body *dto.AddStickerPack `json:"body"`
}

type UploadStickerPackRequest struct {
	Body *dto.UploadStickerPack `json:"body"`
}

type UploadStickerPackResponse struct {
	Body UploadStickerPackBody `json:"body"`
}

type UploadStickerPackBody struct {
	Stickerpack *entities.StickerPack `json:"stickerpack"`
}
