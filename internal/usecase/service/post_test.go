package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPostService_GetPostById(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.PostGetByID

		ExpectedOutput *entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository, email string, postID uint)
	}{
		{
			Name:  "Existing post with access",
			Email: "test@gmail.com",
			Dto: &dto.PostGetByID{
				PostID: 1,
			},

			ExpectedOutput: &entities.Post{
				ID: 1,
			},
			ExpectedErr: nil,

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, postID uint) {
				repo.EXPECT().CheckReadAccess(email).Return(true, nil)
				repo.EXPECT().GetPostSenderInfo(postID).Return(nil, nil, nil)
				repo.EXPECT().SelectPostById(postID, email).Return(&entities.Post{
					ID: 1,
				}, nil)
			},
		},

		{

			Name:  "Existing post with no access",
			Email: "test@gmail.com",
			Dto: &dto.PostGetByID{
				PostID: 1,
			},

			ExpectedOutput: &entities.Post{
				ID: 1,
			},
			ExpectedErr: errors.New("access to post denied"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, postID uint) {
				repo.EXPECT().CheckReadAccess(email).Return(false, nil)
				repo.EXPECT().GetPostSenderInfo(postID).Return(nil, nil, nil).AnyTimes()
				repo.EXPECT().SelectPostById(postID, email).Return(&entities.Post{
					ID: 1,
				}, nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			postService := NewPostService(mockPostRepository)

			test.SetupPostMock(mockPostRepository, test.Email, test.Dto.PostID)
			posts, err := postService.GetPostById(test.Email, test.Dto)
			if err != nil {
				require.Equal(t, err, test.ExpectedErr)
			} else {
				require.Equal(t, test.ExpectedOutput, posts)
			}
		})
	}
}

var communityLink string = "someLink"

func TestPostService_GetPostByCommunityLink(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.PostsGetByLink

		ExpectedOutput []*entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink)
	}{
		{
			Name:  "Existing posts with access",
			Email: "test@gmail.com",
			Dto: &dto.PostsGetByLink{
				CommunityLink: &communityLink,
			},

			ExpectedOutput: []*entities.Post{
				{
					ID:           1,
					Text:         "First post",
					CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
				},
			},
			ExpectedErr: nil,

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink) {
				repo.EXPECT().CheckReadAccess(email).Return(true, nil)
				repo.EXPECT().SelectPostsByCommunityLink(dto, email).Return([]*entities.Post{
					{
						ID:           1,
						Text:         "First post",
						CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
					},
				}, nil).AnyTimes()
				var ind uint = 1
				repo.EXPECT().GetPostSenderInfo(ind).Return(nil, nil, nil).AnyTimes()
			},
		},

		{

			Name:  "Existing posts with no access",
			Email: "test@gmail.com",
			Dto: &dto.PostsGetByLink{
				CommunityLink: &communityLink,
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("access to posts denied"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink) {
				repo.EXPECT().CheckReadAccess(email).Return(false, nil)
				repo.EXPECT().SelectPostsByCommunityLink(dto, email).Return([]*entities.Post{
					{
						ID:           1,
						Text:         "First post",
						CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
						CommunityInfo: &entities.CommunityInfo{
							Link: &communityLink,
						},
					},
				}, nil).AnyTimes()

				repo.EXPECT().GetPostSenderInfo(1).Return(nil, nil, nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			postService := NewPostService(mockPostRepository)

			test.SetupPostMock(mockPostRepository, test.Email, test.Dto)
			posts, err := postService.GetPostsByCommunityLink(test.Email, test.Dto)
			if err != nil {
				require.Equal(t, err, test.ExpectedErr)
			} else {
				require.Equal(t, test.ExpectedOutput, posts)
			}
		})
	}
}

var ownerLink string = "someLink"

func TestPostService_GetPostsByUserLink(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.PostsGetByLink

		ExpectedOutput []*entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink)
	}{
		{
			Name:  "Existing posts with access",
			Email: "test@gmail.com",
			Dto: &dto.PostsGetByLink{
				OwnerLink: &ownerLink,
			},

			ExpectedOutput: []*entities.Post{
				{
					ID:           1,
					Text:         "First post",
					CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
				},
			},
			ExpectedErr: nil,

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink) {
				repo.EXPECT().CheckReadAccess(email).Return(true, nil)
				repo.EXPECT().SelectPostsByUserLink(dto, email).Return([]*entities.Post{
					{
						ID:           1,
						Text:         "First post",
						CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
					},
				}, nil).AnyTimes()
				var ind uint = 1
				repo.EXPECT().GetPostSenderInfo(ind).Return(nil, nil, nil).AnyTimes()
			},
		},

		{

			Name:  "Existing posts with no access",
			Email: "test@gmail.com",
			Dto: &dto.PostsGetByLink{
				OwnerLink: &ownerLink,
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("access to posts denied"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostsGetByLink) {
				repo.EXPECT().CheckReadAccess(email).Return(false, nil)
				repo.EXPECT().SelectPostsByUserLink(email, dto).Return([]*entities.Post{
					{
						ID:           1,
						Text:         "First post",
						CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
					},
				}, nil).AnyTimes()

				repo.EXPECT().GetPostSenderInfo(1).Return(nil, nil, nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			postService := NewPostService(mockPostRepository)

			test.SetupPostMock(mockPostRepository, test.Email, test.Dto)
			posts, err := postService.GetPostsByUserLink(test.Email, test.Dto)
			if err != nil {
				require.Equal(t, err, test.ExpectedErr)
			} else {
				require.Equal(t, test.ExpectedOutput, posts)
			}
		})
	}
}

