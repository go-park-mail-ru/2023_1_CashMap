package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"depeche/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFeedService_CollectPosts(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.FeedDTO

		ExpectedOutput []*entities.Post
		ExpectedErr    error

		SetupPostMock func(repo *mock_repository.MockPostRepository)

		SetupFeedMock func(repo *mock_repository.MockFeedRepository, email string, dto *dto.FeedDTO)
	}{
		{
			Name:  "Not empty sorted by date feed",
			Email: "test@gmail.com",
			Dto: &dto.FeedDTO{
				BatchSize: 10,
			},

			ExpectedOutput: []*entities.Post{
				somePosts[2],
				somePosts[0],
				somePosts[1],
				somePosts[4],
				somePosts[3],
				somePosts[5],
			},
			ExpectedErr: nil,

			SetupFeedMock: func(repo *mock_repository.MockFeedRepository, email string, dto *dto.FeedDTO) {
				repo.EXPECT().GetGroupsPosts(email, dto).Return(nil, nil)

				repo.EXPECT().GetFriendsPosts(email, dto).Return(somePosts, nil)
			},

			SetupPostMock: func(repo *mock_repository.MockPostRepository) {
				for ind, _ := range somePosts {
					repo.EXPECT().GetPostSenderInfo(uint(ind)+1).Return(nil, nil, nil)
				}

			},
		},

		{
			Name:  "Feed with batch size",
			Email: "test@gmail.com",
			Dto: &dto.FeedDTO{
				BatchSize: 2,
			},

			ExpectedOutput: []*entities.Post{
				somePosts[2],
				somePosts[0],
			},
			ExpectedErr: nil,

			SetupFeedMock: func(repo *mock_repository.MockFeedRepository, email string, dto *dto.FeedDTO) {
				repo.EXPECT().GetGroupsPosts(email, dto).Return(nil, nil)

				repo.EXPECT().GetFriendsPosts(email, dto).Return(somePosts, nil)
			},

			SetupPostMock: func(repo *mock_repository.MockPostRepository) {
				for ind, _ := range somePosts {
					repo.EXPECT().GetPostSenderInfo(uint(ind)+1).Return(nil, nil, nil)
				}

			},
		},

		{
			Name:  "Batch size = 0",
			Email: "test@gmail.com",
			Dto: &dto.FeedDTO{
				BatchSize: 0,
			},

			ExpectedOutput: nil,
			ExpectedErr:    apperror.BadRequest,

			SetupFeedMock: func(repo *mock_repository.MockFeedRepository, email string, dto *dto.FeedDTO) {
				repo.EXPECT().GetGroupsPosts(email, dto).Return(nil, nil).AnyTimes()

				repo.EXPECT().GetFriendsPosts(email, dto).Return(nil, nil).AnyTimes()
			},

			SetupPostMock: func(repo *mock_repository.MockPostRepository) {
				for ind, _ := range somePosts {
					repo.EXPECT().GetPostSenderInfo(uint(ind)+1).Return(nil, nil, nil).AnyTimes()
				}

			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFeedRepository := mock_repository.NewMockFeedRepository(ctrl)
			mockPostRepository := mock_repository.NewMockPostRepository(ctrl)

			feedService := FeedService{
				repository: mockFeedRepository,
				postRepo:   mockPostRepository,
			}

			test.SetupFeedMock(mockFeedRepository, test.Email, test.Dto)
			test.SetupPostMock(mockPostRepository)
			posts, err := feedService.CollectPosts(test.Email, test.Dto)
			if err != nil {
				require.ErrorIs(t, err, test.ExpectedErr)
			} else {
				require.Equal(t, test.ExpectedOutput, posts)
			}
		})
	}
}

var somePosts = []*entities.Post{
	{
		ID:           1,
		Text:         "First post",
		CreationDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).String(),
	},

	{
		ID:           2,
		Text:         "Second post",
		CreationDate: time.Date(2023, 9, 10, 0, 0, 0, 0, time.UTC).String(),
	},

	{
		ID:           3,
		Text:         "Third post",
		CreationDate: time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC).String(),
	},

	{
		ID:           4,
		Text:         "Fourth post",
		CreationDate: time.Date(2023, 3, 10, 0, 0, 0, 0, time.UTC).String(),
	},

	{
		ID:           5,
		Text:         "Fifth post",
		CreationDate: time.Date(2023, 3, 12, 0, 0, 0, 0, time.UTC).String(),
	},

	{
		ID:           6,
		Text:         "Fifth post",
		CreationDate: time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC).String(),
	},
}
