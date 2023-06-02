package postgres

import (
	"database/sql"
	"depeche/internal/entities"
	"depeche/pkg/apperror"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupRepository_GetGroupByLink(t *testing.T) {
	type dbBehaviour struct {
		error error
	}
	tests := []struct {
		name string
		link string

		expectedError error
		expectedGroup *entities.Group

		dbBehaviour
		setupMock func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour)
	}{
		//{
		//	name: "Success",
		//	link: "id1234",
		//
		//	expectedError: nil,
		//	expectedGroup: &entities.Group{
		//		Link:         "id1235",
		//		Title:        "Group",
		//		CreationDate: "2023-04-30T10:56:01Z",
		//		Privacy:      "open",
		//		MembersCount: 10,
		//		OwnerLink:    "id1",
		//	},
		//	dbBehaviour: dbBehaviour{
		//		data: sqlmock.NewRows([]string{
		//			"title", "link", "info",
		//			"privacy", "creation_date",
		//			"hide_author", "owner_link", "is_deleted",
		//			"subscribers", "avatar"}).
		//			AddRow("Group", "id12345", "", "open",
		//				"2023-04-30T10:56:01Z",
		//				false, "id1", false, 10, ""),
		//	},
		//	setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
		//		mock.ExpectQuery(GroupByLink).WithArgs(link).WillReturnRows(behaviour.data)
		//	},
		//},
		{
			name: "Not found",
			link: "id1234",

			expectedError: apperror.GroupNotFound,
			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupByLink).WithArgs(link).WillReturnError(behaviour.error)
			},
		},
		{
			name: "SQL error",
			link: "id1234",

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupByLink).WithArgs(link).WillReturnError(behaviour.error)
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
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			_, err = groupRepo.GetGroupByLink(test.link)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetUserGroupsByEmail(t *testing.T) {
	type dbBehaviour struct {
		error error
	}
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		dbBehaviour

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	dbBehaviour: dbBehaviour{
		//		data: sqlmock.NewRows([]string{
		//			"title", "link", "info",
		//			"privacy", "creation_date",
		//			"hide_author", "owner_link", "is_deleted",
		//			"subscribers", "avatar",
		//		}).
		//			AddRow(
		//				"Group#1", "id1", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id1", false,
		//				10, "").
		//			AddRow(
		//				"Group#2", "id2", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id2", false,
		//				10, ""),
		//	},
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:         "id1",
		//			Title:        "Group#1",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id1",
		//		},
		//		{
		//			Link:         "id2",
		//			Title:        "Group#2",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id2",
		//		},
		//	},
		//	setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
		//		mock.ExpectQuery(GroupsByUserEmail).WithArgs(email, limit, offset).WillReturnRows(behaviour.data)
		//	},
		//},
		{
			name:   "No rows",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupsByUserEmail).WithArgs(email, limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupsByUserEmail).WithArgs(email, limit, offset).WillReturnError(behaviour.error)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			t.Parallel()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			require.NoError(t, err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)

			test.setupMock(mock, test.email, test.limit, test.offset, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			groups, err := groupRepo.GetUserGroupsByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetUserGroupsByLink(t *testing.T) {
	type dbBehaviour struct {
		error error
	}
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		dbBehaviour

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour)
	}{
		//{
		//	name:   "Success",
		//	link:   "id123",
		//	limit:  2,
		//	offset: 0,
		//
		//	dbBehaviour: dbBehaviour{
		//		data: sqlmock.NewRows([]string{
		//			"title", "link", "info",
		//			"privacy", "creation_date",
		//			"hide_author", "owner_link", "is_deleted",
		//			"subscribers", "avatar",
		//		}).
		//			AddRow(
		//				"Group#1", "id1", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id1", false,
		//				10, "").
		//			AddRow(
		//				"Group#2", "id2", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id2", false,
		//				10, ""),
		//	},
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:         "id1",
		//			Title:        "Group#1",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id1",
		//		},
		//		{
		//			Link:         "id2",
		//			Title:        "Group#2",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id2",
		//		},
		//	},
		//	setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
		//		mock.ExpectQuery(GroupsByUserlink).WithArgs(link, limit, offset).WillReturnRows(behaviour.data)
		//	},
		//},
		{
			name:   "No rows",
			link:   "id123",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupsByUserlink).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			link:   "id123",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupsByUserlink).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
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

			test.setupMock(mock, test.link, test.limit, test.offset, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			groups, err := groupRepo.GetUserGroupsByLink(test.link, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetManagedGroups(t *testing.T) {
	type dbBehaviour struct {
		error error
	}
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		dbBehaviour

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	dbBehaviour: dbBehaviour{
		//		data: sqlmock.NewRows([]string{
		//			"title", "link", "info",
		//			"privacy", "creation_date",
		//			"hide_author", "owner_link", "is_deleted",
		//			"subscribers", "avatar",
		//		}).
		//			AddRow(
		//				"Group#1", "id1", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id1", false,
		//				10, "").
		//			AddRow(
		//				"Group#2", "id2", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id1", false,
		//				10, ""),
		//	},
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:         "id1",
		//			Title:        "Group#1",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id1",
		//		},
		//		{
		//			Link:         "id2",
		//			Title:        "Group#2",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id1",
		//		},
		//	},
		//	setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
		//		mock.ExpectQuery(GetManaged).WithArgs(email, limit, offset).WillReturnRows(behaviour.data)
		//	},
		//},
		{
			name:   "No rows",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GetManaged).WithArgs(email, limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GetManaged).WithArgs(email, limit, offset).WillReturnError(behaviour.error)
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

			test.setupMock(mock, test.email, test.limit, test.offset, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			groups, err := groupRepo.GetManagedGroups(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetPopularGroups(t *testing.T) {
	type dbBehaviour struct {
		error error
	}
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		dbBehaviour

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	dbBehaviour: dbBehaviour{
		//		data: sqlmock.NewRows([]string{
		//			"title", "link", "info",
		//			"privacy", "creation_date",
		//			"hide_author", "owner_link", "is_deleted",
		//			"subscribers", "avatar",
		//		}).
		//			AddRow(
		//				"Group#1", "id1", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id1", false,
		//				10, "").
		//			AddRow(
		//				"Group#2", "id2", "",
		//				"open", "2023-04-30T10:56:01Z",
		//				false, "id2", false,
		//				10, ""),
		//	},
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:         "id1",
		//			Title:        "Group#1",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id1",
		//		},
		//		{
		//			Link:         "id2",
		//			Title:        "Group#2",
		//			CreationDate: "2023-04-30T10:56:01Z",
		//			Privacy:      "open",
		//			MembersCount: 10,
		//			OwnerLink:    "id2",
		//		},
		//	},
		//	setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
		//		mock.ExpectQuery(GetGroups).WithArgs(limit, offset).WillReturnRows(behaviour.data)
		//	},
		//},
		{
			name:   "No rows",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GetGroups).WithArgs(limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, email string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GetGroups).WithArgs(limit, offset).WillReturnError(behaviour.error)
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

			test.setupMock(mock, test.email, test.limit, test.offset, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			groups, err := groupRepo.GetPopularGroups(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetSubscribers(t *testing.T) {
	type dbBehaviour struct {
		data  *sqlmock.Rows
		error error
	}
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		dbBehaviour

		expectedUsers []*entities.User
		expectedError error

		setupMock func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour)
	}{
		{
			name:   "Success",
			link:   "id123",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				data: sqlmock.
					NewRows([]string{
						"id", "link", "email",
						"first_name", "last_name",
						"sex", "bio", "status",
						"birthday", "last_active",
						"avatar"}).
					AddRow(
						1, "id1", "e.larkin@mail.ru",
						"Egor", "Larkin",
						"male", "Bio", "Status",
						"30.04.2002", "11.04.2023",
						"").
					AddRow(
						2, "id2", "other@mail.ru",
						"Pavel", "Repin",
						"male", "Bio", "Status",
						"17.07.2003", "11.04.2023",
						""),
			},

			expectedUsers: []*entities.User{
				{
					ID:         1,
					Link:       "id1",
					Email:      "e.larkin@mail.ru",
					FirstName:  "Egor",
					LastName:   "Larkin",
					Sex:        "male",
					Bio:        "Bio",
					Status:     "Status",
					BirthDay:   "30.04.2002",
					LastActive: "11.04.2023",
				},
				{
					ID:         2,
					Link:       "id2",
					Email:      "other@mail.ru",
					FirstName:  "Pavel",
					LastName:   "Repin",
					Sex:        "male",
					Bio:        "Bio",
					Status:     "Status",
					BirthDay:   "17.07.2003",
					LastActive: "11.04.2023",
				},
			},
			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupSubscribers).WithArgs(link, limit, offset).WillReturnRows(behaviour.data)
			},
		},
		{
			name:   "No rows",
			link:   "id123",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupSubscribers).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			link:   "id123",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(GroupSubscribers).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
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

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			test.setupMock(mock, test.link, test.limit, test.offset, test.dbBehaviour)

			users, err := groupRepo.GetSubscribers(test.link, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedUsers, users)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_GetPendingRequests(t *testing.T) {
	type dbBehaviour struct {
		data  *sqlmock.Rows
		error error
	}
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		dbBehaviour

		expectedUsers []*entities.User
		expectedError error

		setupMock func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour)
	}{
		{
			name:   "Success",
			link:   "id100",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				data: sqlmock.
					NewRows([]string{
						"id", "link", "email",
						"first_name", "last_name",
						"sex", "bio", "status",
						"birthday", "last_active",
						"avatar"}).
					AddRow(
						1, "id1", "e.larkin@mail.ru",
						"Egor", "Larkin",
						"male", "Bio", "Status",
						"30.04.2002", "11.04.2023",
						"").
					AddRow(
						2, "id2", "other@mail.ru",
						"Pavel", "Repin",
						"male", "Bio", "Status",
						"17.07.2003", "11.04.2023",
						""),
			},

			expectedUsers: []*entities.User{
				{
					ID:         1,
					Link:       "id1",
					Email:      "e.larkin@mail.ru",
					FirstName:  "Egor",
					LastName:   "Larkin",
					Sex:        "male",
					Bio:        "Bio",
					Status:     "Status",
					BirthDay:   "30.04.2002",
					LastActive: "11.04.2023",
				},
				{
					ID:         2,
					Link:       "id2",
					Email:      "other@mail.ru",
					FirstName:  "Pavel",
					LastName:   "Repin",
					Sex:        "male",
					Bio:        "Bio",
					Status:     "Status",
					BirthDay:   "17.07.2003",
					LastActive: "11.04.2023",
				},
			},
			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(PendingGroupRequests).WithArgs(link, limit, offset).WillReturnRows(behaviour.data)
			},
		},
		{
			name:   "No rows",
			link:   "id123",
			limit:  2,
			offset: 0,

			dbBehaviour: dbBehaviour{
				error: sql.ErrNoRows,
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(PendingGroupRequests).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
			},
		},
		{
			name:   "SQL error",
			link:   "id123",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			setupMock: func(mock sqlmock.Sqlmock, link string, limit, offset int, behaviour dbBehaviour) {
				mock.ExpectQuery(PendingGroupRequests).WithArgs(link, limit, offset).WillReturnError(behaviour.error)
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

			test.setupMock(mock, test.link, test.limit, test.offset, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			users, err := groupRepo.GetPendingRequests(test.link, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedUsers, users)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroupRepository_IsOwner(t *testing.T) {
	type dbBehaviour struct {
		data  *sqlmock.Rows
		error error
	}
	tests := []struct {
		name  string
		email string
		link  string

		dbBehaviour

		expectedRes   bool
		expectedError error

		setupMock func(mock sqlmock.Sqlmock, email, link string, behaviour dbBehaviour)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",
			link:  "id1234",
			dbBehaviour: dbBehaviour{
				data: sqlmock.
					NewRows([]string{"is_owner"}).
					AddRow(true),
			},
			expectedRes: true,
			setupMock: func(mock sqlmock.Sqlmock, email, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(IsOwner).WithArgs(email, link).WillReturnRows(behaviour.data)
			},
		},
		{
			name:  "SQL error",
			email: "e.larkin@mail.ru",
			link:  "id1234",
			dbBehaviour: dbBehaviour{
				error: errors.New("some sql error"),
			},

			expectedError: apperror.InternalServerError,
			setupMock: func(mock sqlmock.Sqlmock, email, link string, behaviour dbBehaviour) {
				mock.ExpectQuery(IsOwner).WithArgs(email, link).WillReturnError(behaviour.error)
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

			test.setupMock(mock, test.email, test.link, test.dbBehaviour)

			sqlxDB := sqlx.NewDb(db, "sqlmock")
			groupRepo := &GroupRepository{
				db: sqlxDB,
			}

			isOwner, err := groupRepo.IsOwner(test.email, test.link)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.expectedRes, isOwner)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}
