package postgres

import (
	"database/sql"
	"database/sql/driver"
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	"depeche/pkg/apperror"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_GetUserById(t *testing.T) {
	type dbBehaviour struct {
		orderIDs []driver.Value
		data     *sqlmock.Rows
		error    error
	}
	tests := []struct {
		name          string
		id            uint
		expectedError error
		expectedUser  *entities.User

		dbBehaviour

		setupMock func(mock sqlmock.Sqlmock, id uint, behaviour dbBehaviour)
	}{
		{
			name:          "Success",
			id:            1,
			expectedError: nil,
			expectedUser:  &entities.User{ID: 1},

			dbBehaviour: dbBehaviour{
				orderIDs: []driver.Value{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				data: sqlmock.
					NewRows([]string{
						"id", "link", "email",
						"first_name", "last_name",
						"sex", "bio", "status",
						"birthday", "last_active",
						"avatar"}).AddRow(
					1, "id1", "e-larkin@mail.ru",
					"Egor", "Larkin",
					"male", "Bio", "Status",
					"30.04.02", "11.04.23",
					"url1"),
				error: nil,
			},
			setupMock: func(mock sqlmock.Sqlmock, id uint, behaviour dbBehaviour) {
				mock.ExpectQuery(UserById).WithArgs(id).WillReturnRows(behaviour.data)
			},
		},
		{
			name:          "NotFound",
			id:            1,
			expectedError: apperror.UserNotFound,
			expectedUser:  nil,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},
			setupMock: func(mock sqlmock.Sqlmock, id uint, behaviour dbBehaviour) {
				mock.ExpectQuery(UserById).WithArgs(id).WillReturnError(behaviour.error)
			},
		},
		{
			name:          "InternalError",
			id:            1,
			expectedError: apperror.InternalServerError,
			expectedUser:  nil,

			dbBehaviour: dbBehaviour{
				error: errors.New("some SQL error"),
			},
			setupMock: func(mock sqlmock.Sqlmock, id uint, behaviour dbBehaviour) {
				mock.ExpectQuery(UserById).WithArgs(id).WillReturnError(behaviour.error)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)

			test.setupMock(mock, test.id, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			userRepo := &UserRepository{
				DB: sqlxDB,
			}

			_, err = userRepo.GetUserById(test.id)
			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestUserRepository_GetUserByLink(t *testing.T) {
	type dbBehaviour struct {
		orderIDs []driver.Value
		data     *sqlmock.Rows
		error    error
	}
	tests := []struct {
		name          string
		link          string
		expectedError error
		expectedUser  *entities.User

		dbBehaviour

		setupMock func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour)
	}{
		{
			name:          "Success",
			link:          "id1",
			expectedError: nil,
			expectedUser:  &entities.User{ID: 1},

			dbBehaviour: dbBehaviour{
				orderIDs: []driver.Value{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				data: sqlmock.
					NewRows([]string{
						"id", "link", "email",
						"first_name", "last_name",
						"sex", "bio", "status",
						"birthday", "last_active",
						"avatar"}).AddRow(
					1, "id1", "e-larkin@mail.ru",
					"Egor", "Larkin",
					"male", "Bio", "Status",
					"30.04.02", "11.04.23",
					"url1"),
				error: nil,
			},
			setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(UserByLink).WithArgs(link).WillReturnRows(behaviour.data)
			},
		},
		{
			name:          "NotFound",
			link:          "id1",
			expectedError: apperror.UserNotFound,
			expectedUser:  nil,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},
			setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(UserByLink).WithArgs(link).WillReturnError(behaviour.error)
			},
		},
		{
			name:          "InternalError",
			link:          "id1",
			expectedError: apperror.InternalServerError,
			expectedUser:  nil,

			dbBehaviour: dbBehaviour{
				error: errors.New("some SQL error"),
			},
			setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(UserByLink).WithArgs(link).WillReturnError(behaviour.error)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)

			test.setupMock(mock, test.link, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			userRepo := &UserRepository{
				DB: sqlxDB,
			}

			_, err = userRepo.GetUserByLink(test.link)
			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestUserRepository_UpdateUser(t *testing.T) {
	type dbBehaviour struct {
		data  *sqlmock.Rows
		error error
	}
	tests := []struct {
		name  string
		email string
		user  *dto.EditProfile

		expectedUser  *entities.User
		expectedError error

		dbBehaviour

		setupMock func(mock sqlmock.Sqlmock, email string, profile *dto.EditProfile, behaviour dbBehaviour)
	}{}
	//	{
	//		name:  "success basic",
	//		email: "e-larkin@mail.ru",
	//		user:  testUtils.InitProfileBasic("Egor", "Larkin", "Bio"),
	//
	//		expectedError: nil,
	//		expectedUser:  nil,
	//
	//		dbBehaviour: dbBehaviour{
	//			error: nil,
	//			data:  sqlmock.NewRows([]string{}),
	//		},
	//
	//		setupMock: func(mock sqlmock.Sqlmock, email string, profile *dto.EditProfile, behaviour dbBehaviour) {
	//			mock.ExpectQuery("update userprofile set first_name = $1, last_name = $2, bio = $3 where email = $4").WithArgs(
	//				*profile.FirstName, *profile.LastName, *profile.Bio, email).WillReturnRows(behaviour.data)
	//		},
	//	},
	//}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)

			test.setupMock(mock, test.email, test.user, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			userRepo := &UserRepository{
				DB: sqlxDB,
			}
			_, err = userRepo.UpdateUser(test.email, test.user)
			require.Equal(t, test.expectedError, err)
		})
	}
}
