package postgres

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/internal/repository"
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
	friendsQuery := `SELECT * FROM Post WHERE owner_id IN (select u.id from
    friendrequests f1
        join friendrequests f2 on
                f1.subscribed = f2.subscriber and
                f2.subscribed = f1.subscriber
        join userprofile u on
            f1.subscribed = u.id
        where
        f1.subscriber = (SELECT id FROM UserProfile WHERE email = $1))
        order by creation_date desc
        LIMIT $2 OFFSET $3`

	var posts []*entities.Post
	err := storage.db.Select(posts, friendsQuery, email, feedDTO.BatchSize, feedDTO.LastPostDate)
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (storage FeedStorage) GetGroupsPosts(email string, feedDTO *dto.FeedDTO) ([]*entities.Post, error) {
	// TODO: к 3 рк :)
	return nil, nil
}
