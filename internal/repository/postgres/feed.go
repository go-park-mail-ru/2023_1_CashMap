package postgres

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
	utildb "depeche/internal/repository/utils"
	"depeche/pkg/apperror"
	"github.com/jmoiron/sqlx"
)

type FeedStorage struct {
	db *sqlx.DB
}

func NewFeedStorage(db *sqlx.DB) repository.FeedRepository {
	return &FeedStorage{
		db: db,
	}
}

func (storage FeedStorage) GetFriendsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error) {
	var posts []*entities.Post
	rows, err := storage.db.Queryx(FriendsPostsQuery, email, feedDTO.BatchSize, feedDTO.LastPostDate)
	defer utildb.CloseRows(rows)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	posts, err = getSliceFromRows[entities.Post](rows, feedDTO.BatchSize)
	if err != nil {
		return nil, apperror.NewServerError(apperror.InternalServerError, err)
	}

	return posts, nil
}

func (storage FeedStorage) GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error) {
	// TODO: к 3 рк :)
	return nil, nil
}
