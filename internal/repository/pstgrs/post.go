package pstgrs

import (
	"database/sql"
	"depeche/internal/entities"
	"depeche/internal/repository"
	"github.com/jmoiron/sqlx"
)

type PostStorage struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) repository.PostRepository {
	return &PostStorage{
		db: db,
	}
}

func (p PostStorage) SelectPostById(postId string) (*entities.Post, error) {
	post := &entities.Post{}
	err := p.db.Get(post, "SELECT * FROM Post WHERE id = ?", postId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p PostStorage) SelectPostsByCommunityLink(communityLink string, batchSize int, offset int) ([]*entities.Post, error) {
	rows, err := p.db.Queryx("SELECT * FROM Post WHERE id = (SELECT id FROM Community WHERE link = ?) ORDER BY creation_date DESC LIMIT ? OFFSET ?", communityLink, batchSize, offset)
	if err != nil {
		return nil, err
	}

	posts, err := getSliceFromRows[entities.Post](rows, batchSize)

	return posts, nil
}

func (p PostStorage) SelectPostsByUserLink(userLink string, batchSize int, offset int) ([]*entities.Post, error) {
	rows, err := p.db.Queryx("SELECT * FROM Post WHERE id = (SELECT id FROM UserProfile WHERE link = ?) ORDER BY creation_date DESC LIMIT ? OFFSET ?", userLink, batchSize, offset)
	if err != nil {
		return nil, err
	}

	posts, err := getSliceFromRows[entities.Post](rows, batchSize)

	return posts, nil
}

func getSliceFromRows[T any](rows *sqlx.Rows, size int) ([]*T, error) {
	posts := make([]*T, 0, size)
	it := 0
	for rows.Next() {
		err := rows.StructScan(posts[it])
		if err != nil {
			return nil, err
		}
		it++
	}

	return posts, nil
}
