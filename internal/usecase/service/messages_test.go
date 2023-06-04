package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"depeche/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	msgText        = "someText"
	createdAt      = utils.CurrentTimeString()
	userID    uint = 1
	chatID    uint = 1
)

func TestMessageService_Send(t *testing.T) {
	tests := []struct {
		Name  string
		Email string
		Dto   *dto.NewMessageDTO

		ExpectedOutput *entities.Message
		ExpectedErr    error

		SetupMessageMock func(repo *mock_repository.MockMessageRepository, email string, dto *dto.NewMessageDTO)

		SetupUserMock func(repo *mock_repository.MockUserRepository, email string)
	}{

		{
			Name:  "Successfully sent",
			Email: "test@gmail.com",
			Dto: &dto.NewMessageDTO{
				UserId: 1,
				Text:   "someText",
			},

			ExpectedErr: nil,
			ExpectedOutput: &entities.Message{
				Text:   &msgText,
				UserId: &userID,
				Id:     &userID,
			},

			SetupMessageMock: func(repo *mock_repository.MockMessageRepository, email string, dto *dto.NewMessageDTO) {
				var id uint = 1
				returned := &entities.Message{
					Text:      &msgText,
					UserId:    &userID,
					Id:        &id,
					CreatedAt: &createdAt,
					ChatId:    &chatID,
				}
				repo.EXPECT().SaveMsg(dto).Return(returned, nil).AnyTimes()

				repo.EXPECT().GetUserInfoByMessageId(id).Return(&entities.UserInfo{}, nil).AnyTimes()
				repo.EXPECT().SetLastRead(email, int(*returned.ChatId), *returned.CreatedAt)
			},

			SetupUserMock: func(repo *mock_repository.MockUserRepository, email string) {
				repo.EXPECT().GetUserByEmail(email).Return(&entities.User{ID: userID}, nil).AnyTimes()
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMessageRepository := mock_repository.NewMockMessageRepository(ctrl)
			mockUserRepository := mock_repository.NewMockUserRepository(ctrl)

			postService := NewMessageService(mockMessageRepository, mockUserRepository)

			test.SetupMessageMock(mockMessageRepository, test.Email, test.Dto)
			test.SetupUserMock(mockUserRepository, test.Email)
			_, err := postService.Send(test.Email, test.Dto)

			require.Equal(t, err, nil)
		})
	}
}
