package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"depeche/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroup_GetGroup(t *testing.T) {
	tests := []struct {
		name string
		link string

		expectedGroup *entities.Group
		expectedError error

		setupMock func(repo *mock_repository.MockGroup, link string)
	}{
		{
			name: "Success",
			link: "id1234",

			expectedGroup: &entities.Group{
				Link:      "id1234",
				Title:     "Group#1234",
				HideOwner: false,
				Management: []entities.GroupManagement{
					{
						Link: "id1",
						Role: "Owner",
					},
				},
			},
			expectedError: nil,
			setupMock: func(repo *mock_repository.MockGroup, link string) {
				group := &entities.Group{
					Link:      "id1234",
					Title:     "Group#1234",
					HideOwner: false,
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "Owner",
						},
					},
				}
				repo.EXPECT().GetGroupByLink(link).Return(group, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.link)

			_, err := groupService.GetGroup(test.link)
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

func TestGroup_GetUserGroupsByEmail(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(repo *mock_repository.MockGroup, email string,
			limit, offset int)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:      "id1234",
		//			Title:     "Group#1234",
		//			OwnerLink: "id1",
		//			HideOwner: false,
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id1",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//		{
		//			Link:      "id10",
		//			Title:     "Group#10",
		//			HideOwner: false,
		//			OwnerLink: "id100",
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id100",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//	},
		//	setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
		//		groups := []*entities.Group{
		//			{
		//				Link:      "id1234",
		//				Title:     "Group#1234",
		//				HideOwner: false,
		//				OwnerLink: "id1",
		//			},
		//			{
		//				Link:      "id10",
		//				Title:     "Group#10",
		//				OwnerLink: "id100",
		//				HideOwner: false,
		//			},
		//		}
		//		repo.EXPECT().GetUserGroupsByEmail(email, limit, offset).Return(groups, nil)
		//	},
		//},
		{
			name:   "Internal Error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				repo.EXPECT().GetUserGroupsByEmail(email, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.email, test.limit, test.offset)

			groups, err := groupService.GetUserGroupsByEmail(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.limit, len(groups))
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroup_GetPopularGroups(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(repo *mock_repository.MockGroup, email string,
			limit, offset int)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:      "id1234",
		//			Title:     "Group#1234",
		//			OwnerLink: "id1",
		//			HideOwner: false,
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id1",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//		{
		//			Link:      "id10",
		//			Title:     "Group#10",
		//			HideOwner: false,
		//			OwnerLink: "id100",
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id100",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//	},
		//	setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
		//		groups := []*entities.Group{
		//			{
		//				Link:      "id1234",
		//				Title:     "Group#1234",
		//				HideOwner: false,
		//				OwnerLink: "id1",
		//			},
		//			{
		//				Link:      "id10",
		//				Title:     "Group#10",
		//				OwnerLink: "id100",
		//				HideOwner: false,
		//			},
		//		}
		//		repo.EXPECT().GetPopularGroups(email, limit, offset).Return(groups, nil)
		//	},
		//},
		{
			name:   "Internal Error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				repo.EXPECT().GetPopularGroups(email, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.email, test.limit, test.offset)

			groups, err := groupService.GetPopularGroups(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.limit, len(groups))
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroup_GetManagedGroups(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		limit  int
		offset int

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(repo *mock_repository.MockGroup, email string,
			limit, offset int)
	}{
		//{
		//	name:   "Success",
		//	email:  "e.larkin@mail.ru",
		//	limit:  2,
		//	offset: 0,
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:      "id1234",
		//			Title:     "Group#1234",
		//			OwnerLink: "id1",
		//			HideOwner: false,
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id1",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//		{
		//			Link:      "id10",
		//			Title:     "Group#10",
		//			HideOwner: false,
		//			OwnerLink: "id1",
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id1",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//	},
		//	setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
		//		groups := []*entities.Group{
		//			{
		//				Link:      "id1234",
		//				Title:     "Group#1234",
		//				HideOwner: false,
		//				OwnerLink: "id1",
		//			},
		//			{
		//				Link:      "id10",
		//				Title:     "Group#10",
		//				OwnerLink: "id1",
		//				HideOwner: false,
		//			},
		//		}
		//		repo.EXPECT().GetManagedGroups(email, limit, offset).Return(groups, nil)
		//	},
		//},
		{
			name:   "Internal Error",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				repo.EXPECT().GetManagedGroups(email, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.email, test.limit, test.offset)

			groups, err := groupService.GetManagedGroups(test.email, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.limit, len(groups))
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroup_GetUserGroupsByLink(t *testing.T) {
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		expectedGroups []*entities.Group
		expectedError  error

		setupMock func(repo *mock_repository.MockGroup, email string,
			limit, offset int)
	}{
		//{
		//	name:   "Success",
		//	link:   "id200",
		//	limit:  2,
		//	offset: 0,
		//
		//	expectedGroups: []*entities.Group{
		//		{
		//			Link:      "id1234",
		//			Title:     "Group#1234",
		//			OwnerLink: "id1",
		//			HideOwner: false,
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id1",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//		{
		//			Link:      "id10",
		//			Title:     "Group#10",
		//			HideOwner: false,
		//			OwnerLink: "id100",
		//			Management: []entities.GroupManagement{
		//				{
		//					Link: "id100",
		//					Role: "owner",
		//				},
		//			},
		//		},
		//	},
		//	setupMock: func(repo *mock_repository.MockGroup, link string, limit, offset int) {
		//		groups := []*entities.Group{
		//			{
		//				Link:      "id1234",
		//				Title:     "Group#1234",
		//				HideOwner: false,
		//				OwnerLink: "id1",
		//			},
		//			{
		//				Link:      "id10",
		//				Title:     "Group#10",
		//				OwnerLink: "id100",
		//				HideOwner: false,
		//			},
		//		}
		//		repo.EXPECT().GetUserGroupsByLink(link, limit, offset).Return(groups, nil)
		//	},
		//},
		{
			name:   "Internal Error",
			link:   "id123",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, link string, limit, offset int) {
				repo.EXPECT().GetUserGroupsByLink(link, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.link, test.limit, test.offset)

			groups, err := groupService.GetUserGroupsByLink(test.link, test.limit, test.offset)
			if test.expectedError != nil {
				uerr, ok := err.(*apperror.ServerError)
				require.Equal(t, true, ok)
				require.Equal(t, test.expectedError, uerr.UserErr)
			} else {
				require.Equal(t, test.limit, len(groups))
				require.Equal(t, test.expectedGroups, groups)
				require.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGroup_CreateGroup(t *testing.T) {
	tests := []struct {
		name  string
		email string
		dto   *dto.Group

		expectedError error
		setupMock     func(repo *mock_repository.MockGroup, email string, group *dto.Group)
	}{
		{
			name:  "Success",
			email: "e.larkin@mail.ru",

			dto: &dto.Group{
				Title: "Group",
			},
			expectedError: nil,

			setupMock: func(repo *mock_repository.MockGroup, email string, group *dto.Group) {
				returned := &entities.Group{
					Title: "Group",
					Link:  "id1",
				}
				repo.EXPECT().CreateGroup(email, group).Return(returned, nil)
				repo.EXPECT().Subscribe(email, returned.Link).Return(nil)
			},
		},
		{
			name:  "Create error",
			email: "e.larkin@mail.ru",

			dto: &dto.Group{
				Title: "Group",
			},
			expectedError: apperror.InternalServerError,

			setupMock: func(repo *mock_repository.MockGroup, email string, group *dto.Group) {
				repo.EXPECT().CreateGroup(email, group).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.dto)
			err := groupService.CreateGroup(test.email, test.dto)
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

func TestGroup_AcceptRequest(t *testing.T) {
	tests := []struct {
		name         string
		managerEmail string
		userLink     string
		groupLink    string

		expectedError error

		setupMock func(repo *mock_repository.MockGroup,
			managerEmail, userLink, groupLink string)
	}{
		{
			name:         "Success",
			managerEmail: "e.larkin@mail.ru",
			userLink:     "id123",
			groupLink:    "id1234",

			setupMock: func(repo *mock_repository.MockGroup, managerEmail, userLink, groupLink string) {
				repo.EXPECT().IsOwner(managerEmail, groupLink).Return(true, nil)
				repo.EXPECT().AcceptRequest(userLink, groupLink).Return(nil)
			},
		},
		{
			name:         "Forbidden",
			managerEmail: "notamanager@mail.ru",
			userLink:     "id123",
			groupLink:    "id1234",

			expectedError: apperror.Forbidden,

			setupMock: func(repo *mock_repository.MockGroup, managerEmail, userLink, groupLink string) {
				repo.EXPECT().IsOwner(managerEmail, groupLink).Return(false, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.managerEmail, test.userLink, test.groupLink)
			err := groupService.AcceptRequest(test.managerEmail, test.userLink, test.groupLink)
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

func TestGroup_AcceptAllRequestsRequest(t *testing.T) {
	tests := []struct {
		name         string
		managerEmail string
		groupLink    string

		expectedError error

		setupMock func(repo *mock_repository.MockGroup,
			managerEmail, groupLink string)
	}{
		{
			name:         "Success",
			managerEmail: "e.larkin@mail.ru",
			groupLink:    "id1234",

			setupMock: func(repo *mock_repository.MockGroup, managerEmail, groupLink string) {
				repo.EXPECT().IsOwner(managerEmail, groupLink).Return(true, nil)
				repo.EXPECT().AcceptAllRequests(groupLink).Return(nil)
			},
		},
		{
			name:         "Forbidden",
			managerEmail: "notamanager@mail.ru",
			groupLink:    "id1234",

			expectedError: apperror.Forbidden,

			setupMock: func(repo *mock_repository.MockGroup, managerEmail, groupLink string) {
				repo.EXPECT().IsOwner(managerEmail, groupLink).Return(false, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.managerEmail, test.groupLink)
			err := groupService.AcceptAllRequests(test.managerEmail, test.groupLink)
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

func TestGroup_GetSubscribers(t *testing.T) {
	tests := []struct {
		name   string
		link   string
		limit  int
		offset int

		expectedError error
		expectedUsers []*entities.User

		setupMock func(repo *mock_repository.MockGroup, link string,
			limit, offset int)
	}{
		{
			name:   "Success",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					Link:      "id1",
					FirstName: "Egor",
					LastName:  "Larkin",
				},
				{
					Link:      "id2",
					FirstName: "Pavel",
					LastName:  "Repin",
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, link string, limit, offset int) {
				users := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
					{
						Link:      "id2",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
				}
				repo.EXPECT().GetSubscribers(link, limit, offset).Return(users, nil)
			},
		},
		{
			name:   "Error GetSubscribers",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, link string, limit, offset int) {
				repo.EXPECT().GetSubscribers(link, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.link, test.limit, test.offset)
			users, err := groupService.GetSubscribers(test.link, test.limit, test.offset)
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

func TestGroup_GetPendingRequests(t *testing.T) {
	tests := []struct {
		name   string
		email  string
		link   string
		limit  int
		offset int

		expectedError error
		expectedUsers []*entities.User

		setupMock func(repo *mock_repository.MockGroup, email, link string,
			limit, offset int)
	}{
		{
			name:   "Success",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedError: nil,
			expectedUsers: []*entities.User{
				{
					Link:      "id1",
					FirstName: "Egor",
					LastName:  "Larkin",
				},
				{
					Link:      "id2",
					FirstName: "Pavel",
					LastName:  "Repin",
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, email, link string, limit, offset int) {
				repo.EXPECT().IsOwner(email, link).Return(true, nil)
				users := []*entities.User{
					{
						Link:      "id1",
						FirstName: "Egor",
						LastName:  "Larkin",
					},
					{
						Link:      "id2",
						FirstName: "Pavel",
						LastName:  "Repin",
					},
				}
				repo.EXPECT().GetPendingRequests(link, limit, offset).Return(users, nil)
			},
		},
		{
			name:   "Error GetPendingRequests",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedError: apperror.InternalServerError,
			setupMock: func(repo *mock_repository.MockGroup, email, link string, limit, offset int) {
				repo.EXPECT().IsOwner(email, link).Return(true, nil)
				repo.EXPECT().GetPendingRequests(link, limit, offset).Return(nil, apperror.NewServerError(apperror.InternalServerError, nil))
			},
		},
		{
			name:   "Forbidden",
			link:   "id1234",
			limit:  2,
			offset: 0,

			expectedError: apperror.Forbidden,
			setupMock: func(repo *mock_repository.MockGroup, email, link string, limit, offset int) {
				repo.EXPECT().IsOwner(email, link).Return(false, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockGroup(ctrl)

			groupService := Group{
				repo: mockRepository,
			}

			test.setupMock(mockRepository, test.email, test.link, test.limit, test.offset)
			users, err := groupService.GetPendingRequests(test.email, test.link, test.limit, test.offset)
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