func TestPostService_CreatePost(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.PostCreate

		ExpectedOutput *entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate)
	}{
		{
			Name:  "Empty post data",
			Email: "test@gmail.com",
			Dto: &dto.PostCreate{
				Text:        "",
				Attachments: nil,
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("empty input data"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate) {
				repo.EXPECT().CheckWriteAccess(email, dto).Return(false, nil).AnyTimes()
				var postID uint = 1
				repo.EXPECT().CreatePost(email, dto).Return(postID, nil).AnyTimes()
				repo.EXPECT().SelectPostById(postID, email).Return(nil, nil).AnyTimes()
			},
		},

		{
			Name:  "Too mane post data",
			Email: "test@gmail.com",
			Dto: &dto.PostCreate{
				CommunityLink: &communityLink,
				OwnerLink:     &ownerLink,
				Text:          "someText",
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("too many data (community_link and owner_link can't come together)"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate) {
				repo.EXPECT().CheckWriteAccess(email, dto).Return(false, nil).AnyTimes()
				var postID uint = 1
				repo.EXPECT().CreatePost(email, dto).Return(postID, nil).AnyTimes()
				repo.EXPECT().SelectPostById(postID, email).Return(nil, nil).AnyTimes()
			},
		},

		{
			Name:  "Not enough input data",
			Email: "test@gmail.com",
			Dto: &dto.PostCreate{
				CommunityLink: nil,
				OwnerLink:     nil,
				Text:          "someText",
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("not enough input data"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate) {
				repo.EXPECT().CheckWriteAccess(email, dto).Return(false, nil).AnyTimes()
				var postID uint = 1
				repo.EXPECT().CreatePost(email, dto).Return(postID, nil).AnyTimes()
				repo.EXPECT().SelectPostById(postID, email).Return(nil, nil).AnyTimes()
			},
		},

		{
			Name:  "Access to posts denied",
			Email: "test@gmail.com",
			Dto: &dto.PostCreate{
				CommunityLink: &communityLink,
				OwnerLink:     nil,
				Text:          "someText",
			},

			ExpectedOutput: nil,
			ExpectedErr:    errors.New("access to posts denied"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate) {
				repo.EXPECT().CheckWriteAccess(email, dto).Return(false, nil).AnyTimes()

				var postID uint = 1
				repo.EXPECT().CreatePost(email, dto).Return(postID, nil).AnyTimes()
				repo.EXPECT().SelectPostById(postID, email).Return(nil, nil).AnyTimes()
			},
		},

		{
			Name:  "Successfully created",
			Email: "test@gmail.com",
			Dto: &dto.PostCreate{
				CommunityLink: &communityLink,
				OwnerLink:     nil,
				Text:          "someText",
			},

			ExpectedOutput: &entities.Post{
				ID: 1,
			},
			ExpectedErr: nil,

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostCreate) {
				repo.EXPECT().CheckWriteAccess(email, dto).Return(true, nil).AnyTimes()
				var postID uint = 1
				repo.EXPECT().CreatePost(email, dto).Return(postID, nil).AnyTimes()

				post := &entities.Post{
					ID: postID,
				}
				repo.EXPECT().SelectPostById(postID, email).Return(post, nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			postService := NewPostService(mockPostRepository)

			test.SetupPostMock(mockPostRepository, test.Email, test.Dto)
			posts, err := postService.CreatePost(test.Email, test.Dto)
			if err != nil {
				require.Equal(t, err, test.ExpectedErr)
			} else {
				require.Equal(t, test.ExpectedOutput, posts)
			}
		})
	}
}

var postID uint = 1

func TestPostService_UpdatePost(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.PostUpdate

		ExpectedOutput []*entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostUpdate)
	}{

		{
			Name:  "Missing post_id",
			Email: "test@gmail.com",
			Dto: &dto.PostUpdate{
				PostID: nil,
			},

			ExpectedErr: errors.New("post_id is required field"),

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostUpdate) {
				repo.EXPECT().UpdatePost(email, dto).Return(nil).AnyTimes()
			},
		},

		{
			Name:  "Missing post_id",
			Email: "test@gmail.com",
			Dto: &dto.PostUpdate{
				PostID: &postID,
			},

			ExpectedErr: nil,

			SetupPostMock: func(repo *mock_repository.MockPostRepository, email string, dto *dto.PostUpdate) {
				repo.EXPECT().UpdatePost(email, dto).Return(nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			postService := NewPostService(mockPostRepository)

			test.SetupPostMock(mockPostRepository, test.Email, test.Dto)
			err := postService.UpdatePost(test.Email, test.Dto)

			require.Equal(t, err, test.ExpectedErr)
		})
	}
}
