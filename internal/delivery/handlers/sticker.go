package handlers

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/delivery/utils"
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"sticker": sticker,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpack": pack,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpack": pack,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpacks": packs,
		},
	})

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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpacks": packs,
		},
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpacks": packs,
		},
	})
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
	body, err := utils.GetBody[dto.UploadStickerPack](ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	email, err := utils.GetEmail(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	pack, err := s.service.UploadStickerPack(email, body)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": gin.H{
			"stickerpack": pack,
		},
	})
}
