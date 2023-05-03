package postgres

import (
	"database/sql"
	"depeche/internal/delivery/dto"
	"depeche/internal/utils"
	"depeche/internal/utils/testUtils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostStorage_CreatePost(t *testing.T) {
	type dbBehaviour struct {
		posts *sqlmock.Rows
		error error
	}
	tests := []struct {
		name  string
		email string
		post  *dto.PostCreate

		expectedId    uint
		expectedError error

		dbBehaviour

		setupMock func(mock sqlmock.Sqlmock, email string, post *dto.PostCreate, behaviour dbBehaviour, id uint)
	}{
		{
			name:  "Community post",
			email: "e-larkin@mail.ru",
			post:  testUtils.InitPostCreateDtoForCommunity("testLink"),

			expectedError: nil,
			expectedId:    3,

			dbBehaviour: dbBehaviour{
				posts: sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2),
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, post *dto.PostCreate, behaviour dbBehaviour, id uint) {
				mock.ExpectBegin()

				mock.ExpectQuery("SELECT id FROM groups WHERE link = $1").WithArgs(*post.CommunityLink).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				currentTime := utils.CurrentTimeString()
				query := mock.ExpectPrepare(CreatePostQuery)
				query.ExpectQuery().WithArgs(1, email, nil, false, "", currentTime, currentTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))

				mock.ExpectCommit()
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)

			test.setupMock(mock, test.email, test.post, test.dbBehaviour, test.expectedId)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			userRepo := &PostStorage{
				db: sqlxDB,
			}

			postID, err := userRepo.CreatePost(test.email, test.post)
			assert.NoError(t, err)

			require.Equal(t, postID, test.expectedId)
		})
	}
}
