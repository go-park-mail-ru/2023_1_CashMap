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
		{
			name:   "Success",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedGroups: []*entities.Group{
				{
					Link:      "id1234",
					Title:     "Group#1234",
					OwnerLink: "id1",
					HideOwner: false,
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "owner",
						},
					},
				},
				{
					Link:      "id10",
					Title:     "Group#10",
					HideOwner: false,
					OwnerLink: "id100",
					Management: []entities.GroupManagement{
						{
							Link: "id100",
							Role: "owner",
						},
					},
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:      "id1234",
						Title:     "Group#1234",
						HideOwner: false,
						OwnerLink: "id1",
					},
					{
						Link:      "id10",
						Title:     "Group#10",
						OwnerLink: "id100",
						HideOwner: false,
					},
				}
				repo.EXPECT().GetUserGroupsByEmail(email, limit, offset).Return(groups, nil)
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
		{
			name:   "Success",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedGroups: []*entities.Group{
				{
					Link:      "id1234",
					Title:     "Group#1234",
					OwnerLink: "id1",
					HideOwner: false,
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "owner",
						},
					},
				},
				{
					Link:      "id10",
					Title:     "Group#10",
					HideOwner: false,
					OwnerLink: "id100",
					Management: []entities.GroupManagement{
						{
							Link: "id100",
							Role: "owner",
						},
					},
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:      "id1234",
						Title:     "Group#1234",
						HideOwner: false,
						OwnerLink: "id1",
					},
					{
						Link:      "id10",
						Title:     "Group#10",
						OwnerLink: "id100",
						HideOwner: false,
					},
				}
				repo.EXPECT().GetPopularGroups(email, limit, offset).Return(groups, nil)
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
		{
			name:   "Success",
			email:  "e.larkin@mail.ru",
			limit:  2,
			offset: 0,

			expectedGroups: []*entities.Group{
				{
					Link:      "id1234",
					Title:     "Group#1234",
					OwnerLink: "id1",
					HideOwner: false,
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "owner",
						},
					},
				},
				{
					Link:      "id10",
					Title:     "Group#10",
					HideOwner: false,
					OwnerLink: "id1",
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "owner",
						},
					},
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, email string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:      "id1234",
						Title:     "Group#1234",
						HideOwner: false,
						OwnerLink: "id1",
					},
					{
						Link:      "id10",
						Title:     "Group#10",
						OwnerLink: "id1",
						HideOwner: false,
					},
				}
				repo.EXPECT().GetManagedGroups(email, limit, offset).Return(groups, nil)
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
		{
			name:   "Success",
			link:   "id200",
			limit:  2,
			offset: 0,

			expectedGroups: []*entities.Group{
				{
					Link:      "id1234",
					Title:     "Group#1234",
					OwnerLink: "id1",
					HideOwner: false,
					Management: []entities.GroupManagement{
						{
							Link: "id1",
							Role: "owner",
						},
					},
				},
				{
					Link:      "id10",
					Title:     "Group#10",
					HideOwner: false,
					OwnerLink: "id100",
					Management: []entities.GroupManagement{
						{
							Link: "id100",
							Role: "owner",
						},
					},
				},
			},
			setupMock: func(repo *mock_repository.MockGroup, link string, limit, offset int) {
				groups := []*entities.Group{
					{
						Link:      "id1234",
						Title:     "Group#1234",
						HideOwner: false,
						OwnerLink: "id1",
					},
					{
						Link:      "id10",
						Title:     "Group#10",
						OwnerLink: "id100",
						HideOwner: false,
					},
				}
				repo.EXPECT().GetUserGroupsByLink(link, limit, offset).Return(groups, nil)
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
				repo.EXPECT().CreateGroup(email, group).Return(&entities.Group{}, nil)
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
