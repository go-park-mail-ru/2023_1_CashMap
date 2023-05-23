package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities"
	"depeche/internal/entities/response"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"strconv"
)

type Sticker struct {
	service usecase.Sticker
}

func NewStickerHandler(service usecase.Sticker) *Sticker {
	return &Sticker{
		service: service,
	}
}

func (s *Sticker) GetStickerById(ctx *gin.Context) {
	sId := ctx.Query("sticker_id")

	stickerId, err := strconv.Atoi(sId)
	if err != nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	sticker, err := s.service.GetStickerByID(uint(stickerId))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetStickerByIdResponse{
		Body: response.GetStickerByIdBody{
			Sticker: sticker,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (s *Sticker) GetStickerPackInfo(ctx *gin.Context) {
	pId := ctx.Query("pack_id")

	packId, err := strconv.Atoi(pId)
	if err != nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	pack, err := s.service.GetStickerpackInfo(uint(packId))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetStickerPackInfoResponse{
		Body: response.GetStickerPackInfoBody{
			Stickerpack: pack,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (s *Sticker) GetStickerPack(ctx *gin.Context) {
	pId := ctx.Query("pack_id")

	packId, err := strconv.Atoi(pId)
	if err != nil {
		_ = ctx.Error(apperror.NewBadRequest())
		return
	}

	pack, err := s.service.GetStickerPack(uint(packId))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetStickerPackInfoResponse{
		Body: response.GetStickerPackInfoBody{
			Stickerpack: pack,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (s *Sticker) GetUserStickerPacks(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	packs, err := s.service.GetStickerPacksInfoByEmail(email, limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetUserStickerPacksResponse{
		Body: response.GetUserStickerPacksBody{
			Stickerpacks: packs,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)

}

func (s *Sticker) GetStickerPacksByAuthor(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	author := ctx.Query("author")
	var packs []*entities.StickerPack
	if author != "" {
		packs, err = s.service.GetStickerPacksInfoByAuthor(author, limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	} else {
		packs, err = s.service.GetDepechePacks(limit, offset)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}

	_response := response.GetUserStickerPacksResponse{
		Body: response.GetUserStickerPacksBody{
			Stickerpacks: packs,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (s *Sticker) GetNewStickerPacks(ctx *gin.Context) {
	limit, offset, err := utils.GetLimitOffset(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	packs, err := s.service.GetNewStickerPacksInfo(limit, offset)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.GetUserStickerPacksResponse{
		Body: response.GetUserStickerPacksBody{
			Stickerpacks: packs,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}

func (s *Sticker) AddStickerPack(ctx *gin.Context) {
	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	body, err := utils.GetBody[dto.AddStickerPack](ctx)

	err = s.service.AddStickerPack(email, body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
}

func (s *Sticker) UploadStickerPack(ctx *gin.Context) {
	inputDTO := new(response.UploadStickerPackRequest)
	if err := easyjson.UnmarshalFromReader(ctx.Request.Body, inputDTO); err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.BadRequest, errors.New("failed to parse struct")))
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	pack, err := s.service.UploadStickerPack(email, inputDTO.Body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	_response := response.UploadStickerPackResponse{
		Body: response.UploadStickerPackBody{
			Stickerpack: pack,
		},
	}

	responseJSON, err := _response.MarshalJSON()
	if err != nil {
		_ = ctx.Error(apperror.NewServerError(apperror.InternalServerError, err))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", responseJSON)
}
