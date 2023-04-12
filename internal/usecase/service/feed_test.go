package service

import (
	"depeche/internal/delivery/dto"
	"depeche/internal/entities"
	mock_repository "depeche/internal/repository/mocks"
	"depeche/internal/utils"
	"depeche/pkg/apperror"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFeedService_CollectPosts(t *testing.T) {
	tests := []struct {
		name          string
		inputDTO      *dto.SignIn
		expectedUser  *entities.User
		expectedError error

		setupMock func(repo *mock_repository.MockUserRepository, in *dto.SignIn)
	}{
		{
			name: "Successful",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  &entities.User{},
			expectedError: nil,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(&entities.User{Password: utils.Hash("Qwerty123")}, nil)
			},
		},
		{
			name: "User doesn't exist",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  nil,
			expectedError: apperror.UserNotFound,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(nil, apperror.UserNotFound)
			},
		},
		{
			name: "Incorrect password",
			inputDTO: &dto.SignIn{
				Email:    "e-larkin@mail.ru",
				Password: "Qwerty123",
			},
			expectedUser:  nil,
			expectedError: apperror.IncorrectCredentials,

			setupMock: func(repo *mock_repository.MockUserRepository, in *dto.SignIn) {
				repo.EXPECT().GetUserByEmail(in.Email).Return(&entities.User{Password: "Another Password"}, nil)
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepository := mock_repository.NewMockUserRepository(ctrl)

			userService := UserService{
				repo: mockRepository,
			}
			test.setupMock(mockRepository, test.inputDTO)
			_, err := userService.SignIn(test.inputDTO)
			require.Equal(t, test.expectedError, err)

		})
	}
}
